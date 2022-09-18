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

func (r Week) Next() Week {
	return Week{days: [daysInWeek]Day{r.days[6].Add(1)}}.fillFromFirstDay()
}

func (r Week) fillFromFirstDay() Week {
	return r.fillFromDay(0)
}

func (r Week) fillFromDay(n int) Week {
	for i := n + 1; i < daysInWeek; i++ {
		r.days[i] = r.days[i-1].Add(1)
	}

	return r
}

func (r Week) HeadMonth() time.Month {
	for _, day := range r.days {
		if day.IsZero() {
			continue
		}

		return day.Month()
	}

	return -1
}

func (r Week) TailMonth() time.Month {
	return r.days[6].Month()
}

func (r Week) ZerofyMonth(mo time.Month) Week {
	for i, day := range r.days {
		if day.Month() == mo {
			r.days[i] = Day{} //nolint:exhaustivestruct
		}
	}

	return r
}

func (r Week) WeekNumber() int {
	_, weekNumber := r.days[0].ISOWeek()

	for _, day := range r.days {
		if _, currDayWeekNumber := day.ISOWeek(); !day.IsZero() && currDayWeekNumber != weekNumber {
			return currDayWeekNumber
		}
	}

	return weekNumber
}

func (r Week) backfill() Week {
	for i := range r.days {
		if r.days[i].IsZero() {
			continue
		}

		for j := i - 1; j >= 0; j-- {
			r.days[j] = r.days[j+1].Add(-1)
		}

		break
	}

	for i := 6; i >= 0; i-- {
		if r.days[i].IsZero() {
			continue
		}

		for j := i + 1; j < 7; j++ {
			r.days[j] = r.days[j-1].Add(1)
		}

		break
	}

	return r
}

func (r Week) TailYear() int {
	return r.days[6].Year()
}

func (r Week) HeadYear() int {
	return r.days[0].Year()
}

func (r Week) SetFirst() Week {
	r.first = true

	return r
}

func (r Week) First() bool {
	return r.first
}

func (r Week) SetLast() Week {
	r.last = true

	return r
}

func (r Week) Last() bool {
	return r.last
}

func (r Week) Days() Days {
	return r.days[:]
}

func (r Week) Quarters(year int, weekday time.Weekday) Quarters {
	quarters := make(Quarters, 0, 2)

	if r.First() {
		return append(quarters, NewQuarter(year, FirstQuarter, weekday))
	}

	if r.Last() {
		return append(quarters, NewQuarter(year, FourthQuarter, weekday))
	}

	return append(
		quarters,
		NewQuarter(year, GetQuarterFromMonth(r.HeadMonth()), weekday),
		NewQuarter(year, GetQuarterFromMonth(r.TailMonth()), weekday),
	)
}

func GetQuarterFromMonth(month time.Month) int {
	switch month {
	case time.January, time.February, time.March:
		return FirstQuarter
	case time.April, time.May, time.June:
		return SecondQuarter
	case time.July, time.August, time.September:
		return ThirdQuarter
	case time.October, time.November, time.December:
		return FourthQuarter
	}

	return -1
}

func (r Week) Months(year int, weekday time.Weekday) Months {
	months := make(Months, 0, 2)

	if r.First() {
		return append(months, NewMonth(year, time.January, weekday))
	}

	if r.Last() {
		return append(months, NewMonth(year, time.December, weekday))
	}

	return append(
		months,
		NewMonth(year, r.HeadMonth(), weekday),
		NewMonth(year, r.TailMonth(), weekday),
	)
}
