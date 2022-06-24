package texcalendar

import (
	"strconv"
	"strings"

	"github.com/kudrykv/latex-yearly-planner/app/tex"
	"github.com/kudrykv/latex-yearly-planner/lib/calendar"
)

type Quarters []Quarter

func (q Quarters) Reverse() Quarters {
	if len(q) == 0 {
		return nil
	}

	quarters := make(Quarters, 0, len(q))

	for i := len(q) - 1; i >= 0; i-- {
		quarters = append(quarters, q[i])
	}

	return quarters
}

type Quarter struct {
	quarter    calendar.Quarter
	parameters Parameters
}

func NewQuarter(quarter calendar.Quarter, options ...ApplyToParameters) Quarter {
	parameters := Parameters{}

	for _, option := range options {
		option(&parameters)
	}

	return Quarter{parameters: parameters, quarter: quarter}
}

func (q Quarter) Row() string {
	monthsRow := make([]string, 0, len(q.quarter.Months))

	for _, month := range q.quarter.Months {
		littleCalendar := NewMonth(month, WithParameters(q.parameters)).LittleCalendar()
		monthsRow = append(monthsRow, tex.AdjustBox(littleCalendar))
	}

	return strings.Join(monthsRow, " & ")
}

func (q Quarter) Column() string {
	months := make([]string, 0, 3)

	for _, month := range q.quarter.Months {
		months = append(months, NewMonth(month, WithParameters(q.parameters)).LittleCalendar())
	}

	return strings.Join(months, "\n\\vfill\n")
}

func (q Quarter) Name() string {
	return "Q" + strconv.Itoa(q.quarter.Number())
}

func (q Quarter) Matches(quarter Quarter) bool {
	return q.Name() == quarter.Name()
}
