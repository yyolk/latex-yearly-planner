package texcalendar

import (
	"strings"

	"github.com/kudrykv/latex-yearly-planner/app/tex"
	"github.com/kudrykv/latex-yearly-planner/lib/calendar"
)

type Quarter struct {
	quarter calendar.Quarter
}

func NewQuarter(quarter calendar.Quarter) Quarter {
	return Quarter{quarter: quarter}
}

func (q Quarter) Row() string {
	monthsRow := make([]string, 0, len(q.quarter.Months))

	for _, month := range q.quarter.Months {
		monthsRow = append(monthsRow, tex.AdjustBox(NewMonth(month).Tabular()))
	}

	return strings.Join(monthsRow, " & ")
}
