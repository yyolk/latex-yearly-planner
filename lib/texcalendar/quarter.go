package texcalendar

import (
	"strings"

	"github.com/kudrykv/latex-yearly-planner/app/tex"
	"github.com/kudrykv/latex-yearly-planner/app2/planners/common"
	"github.com/kudrykv/latex-yearly-planner/lib/calendar"
)

type Quarter struct {
	quarter calendar.Quarter
	hand    common.MainHand
}

func NewQuarter(hand common.MainHand, quarter calendar.Quarter) Quarter {
	return Quarter{hand: hand, quarter: quarter}
}

func (q Quarter) Row() string {
	monthsRow := make([]string, 0, len(q.quarter.Months))

	for _, month := range q.quarter.Months {
		monthsRow = append(monthsRow, tex.AdjustBox(NewMonth(month, q.hand).LittleCalendar()))
	}

	return strings.Join(monthsRow, " & ")
}

func (q Quarter) Column() string {
	months := make([]string, 0, 3)

	for _, month := range q.quarter.Months {
		months = append(months, NewMonth(month, q.hand).LittleCalendar())
	}

	return strings.Join(months, "\n\\vfill\n")
}
