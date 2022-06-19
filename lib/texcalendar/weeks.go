package texcalendar

import (
	"strconv"
	"strings"

	"github.com/kudrykv/latex-yearly-planner/app2/planners/common"
	"github.com/kudrykv/latex-yearly-planner/app2/tex/ref"
	"github.com/kudrykv/latex-yearly-planner/lib/calendar"
)

type Weeks struct {
	weeks    calendar.Weeks
	hand     common.MainHand
	forLarge bool
}

func NewWeeks(hand common.MainHand, weeks calendar.Weeks, forLarge bool) Weeks {
	return Weeks{hand: hand, weeks: weeks, forLarge: forLarge}
}

func (r Weeks) Weekdays() []string {
	if len(r.weeks) == 0 {
		return nil
	}

	weekdays := make([]string, 0, 8)
	for _, day := range r.weeks[0].Next().Days {
		weekdays = append(weekdays, `\hfil{}`+day.Weekday().String())
	}

	if r.hand == common.RightHand {
		weekdays = append([]string{"W"}, weekdays...)
	} else {
		weekdays = append(weekdays, "W")
	}

	return weekdays
}

func (r Weeks) Tabular() string {
	out := make([]string, 0, len(r.weeks))

	for _, week := range r.weeks {
		out = append(out, NewWeek(r.hand, week, r.forLarge).Tabular())
	}

	return strings.Join(out, `\\`+"\n")
}

func (r Weeks) Matrix() [][]string {
	rows := make([][]string, 0, len(r.weeks))

	for _, week := range r.weeks {
		rows = append(rows, NewWeek(r.hand, week, r.forLarge).Row())
	}

	return rows
}

type Week struct {
	week     calendar.Week
	hand     common.MainHand
	forLarge bool
}

func NewWeek(hand common.MainHand, week calendar.Week, forLarge bool) Week {
	return Week{hand: hand, week: week, forLarge: forLarge}
}

func (r Week) Tabular() string {
	return strings.Join(r.Row(), " & ")
}

func (r Week) weekDays() []string {
	names := make([]string, 0, 7)

	for _, day := range r.week.Days {
		if day.IsZero() {
			names = append(names, "")

			continue
		}

		name := strconv.Itoa(day.Day())

		if r.forLarge {
			name = `{\renewcommand{\arraystretch}{1.2}\begin{tabular}{@{}p{5mm}@{}|}\hfil{}` + name + `\\ \hline\end{tabular}}`
		}

		names = append(names, name)
	}

	return names
}

func (r Week) Row() []string {
	weekName := strconv.Itoa(r.week.WeekNumber())

	if r.forLarge {
		weekName = `\rotatebox[origin=tr]{90}{\makebox[2cm][c]{` + "Week " + weekName + `}}`
	}

	weekName = ref.NewLinkWithRef(weekName, r.Ref()).Build()

	if r.hand == common.LeftHand {
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
