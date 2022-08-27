package sections

import (
	"time"

	"github.com/kudrykv/latex-yearly-planner/lib/calendar"
)

type DailyParameters struct {
}

type Daily struct {
	day        calendar.Day
	parameters DailyParameters
}

func NewDaily(day calendar.Day, parameters DailyParameters) Daily {
	return Daily{
		day:        day,
		parameters: parameters,
	}
}

func (d Daily) Build() ([]string, error) {
	return []string{d.day.Format(time.RFC822)}, nil
}
