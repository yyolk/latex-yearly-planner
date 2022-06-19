package texcalendar

import (
	"time"

	"github.com/kudrykv/latex-yearly-planner/lib/calendar"
)

type Days []Day

func NewDays(days []calendar.Day) Days {
	if len(days) == 0 {
		return nil
	}

	out := make(Days, 0, len(days))

	for _, day := range days {
		out = append(out, NewDay(day))
	}

	return out
}

type Day struct {
	Day calendar.Day
}

func NewDay(day calendar.Day) Day {
	return Day{Day: day}
}

func (d Day) NameAndDate() string {
	return d.Day.Format("Monday, _2")
}

func (d Day) Ref() string {
	return d.Day.Format(time.RFC3339)
}
