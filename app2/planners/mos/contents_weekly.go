package mos

import (
	"time"

	"github.com/kudrykv/latex-yearly-planner/app2/tex/ref"
	"github.com/kudrykv/latex-yearly-planner/lib/texcalendar"
)

type weeklyContents struct {
	week texcalendar.Week
}

func (m weeklyContents) Build() ([]string, error) {
	days := m.week.Days()

	dotsMash := `\vskip5mm\myDotGrid{12}{29}`

	return []string{
		`\vskip1mm\parbox{\myLengthThreeColumnWidth}{\myUnderline{` + m.dayName(days[0]) + `}}%
\hspace{\myLengthThreeColumnsSeparatorWidth}%
\parbox{\myLengthThreeColumnWidth}{\myUnderline{` + m.dayName(days[1]) + `}}%
\hspace{\myLengthThreeColumnsSeparatorWidth}%
\parbox{\myLengthThreeColumnWidth}{\myUnderline{` + m.dayName(days[2]) + `}}
` + dotsMash + `
\vfill

\parbox{\myLengthThreeColumnWidth}{\myUnderline{` + m.dayName(days[3]) + `}}%
\hspace{\myLengthThreeColumnsSeparatorWidth}%
\parbox{\myLengthThreeColumnWidth}{\myUnderline{` + m.dayName(days[4]) + `}}%
\hspace{\myLengthThreeColumnsSeparatorWidth}%
\parbox{\myLengthThreeColumnWidth}{\myUnderline{` + m.dayName(days[5]) + `}}
` + dotsMash + `
\vfill

\parbox{\myLengthThreeColumnWidth}{\myUnderline{` + m.dayName(days[6]) + `}}%
\hspace{\myLengthThreeColumnsSeparatorWidth}%
\parbox{\dimexpr2\myLengthThreeColumnWidth+\myLengthThreeColumnsSeparatorWidth}{\myUnderline{Notes\textcolor{white}{Q}}}
` + dotsMash,
	}, nil
}

func (m weeklyContents) dayName(day texcalendar.Day) string {
	if m.week.First() && day.Day.Month() != time.January {
		return `\textcolor{white}{Q}`
	}

	if m.week.Last() && day.Day.Month() != time.December {
		return `\textcolor{white}{Q}`
	}

	return ref.NewLinkWithRef(day.NameAndDate(), day.Ref()).Build()
}
