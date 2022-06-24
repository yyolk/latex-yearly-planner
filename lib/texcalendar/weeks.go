package texcalendar

import (
	"strings"

	"github.com/kudrykv/latex-yearly-planner/app2/planners/common"
	"github.com/kudrykv/latex-yearly-planner/lib/calendar"
)

type Weeks struct {
	weeks      calendar.Weeks
	parameters Parameters
}

func NewWeeks(weeks calendar.Weeks, options ...ApplyToParameters) Weeks {
	parameters := Parameters{}

	for _, option := range options {
		option(&parameters)
	}

	return Weeks{weeks: weeks, parameters: parameters}
}

func (r Weeks) WeekdaysFullNames() []string {
	if len(r.weeks) == 0 {
		return nil
	}

	return r.centerline(r.extendWithW(r.buildWeekdays()))
}

func (r Weeks) WeekdaysShortNames() []string {
	if len(r.weeks) == 0 {
		return nil
	}

	return r.centerline(r.extendWithW(r.buildWeekdaysWithShortNames()))
}

func (r Weeks) centerline(row []string) []string {
	for i, item := range row {
		row[i] = `\hfil{}` + item
	}

	return row
}

func (r Weeks) extendWithW(row []string) []string {
	if r.parameters.Hand == common.RightHand {
		row = append([]string{"W"}, row...)
	} else {
		row = append(row, "W")
	}

	return row
}

func (r Weeks) buildWeekdaysWithShortNames() []string {
	weekdays := make([]string, 0, 8)

	for _, weekday := range r.buildWeekdays() {
		weekdays = append(weekdays, weekday[:1])
	}

	return weekdays
}

func (r Weeks) buildWeekdays() []string {
	weekdays := make([]string, 0, 8)

	for _, day := range r.weeks[0].Next().Days() {
		weekdays = append(weekdays, day.Weekday().String())
	}

	return weekdays
}

func (r Weeks) Tabular() string {
	out := make([]string, 0, len(r.weeks))

	for _, week := range r.weeks {
		out = append(out, NewWeek(week, WithParameters(r.parameters)).Tabular())
	}

	return strings.Join(out, `\\`+"\n")
}

func (r Weeks) Matrix() [][]string {
	rows := make([][]string, 0, len(r.weeks))

	for _, week := range r.weeks {
		rows = append(rows, NewWeek(week, WithParameters(r.parameters)).Row())
	}

	return rows
}
