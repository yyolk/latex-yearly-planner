package calendar

import (
	"time"
)

type Months []Month

type Month struct {
	year    *Year
	quarter *Quarter

	Weeks Weeks
}

func newMonth(year *Year, quarter *Quarter, month time.Month, weekday time.Weekday) Month {
	calendarMonth := NewMonth(year.Year(), month, weekday)

	calendarMonth.year = year
	calendarMonth.quarter = quarter

	return calendarMonth
}

func NewMonth(year int, mo time.Month, wd time.Weekday) Month {
	week := NewWeek(FromMonth(year, mo, wd))
	if mo == time.January {
		week = week.SetFirst()
	}

	currMo := week.TailMonth()
	month := Month{Weeks: append(make(Weeks, 0, usualWeeksInMonth), week)}

	for {
		week = week.Next()

		if week.HeadMonth() != currMo {
			break
		}

		stop := false

		if tailMo := week.TailMonth(); tailMo != currMo {
			week = week.ZerofyMonth(tailMo)
			stop = true
		}

		month.Weeks = append(month.Weeks, week)

		if stop {
			break
		}
	}

	return month
}

func (m Month) Month() time.Month {
	if len(m.Weeks) == 0 {
		return -1
	}

	return m.Weeks[0].HeadMonth()
}

func (m Month) Weekdays() []time.Weekday {
	weekdays := make([]time.Weekday, 0, 7)

	for _, day := range m.Weeks[1].days {
		weekdays = append(weekdays, day.Weekday())
	}

	return weekdays
}
