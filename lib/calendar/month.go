package calendar

import "time"

type Month struct{ Weeks Weeks }

func NewMonth(year int, mo time.Month, wd time.Weekday) Month {
	week := NewWeek(FromMonth(year, mo, wd))
	currMo := week.TailMonth()
	month := Month{Weeks: append(make(Weeks, 0, 5), week)}

	for {
		week = week.Next()

		if week.HeadMonth() == currMo {
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
	}

	return month
}
