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
	return Week{Days: [7]Day{wo()}}.fillFromFirstDay()
}

func (h Week) Next() Week {
	return Week{Days: [7]Day{h.Days[6].Add(1)}}.fillFromFirstDay()
}

func (h Week) fillFromFirstDay() Week {
	for i := 1; i < 7; i++ {
		h.Days[i] = h.Days[i-1].Add(1)
	}

	return h
}

func (h Week) TailMonth() time.Month {
	return h.Days[6].Month()
}

func (h Week) HeadMonth() time.Month {
	return h.Days[0].Month()
}
