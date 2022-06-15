package mos

import (
	"github.com/kudrykv/latex-yearly-planner/lib/calendar"
)

type mosMonthlyContents struct {
	month calendar.Month
}

func (m mosMonthlyContents) Build() ([]string, error) {
	return []string{m.month.Month().String()}, nil
}
