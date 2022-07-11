package components

import (
	"fmt"
	"strings"
	"time"
)

type Schedule []ScheduleHour

func NewSchedule(fromHour, toHour int, format string) Schedule {
	schedule := make(Schedule, 0, toHour-fromHour+1)

	for i := fromHour; i <= toHour; i++ {
		schedule = append(schedule, NewScheduleHour(i, format))
	}

	return schedule
}

func (r Schedule) Build() string {
	var schedule []string

	for _, hour := range r {
		schedule = append(schedule, hour.Build())
	}

	return strings.Join(schedule, "\n")
}

type ScheduleHour struct {
	Hour   int
	Format string
}

func NewScheduleHour(hour int, format string) ScheduleHour {
	return ScheduleHour{
		Hour:   hour,
		Format: format,
	}
}

func (r ScheduleHour) Build() string {
	return fmt.Sprintf(scheduleHourFormat, r.newHour().Format(r.Format))
}

func (r ScheduleHour) newHour() time.Time {
	return time.Date(0, 0, 0, r.Hour, 0, 0, 0, time.Local)
}

const scheduleHourFormat = `\parbox{0pt}{\vskip5mm}%s\myLineLightGray` + "\n" + `\vskip5mm\myLineGray`
