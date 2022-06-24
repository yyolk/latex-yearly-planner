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
	week       calendar.Week
	parameters Parameters
}

func NewWeek(week calendar.Week, options ...ApplyToParameters) Week {
	parameters := Parameters{}

	for _, option := range options {
		option(&parameters)
	}

	return Week{week: week, parameters: parameters}
}

func (r Week) Tabular() string {
	return strings.Join(r.Row(), " & ")
}

func (r Week) weekDays() []string {
	names := make([]string, 0, 7)

	for _, day := range r.week.Days() {
		if day.IsZero() {
			names = append(names, "")

			continue
		}

		name := strconv.Itoa(day.Day())

		if r.parameters.ForLarge {
			name = `{\renewcommand{\arraystretch}{1.2}\begin{tabular}{@{}p{5mm}@{}|}\hfil{}` + name + `\\ \hline\end{tabular}}`
		}

		names = append(names, ref.NewLinkWithRef(name, NewDay(day).Ref()).Build())
	}

	return names
}

func (r Week) Row() []string {
	weekName := strconv.Itoa(r.week.WeekNumber())

	if r.parameters.ForLarge {
		weekName = `\rotatebox[origin=tr]{90}{\makebox[2cm][c]{` + "Week " + weekName + `}}`
	}

	weekName = ref.NewLinkWithRef(weekName, r.Ref()).Build()

	if r.parameters.Hand == common.LeftHand {
		return append(r.weekDays(), weekName)
	}

	return append([]string{weekName}, r.weekDays()...)
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
