package mos

import (
	"time"

	"github.com/kudrykv/latex-yearly-planner/app2/tex/ref"
	"github.com/kudrykv/latex-yearly-planner/lib/calendar"
	"github.com/kudrykv/latex-yearly-planner/lib/texcalendar"
)

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

	return ref.NewLinkWithRef(day.NameAndDate(), day.Ref()).Build()
}
