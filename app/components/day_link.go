package components

import (
	"fmt"

	"github.com/kudrykv/latex-yearly-planner/app/calendar"
)

type DayLink struct {
	day    calendar.Day
	format string
}

func NewDayLink(day calendar.Day) DayLink {
	return DayLink{day: day}
}

func (r DayLink) Format(format string) DayLink {
	r.format = format

	return r
}

func (r DayLink) String() string {
	format := r.format
	if format == "" {
		format = "2"
	}

	return fmt.Sprintf(`\hyperlink{%s}{%s}`, r.day.Format("2006-01-02"), r.day.Format(format))
}
