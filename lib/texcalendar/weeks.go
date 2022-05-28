package texcalendar

import (
	"strconv"
	"strings"

	"github.com/kudrykv/latex-yearly-planner/lib/calendar"
)

type Weeks struct {
	weeks calendar.Weeks
}

func NewWeeks(weeks calendar.Weeks) Weeks {
	return Weeks{weeks: weeks}
}

func (r Weeks) Tabular() string {
	out := make([]string, 0, len(r.weeks))

	for _, week := range r.weeks {
		out = append(out, NewWeek(week).Tabular())
	}

	return strings.Join(out, `\\`+"\n")
}

type Week struct {
	week calendar.Week
}

func NewWeek(week calendar.Week) Week {
	return Week{week: week}
}

func (r Week) Tabular() string {
	return strings.Join(append([]string{strconv.Itoa(r.week.WeekNumber())}, r.weekDays()...), " & ")
}

func (r Week) weekDays() []string {
	names := make([]string, 0, 7)

	for _, day := range r.week.Days {
		if day.IsZero() {
			names = append(names, "")

			continue
		}

		names = append(names, strconv.Itoa(day.Day()))
	}

	return names
}
