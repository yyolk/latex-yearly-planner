package sections

import (
	"fmt"

	"github.com/kudrykv/latex-yearly-planner/app3/components"
	"github.com/kudrykv/latex-yearly-planner/lib/calendar"
)

type DailyParameters struct {
	ScheduleParameters components.ScheduleParameters
}

type Daily struct {
	day        calendar.Day
	parameters DailyParameters
	schedule   components.Schedule
}

func NewDaily(day calendar.Day, parameters DailyParameters) (Daily, error) {
	schedule, err := components.NewSchedule(day.Moment(), parameters.ScheduleParameters)
	if err != nil {
		return Daily{}, fmt.Errorf("new schedule: %w", err)
	}

	return Daily{
		day:        day,
		parameters: parameters,
		schedule:   schedule,
	}, nil
}

func (d Daily) Build() ([]string, error) {
	return []string{d.schedule.Build()}, nil
}
