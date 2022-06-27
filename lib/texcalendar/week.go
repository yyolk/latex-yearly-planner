package texcalendar

import (
	"strconv"
	"strings"
	"time"

	"github.com/kudrykv/latex-yearly-planner/app2/planners/common"
	"github.com/kudrykv/latex-yearly-planner/app2/tex/ref"
	"github.com/kudrykv/latex-yearly-planner/lib/calendar"
)

type Week struct {
	week        calendar.Week
	parameters  Parameters
	selectedDay Day
}

func NewWeek(week calendar.Week, options ...ApplyToParameters) Week {
	parameters := Parameters{}

	for _, option := range options {
		option(&parameters)
	}

	return Week{week: week, parameters: parameters}
}

func (r Week) Tabular() string {
	return strings.Join(r.BuildLargeCalRow(), " & ")
}

func (r Week) Ref() string {
	refer := "W" + strconv.Itoa(r.week.WeekNumber())

	if r.week.First() {
		refer += "-first"
	}

	return refer
}

func (r Week) Title() string {
	return "Week " + strconv.Itoa(r.week.WeekNumber())
}

func (r Week) Days() Days {
	days := make(Days, 0, 7)

	for _, day := range r.week.Days() {
		days = append(days, NewDay(day))
	}

	return days
}

func (r Week) First() bool {
	return r.week.First()
}

func (r Week) Last() bool {
	return r.week.Last()
}

func (r Week) TailMonth() time.Month {
	return r.week.TailMonth()
}

func (r Week) HeadMonth() time.Month {
	return r.week.HeadMonth()
}

func (r Week) BuildLargeCalRow() []string {
	weekName := `\rotatebox[origin=tr]{90}{\makebox[2cm][c]{` + "Week " + strconv.Itoa(r.week.WeekNumber()) + `}}`
	weekName = ref.NewLinkWithRef(weekName, r.Ref()).Build()

	return r.appendWeekName(weekName, r.Days().BuildLarge())
}

func (r Week) BuildLittleCalRow() []string {
	weekName := ref.NewLinkWithRef(strconv.Itoa(r.week.WeekNumber()), r.Ref()).Build()

	return r.appendWeekName(weekName, r.Days().BuildLittle(r.selectedDay))
}

func (r Week) appendWeekName(name string, weekdays []string) []string {
	if r.parameters.Hand == common.LeftHand {
		return append(weekdays, name)
	}

	return append([]string{name}, weekdays...)
}

func (r Week) Selected(day Day) Week {
	r.selectedDay = day

	return r
}
