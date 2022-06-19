package mos

import (
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
