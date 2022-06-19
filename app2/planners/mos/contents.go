package mos

import (
	"time"

	"github.com/kudrykv/latex-yearly-planner/app2/planners/common"
	"github.com/kudrykv/latex-yearly-planner/lib/calendar"
	"github.com/kudrykv/latex-yearly-planner/lib/texcalendar"
)

type annualContents struct {
	year calendar.Year
	hand common.MainHand
}

func (m annualContents) Build() ([]string, error) {
	texYear := texcalendar.NewYear(m.hand, m.year)

	return []string{texYear.BuildCalendar()}, nil
}

type quarterlyContents struct {
	quarter calendar.Quarter
	hand    common.MainHand
}

func (r quarterlyContents) Build() ([]string, error) {
	monthsColumn := texcalendar.NewQuarter(r.hand, r.quarter).Column()
	_ = monthsColumn

	return []string{
		`\begin{minipage}[t][20.942cm]{4.5cm}
` + monthsColumn + `
\end{minipage}%
\hspace{6mm}%
\begin{minipage}[t][1cm]{1cm}%
\vbox to -1.8mm{\myDotGrid{41}{19}}
\end{minipage}`,
	}, nil
}

type monthlyContents struct {
	month calendar.Month
	hand  common.MainHand
}

func (m monthlyContents) Build() ([]string, error) {
	month := texcalendar.NewMonth(m.month, m.hand)

	return []string{
		month.LargeCalendar() +
			`

\vspace{3mm}
\myUnderline{Notes}
\vspace{5mm}
\hspace{0.5mm}\vbox to 0mm{\myDotGrid{25}{29}}`,
	}, nil
}

type weeklyContents struct {
	week calendar.Week
}

func (m weeklyContents) Build() ([]string, error) {
	days := texcalendar.NewDays(m.week.Days[:])

	dotsMash := `\vskip5mm\myDotGrid{12}{29}`

	return []string{
		`\vskip1mm\parbox{\myLengthThreeColumnWidth}{\myUnderline{` + m.dayName(m.week, days[0]) + `}}%
\hspace{\myLengthThreeColumnsSeparatorWidth}%
\parbox{\myLengthThreeColumnWidth}{\myUnderline{` + m.dayName(m.week, days[1]) + `}}%
\hspace{\myLengthThreeColumnsSeparatorWidth}%
\parbox{\myLengthThreeColumnWidth}{\myUnderline{` + m.dayName(m.week, days[2]) + `}}
` + dotsMash + `
\vfill

\parbox{\myLengthThreeColumnWidth}{\myUnderline{` + m.dayName(m.week, days[3]) + `}}%
\hspace{\myLengthThreeColumnsSeparatorWidth}%
\parbox{\myLengthThreeColumnWidth}{\myUnderline{` + m.dayName(m.week, days[4]) + `}}%
\hspace{\myLengthThreeColumnsSeparatorWidth}%
\parbox{\myLengthThreeColumnWidth}{\myUnderline{` + m.dayName(m.week, days[5]) + `}}
` + dotsMash + `
\vfill

\parbox{\myLengthThreeColumnWidth}{\myUnderline{` + m.dayName(m.week, days[6]) + `}}%
\hspace{\myLengthThreeColumnsSeparatorWidth}%
\parbox{\dimexpr2\myLengthThreeColumnWidth+\myLengthThreeColumnsSeparatorWidth}{\myUnderline{Notes\textcolor{white}{Q}}}
` + dotsMash,
	}, nil
}

func (m weeklyContents) dayName(week calendar.Week, day texcalendar.Day) string {
	if week.First() && day.Day.Month() != time.January {
		return `\textcolor{white}{Q}`
	}

	if week.Last() && day.Day.Month() != time.December {
		return `\textcolor{white}{Q}`
	}

	return day.NameAndDate()
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
