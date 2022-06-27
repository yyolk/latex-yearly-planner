package texcalendar

import (
	"strings"
	"time"

	"github.com/kudrykv/latex-yearly-planner/app/tex"
	"github.com/kudrykv/latex-yearly-planner/app2/planners/common"
	"github.com/kudrykv/latex-yearly-planner/app2/tex/ref"
	"github.com/kudrykv/latex-yearly-planner/lib/calendar"
)

type Months []Month

type Month struct {
	month       calendar.Month
	selectedDay Day
	parameters  Parameters
}

func NewMonth(month calendar.Month, options ...ApplyToParameters) Month {
	parameters := Parameters{}

	for _, option := range options {
		option(&parameters)
	}

	mo := Month{parameters: parameters, month: month}

	if mo.month.Month() == parameters.FirstMonth {
		mo.month.Weeks[0] = mo.month.Weeks[0].SetFirst()
	}

	return mo
}

func (r Month) LittleCalendar() string {
	weeks := r.weeks()
	weekdays := strings.Join(weeks.WeekdaysShortNames(), " & ")
	littleCalMatrix := r.tabulate(weeks.BuildLittleCalMatrix(), `\\`)

	name := ref.NewText(r.Name(), "").Build()

	tabularContents := `\multicolumn{8}{c}{` + name + `} \\ \hline` + "\n" +
		weekdays + `\\ \hline` + "\n" +
		littleCalMatrix

	return tex.RenewArrayStretch(r.parameters.LittleCalArrayStretch) + `%` + "\n" +
		tex.TabularXAlignTopLineWidth(r.littleTableRule(), tabularContents)
}

func (r Month) LargeCalendar() string {
	weeks := NewWeeks(r.month.Weeks, WithParameters(r.parameters))
	weekdays := strings.Join(weeks.WeekdaysFullNames(), ` & `)
	weeksMatrix := r.tabulate(weeks.BuildLargeCalMatrix(), `\\ \hline`)
	lineHeight := tex.LineHeight(r.parameters.LargeCalHeaderHeight)
	tabularContents := weekdays + lineHeight + ` \\ \hline` + "\n" +
		weeksMatrix + "\\\\ \\hline\n"

	return tex.TabularXAlignTopLineWidth(r.largeTableRule(), tabularContents)
}

func (r Month) name() string {
	return r.month.Month().String()
}

func (r Month) weeks() Weeks {
	return NewWeeks(r.month.Weeks, WithParameters(r.parameters))
}

func (r Month) littleTableRule() string {
	if r.parameters.Hand == common.LeftHand {
		return "*{7}{@{}Y@{}}|c"
	}

	return "c|*{7}{@{}Y@{}}"
}

func (r Month) largeTableRule() string {
	if r.parameters.Hand == common.LeftHand {
		return `*{7}{|@{}X@{}}|@{ }c@{ }|`
	}

	return `|@{ }c@{ }|*{7}{@{}X@{}|}`
}

func (r Month) weekdaysShort() []string {
	weekdays := make([]string, 0, 7)

	for _, weekday := range r.month.Weekdays() {
		weekdays = append(weekdays, weekday.String()[:1])
	}

	return weekdays
}

func (r Month) tabulate(matrix [][]string, join string) string {
	rows := make([]string, 0, len(matrix))

	for _, row := range matrix {
		rows = append(rows, strings.Join(row, " & "))
	}

	return strings.Join(rows, join+"\n")
}

func (r Month) ShortName() string {
	return r.month.Month().String()[:3]
}

func (r Month) Month() time.Month {
	return r.month.Month()
}

func (r Month) IntersectsWith(selectedMonths []time.Month) bool {
	for _, selectedMonth := range selectedMonths {
		if r.Month() == selectedMonth {
			return true
		}
	}

	return false
}

func (r Month) Name() string {
	return r.month.Month().String()
}

func (r Month) Selected(selectedDay Day) Month {
	r.selectedDay = selectedDay

	return r
}
