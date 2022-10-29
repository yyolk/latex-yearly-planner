package components

import (
	"fmt"

	"github.com/kudrykv/latex-yearly-planner/app/calendar"
)

type WeekLink struct {
	week   calendar.Week
	format string
}

func NewWeekLink(week calendar.Week) WeekLink {
	return WeekLink{week: week}
}

func (r WeekLink) Format(format string) WeekLink {
	r.format = format

	return r
}

func (r WeekLink) String() string {
	format := r.format
	if format == "" {
		format = "Week %d"
	}

	return fmt.Sprintf(`\hyperlink{%s}{`+format+`}`, r.Reference(), r.week.WeekNumber())
}

func (r WeekLink) Reference() string {
	var prefix string

	if r.week.First() {
		prefix = "first-"
	}

	return fmt.Sprintf("%sweek-%d", prefix, r.week.WeekNumber())
}
