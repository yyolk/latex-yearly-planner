package calendar

import (
	"time"
)

type Weeks []Week

type Week struct {
	days  [daysInWeek]Day
	first bool
	last  bool
}

type WeekOption func() Week

func FromDay(day Day) WeekOption {
	return func() Week {
		return Week{days: [daysInWeek]Day{day}}.fillFromFirstDay()
	}
}

func FromTime(moment time.Time) WeekOption {
	return func() Week {
		return Week{days: [daysInWeek]Day{{moment: moment}}}.fillFromFirstDay()
	}
}

func FromWeek(days [daysInWeek]Day) WeekOption {
	return func() Week {
		return Week{days: days}
	}
}

func FromMonth(year int, mo time.Month, wd time.Weekday) WeekOption {
	return func() Week {
		day := Day{moment: time.Date(year, mo, 1, 0, 0, 0, 0, time.Local)}
		week := Week{} //nolint:exhaustivestruct

		pos := (day.Weekday() - wd + daysInWeek) % daysInWeek
		week.days[pos] = day

		return week.fillFromDay(int(pos))
	}
}

func NewWeek(wo WeekOption) Week {
	return wo()
}

func (h Week) Next() Week {
	return Week{days: [daysInWeek]Day{h.days[6].Add(1)}}.fillFromFirstDay()
}

func (h Week) fillFromFirstDay() Week {
	return h.fillFromDay(0)
}

func (h Week) fillFromDay(n int) Week {
	for i := n + 1; i < daysInWeek; i++ {
		h.days[i] = h.days[i-1].Add(1)
	}

	return h
}

func (h Week) HeadMonth() time.Month {
	for _, day := range h.days {
		if day.IsZero() {
			continue
		}

		return day.Month()
	}

	return -1
}

func (h Week) TailMonth() time.Month {
	return h.days[6].Month()
}

func (h Week) ZerofyMonth(mo time.Month) Week {
	for i, day := range h.days {
		if day.Month() == mo {
			h.days[i] = Day{} //nolint:exhaustivestruct
		}
	}

	return h
}

func (h Week) WeekNumber() int {
	_, weekNumber := h.days[0].ISOWeek()

	for _, day := range h.days {
		if _, currDayWeekNumber := day.ISOWeek(); !day.IsZero() && currDayWeekNumber != weekNumber {
			return currDayWeekNumber
		}
	}

	return weekNumber
}

func (h Week) backfill() Week {
	for i := range h.days {
		if h.days[i].IsZero() {
			continue
		}

		for j := i - 1; j >= 0; j-- {
			h.days[j] = h.days[j+1].Add(-1)
		}

		break
	}

	for i := 6; i >= 0; i-- {
		if h.days[i].IsZero() {
			continue
		}

		for j := i + 1; j < 7; j++ {
			h.days[j] = h.days[j-1].Add(1)
		}

		break
	}

	return h
}

func (h Week) TailYear() int {
	return h.days[6].Year()
}

func (h Week) HeadYear() int {
	return h.days[0].Year()
}

func (h Week) SetFirst() Week {
	h.first = true

	return h
}

func (h Week) First() bool {
	return h.first
}

func (h Week) SetLast() Week {
	h.last = true

	return h
}

func (h Week) Last() bool {
	return h.last
}

func (h Week) Days() Days {
	return h.days[:]
}
