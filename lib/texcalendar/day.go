package texcalendar

import (
	"time"

	"github.com/kudrykv/latex-yearly-planner/lib/calendar"
)

type Days []Day

type Day struct {
	Day        calendar.Day
	parameters Parameters
}

func NewDay(day calendar.Day, options ...ApplyToParameters) Day {
	parameters := Parameters{}

	for _, option := range options {
		option(&parameters)
	}

	return Day{Day: day, parameters: parameters}
}

func (d Day) NameAndDate() string {
	return d.Day.Format("Monday, _2")
}

func (d Day) Ref() string {
	return d.Day.Format(time.RFC3339)
}

func (d Day) Month() time.Month {
	return d.Day.Month()
}

func (d Day) Week() Week {
	return NewWeek(*d.Day.Week())
}
