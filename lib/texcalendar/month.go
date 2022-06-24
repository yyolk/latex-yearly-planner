package texcalendar

import (
	"strings"
	"time"

	"github.com/kudrykv/latex-yearly-planner/app2/planners/common"
	"github.com/kudrykv/latex-yearly-planner/lib/calendar"
)

type Months []Month

type Month struct {
	month      calendar.Month
	parameters Parameters
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

func (m Month) LittleCalendar() string {
	monthName := m.month.Month().String()
	weekdays := strings.Join(append([]string{"W"}, m.weekdaysShort()...), ` & `)
	weeksMatrix := m.tabulate(NewWeeks(m.month.Weeks, WithParameters(m.parameters)).Matrix(), `\\`)

	return `\renewcommand{\arraystretch}{1.5}%` + "\n" +
		`%\setlength{\tabcolsep}{3.5pt}%` + "\n" +
		`\begin{tabularx}{\linewidth}[t]{` + m.littleTableRule() + `}` + "\n" +
		`\multicolumn{8}{c}{` + monthName + `} \\ \hline` + "\n" +
		weekdays + `\\ \hline` + "\n" +
		weeksMatrix + "\n" +
		`\end{tabularx}`
}

func (m Month) LargeCalendar() string {
	weeks := NewWeeks(m.month.Weeks, WithParameters(m.parameters))
	weekdays := strings.Join(weeks.Weekdays(), ` & `)
	weeksMatrix := m.tabulate(weeks.Matrix(), `\\ \hline`)

	return `\renewcommand{\arraystretch}{0}%` + "\n" +
		`%\setlength{\tabcolsep}{0pt}%` + "\n" +
		`\begin{tabularx}{\linewidth}[t]{` + m.largeTableRule() + `}` + "\n" +
		weekdays + ` \raisebox{4mm}{} \\[2mm] \hline` + "\n" +
		weeksMatrix + "\\\\ \\hline\n" +
		`\end{tabularx}`
}

func (m Month) littleTableRule() string {
	if m.parameters.Hand == common.LeftHand {
		return "*{7}{@{}Y@{}}|c"
	}

	return "c|*{7}{@{}Y@{}}"
}

func (m Month) largeTableRule() string {
	if m.parameters.Hand == common.LeftHand {
		return `*{7}{|@{}X@{}}|@{ }c@{ }|`
	}

	return `|@{ }c@{ }|*{7}{@{}X@{}|}`
}

func (m Month) weekdaysShort() []string {
	weekdays := make([]string, 0, 7)

	for _, weekday := range m.month.Weekdays() {
		weekdays = append(weekdays, weekday.String()[:1])
	}

	return weekdays
}

func (m Month) tabulate(matrix [][]string, join string) string {
	rows := make([]string, 0, len(matrix))

	for _, row := range matrix {
		rows = append(rows, strings.Join(row, " & "))
	}

	return strings.Join(rows, join+"\n")
}

func (m Month) ShortName() string {
	return m.month.Month().String()[:3]
}

func (m Month) Month() time.Month {
	return m.month.Month()
}

func (m Month) IntersectsWith(selectedMonths []time.Month) bool {
	for _, selectedMonth := range selectedMonths {
		if m.Month() == selectedMonth {
			return true
		}
	}

	return false
}
