package calendar

import (
	"time"
)

type Month struct{ Weeks Weeks }

func NewMonth(year int, mo time.Month, wd time.Weekday) Month {
	week := NewWeek(FromMonth(year, mo, wd))
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

	for _, day := range m.Weeks[1].Days {
		weekdays = append(weekdays, day.Weekday())
	}

	return weekdays
}
