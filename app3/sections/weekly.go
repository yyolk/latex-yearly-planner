package sections

import (
	"fmt"

	"github.com/kudrykv/latex-yearly-planner/lib/calendar"
)

type WeeklyParameters struct {
}

type Weekly struct {
	parameters WeeklyParameters
	week       calendar.Week
}

func NewWeekly(week calendar.Week, parameters WeeklyParameters) (Weekly, error) {
	return Weekly{
		parameters: parameters,
		week:       week,
	}, nil
}

func (w Weekly) Reference() string {
	var prefix string

	if w.week.First() {
		prefix = "first-"
	}

	return fmt.Sprintf("%sweek-%d", prefix, w.week.WeekNumber())
}

func (w Weekly) Title() string {
	return fmt.Sprintf("Week %d", w.week.WeekNumber())
}

func (w Weekly) Build() ([]string, error) {
	return []string{
		"hello week",
	}, nil
}
