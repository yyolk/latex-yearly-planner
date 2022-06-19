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
