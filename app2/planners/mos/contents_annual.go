package mos

import (
	"github.com/kudrykv/latex-yearly-planner/app2/planners/common"
	"github.com/kudrykv/latex-yearly-planner/lib/texcalendar"
)

type annualContents struct {
	year texcalendar.Year
	hand common.MainHand
}

func (m annualContents) Build() ([]string, error) {
	return []string{m.year.BuildCalendar()}, nil
}
