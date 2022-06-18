package mos

import (
	"strconv"
	"time"

	"github.com/kudrykv/latex-yearly-planner/lib/calendar"
	"github.com/kudrykv/latex-yearly-planner/lib/texcalendar"
)

type annualContents struct {
	year calendar.Year
}

func (m annualContents) Build() ([]string, error) {
	texYear := texcalendar.NewYear(m.year)

	return []string{texYear.BuildCalendar()}, nil
}

type quarterlyContents struct {
	quarter calendar.Quarter
}

func (r quarterlyContents) Build() ([]string, error) {
	monthsColumn := texcalendar.NewQuarter(r.quarter).Column()
	_ = monthsColumn

	return []string{
		`\begin{minipage}[t][20cm]{5cm}
` + monthsColumn + `
\end{minipage}%
\hspace{6mm}%
\begin{minipage}[t][20cm]{8cm}%
\vbox to -1.8mm{\myDotGrid{41}{18}}
\end{minipage}`,
	}, nil
}

type monthlyContents struct {
	month calendar.Month
}

func (m monthlyContents) Build() ([]string, error) {
	return []string{m.month.Month().String()}, nil
}

type weeklyContents struct {
	week calendar.Week
}

func (m weeklyContents) Build() ([]string, error) {
	return []string{strconv.Itoa(m.week.WeekNumber())}, nil
}

type dailyContents struct {
	day calendar.Day
}

func (m dailyContents) Build() ([]string, error) {
	return []string{m.day.Format(time.RFC3339)}, nil
}

type todoIndex struct{}

func (i todoIndex) Build() ([]string, error) {
	return []string{"index"}, nil
}

type todoContents struct{}

func (t todoContents) Build() ([]string, error) {
	return []string{"page with todos"}, nil
}

type notesIndex struct{}

func (r notesIndex) Build() ([]string, error) {
	return []string{"notes index"}, nil
}

type notesContents struct{}

func (r notesContents) Build() ([]string, error) {
	return []string{"notes"}, nil
}
