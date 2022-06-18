package texcalendar

import (
	"strconv"
	"strings"

	"github.com/kudrykv/latex-yearly-planner/app2/planners/common"
	"github.com/kudrykv/latex-yearly-planner/lib/calendar"
)

type Weeks struct {
	weeks calendar.Weeks
	hand  common.MainHand
}

func NewWeeks(hand common.MainHand, weeks calendar.Weeks) Weeks {
	return Weeks{hand: hand, weeks: weeks}
}

func (r Weeks) Tabular() string {
	out := make([]string, 0, len(r.weeks))

	for _, week := range r.weeks {
		out = append(out, NewWeek(r.hand, week).Tabular())
	}

	return strings.Join(out, `\\`+"\n")
}

func (r Weeks) Matrix() [][]string {
	rows := make([][]string, 0, len(r.weeks))

	for _, week := range r.weeks {
		rows = append(rows, NewWeek(r.hand, week).Row())
	}

	return rows
}

type Week struct {
	week calendar.Week
	hand common.MainHand
}

func NewWeek(hand common.MainHand, week calendar.Week) Week {
	return Week{hand: hand, week: week}
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

		names = append(names, strconv.Itoa(day.Day()))
	}

	return names
}

func (r Week) Row() []string {
	if r.hand == common.LeftHand {
		return append(r.weekDays(), strconv.Itoa(r.week.WeekNumber()))
	}

	return append([]string{strconv.Itoa(r.week.WeekNumber())}, r.weekDays()...)
}
