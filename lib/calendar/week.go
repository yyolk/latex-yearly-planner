package calendar

import "time"

type Weeks []Week

type Week struct {
	Days [7]Day
}

type WeekOption func() Day

func FromDay(day Day) WeekOption {
	return func() Day {
		return day
	}
}

func FromTime(moment time.Time) WeekOption {
	return func() Day {
		return Day{Time: moment}
	}
}

func NewWeek(wo WeekOption) Week {
	week := Week{} //nolint:exhaustivestruct
	week.Days[0] = wo()

	for i := 1; i < 7; i++ {
		week.Days[i] = week.Days[i-1].Add(1)
	}

	return week
}
