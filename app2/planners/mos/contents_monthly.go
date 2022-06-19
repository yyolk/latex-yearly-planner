package mos

import (
	"github.com/kudrykv/latex-yearly-planner/app2/planners/common"
	"github.com/kudrykv/latex-yearly-planner/lib/calendar"
	"github.com/kudrykv/latex-yearly-planner/lib/texcalendar"
)

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
