package mos

import (
	"github.com/kudrykv/latex-yearly-planner/app2/planners/common"
	"github.com/kudrykv/latex-yearly-planner/lib/calendar"
	"github.com/kudrykv/latex-yearly-planner/lib/texcalendar"
)

type quarterlyContents struct {
	quarter calendar.Quarter
	hand    common.MainHand
}

func (r quarterlyContents) Build() ([]string, error) {
	calendarColumn := r.calendarColumn()
	notesColumn := r.notesColumn()

	if r.hand == common.LeftHand {
		calendarColumn, notesColumn = notesColumn, calendarColumn
	}

	return []string{
		calendarColumn + `\hspace{\myLengthThreeColumnsSeparatorWidth}` + notesColumn,
	}, nil
}

func (r quarterlyContents) notesColumn() string {
	return `\begin{minipage}[t][1cm]{\myLengthTwoThirdsColumnWidth}
\vbox to -1.8mm{\myDotGrid{41}{19}}
\end{minipage}`
}

func (r quarterlyContents) calendarColumn() string {
	return `\begin{minipage}[t][20.942cm]{\myLengthThreeColumnWidth}
` + texcalendar.NewQuarter(r.hand, r.quarter).Column() + `
\end{minipage}`
}
