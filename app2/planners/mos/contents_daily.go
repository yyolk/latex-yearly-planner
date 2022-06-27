package mos

import (
	"fmt"
	"strings"

	"github.com/kudrykv/latex-yearly-planner/app2/planners/common"
	"github.com/kudrykv/latex-yearly-planner/app2/tex/ref"
	"github.com/kudrykv/latex-yearly-planner/lib/texcalendar"
)

type dailyContents struct {
	day  texcalendar.Day
	hand common.MainHand
}

func (r dailyContents) Build() ([]string, error) {
	leftColumn := r.scheduleColumn()
	rightColumn := r.prioritiesAndNotesColumn()

	if r.hand == common.LeftHand {
		leftColumn, rightColumn = rightColumn, leftColumn
	}

	return []string{
		`\noindent\vskip1mm` + leftColumn + `\hspace{5mm}` + rightColumn,
	}, nil
}

func (r dailyContents) prioritiesAndNotesColumn() string {
	var priorities []string

	for i := 0; i < 8; i++ {
		priorities = append(priorities, r.height()+`$\square$\myLineGray`)
	}

	moreNotes := ref.NewNote("More", r.day.Ref()).Build()
	dailyReflect := ref.NewReflect("Reflect", r.day.Ref()).Build()

	return `\begin{minipage}[t]{\dimexpr2\myLengthThreeColumnWidth+\myLengthThreeColumnsSeparatorWidth}
\myUnderline{Top Priorities}
` + strings.Join(priorities, "\n") + `
\vskip7mm\myUnderline{Notes | ` + moreNotes + ` \hfill{}` + dailyReflect + `}
\vspace{5mm}\hspace{.5mm}\vbox to 0mm{\myDotGrid{30}{19}}
\end{minipage}`
}

func (r dailyContents) scheduleColumn() string {
	var hours []string

	for i := 8; i <= 21; i++ {
		strHour := fmt.Sprintf("%0d", i)
		hours = append(hours, r.height()+strHour+`\myLineLightGray
\vskip5mm\myLineGray`)
	}

	return `\begin{minipage}[t]{\myLengthThreeColumnWidth}
\myUnderline{Schedule\textcolor{white}{g}}
` + strings.Join(hours, "\n") + `
\vskip5mm` + r.day.CalendarMonth().Selected(r.day).LittleCalendar() + `
\end{minipage}`
}

func (r dailyContents) height() string {
	return `\parbox{0pt}{\vskip5mm}`
}
