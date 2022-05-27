package texcalendar

import (
	"strings"

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

	adjustBox := func(str string) string { return `\adjustbox{valign=t}{` + str + `}` }

	for _, month := range q.quarter.Months {
		monthsRow = append(monthsRow, adjustBox(NewMonth(month).Tabular()))
	}

	return strings.Join(monthsRow, " & ")
}
