package calendar

import (
	"time"
)

type Weeks []Week

type Week struct {
	Days  [daysInWeek]Day
	first bool
	last  bool
}

type WeekOption func() Week

func FromDay(day Day) WeekOption {
	return func() Week {
		return Week{Days: [daysInWeek]Day{day}}.fillFromFirstDay()
	}
}

func FromTime(moment time.Time) WeekOption {
	return func() Week {
		return Week{Days: [daysInWeek]Day{{moment: moment}}}.fillFromFirstDay()
	}
}

func FromMonth(year int, mo time.Month, wd time.Weekday) WeekOption {
	return func() Week {
		day := Day{moment: time.Date(year, mo, 1, 0, 0, 0, 0, time.Local)}
		week := Week{} //nolint:exhaustivestruct

		pos := (day.Weekday() - wd + daysInWeek) % daysInWeek
		week.Days[pos] = day

		return week.fillFromDay(int(pos))
	}
}

func NewWeek(wo WeekOption) Week {
	return wo()
}

func (h Week) Next() Week {
	return Week{Days: [daysInWeek]Day{h.Days[6].Add(1)}}.fillFromFirstDay()
}

func (h Week) fillFromFirstDay() Week {
	return h.fillFromDay(0)
}

func (h Week) fillFromDay(n int) Week {
	for i := n + 1; i < daysInWeek; i++ {
		h.Days[i] = h.Days[i-1].Add(1)
	}

	return h
}

func (h Week) HeadMonth() time.Month {
	for _, day := range h.Days {
		if day.IsZero() {
			continue
		}

		return day.Month()
	}

	return -1
}

func (h Week) TailMonth() time.Month {
	return h.Days[6].Month()
}

func (h Week) ZerofyMonth(mo time.Month) Week {
	for i, day := range h.Days {
		if day.Month() == mo {
			h.Days[i] = Day{} //nolint:exhaustivestruct
		}
	}

	return h
}

func (h Week) WeekNumber() int {
	_, weekNumber := h.Days[0].ISOWeek()

	for _, day := range h.Days {
		if _, currDayWeekNumber := day.ISOWeek(); !day.IsZero() && currDayWeekNumber != weekNumber {
			return currDayWeekNumber
		}
	}

	return weekNumber
}

func (h Week) backfill() Week {
	for i := range h.Days {
		if h.Days[i].IsZero() {
			continue
		}

		for j := i - 1; j >= 0; j-- {
			h.Days[j] = h.Days[j+1].Add(-1)
		}

		break
	}

	for i := 6; i >= 0; i-- {
		if h.Days[i].IsZero() {
			continue
		}

		for j := i + 1; j < 7; j++ {
			h.Days[j] = h.Days[j-1].Add(1)
		}

		break
	}

	return h
}

func (h Week) TailYear() int {
	return h.Days[6].Year()
}

func (h Week) HeadYear() int {
	return h.Days[0].Year()
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
