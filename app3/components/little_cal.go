package components

import (
	"errors"
	"fmt"

	"github.com/kudrykv/latex-yearly-planner/lib/calendar"
)

type LittleCalendarParameters struct {
}

type LittleCalendar struct {
	today      calendar.Day
	parameters LittleCalendarParameters
}

var ErrNoMonth = errors.New("no month")

func NewLittleCalendar(today calendar.Day, parameters LittleCalendarParameters) (LittleCalendar, error) {
	month := today.CalendarMonth()
	if month == nil {
		return LittleCalendar{}, fmt.Errorf("day doesn't return its month: %w", ErrNoMonth)
	}

	return LittleCalendar{
		today:      today,
		parameters: parameters,
	}, nil
}

func (r LittleCalendar) Build() string {
	panic("implement me")
}
