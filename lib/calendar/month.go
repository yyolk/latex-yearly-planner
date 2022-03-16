package calendar

type Month struct{ Weeks Weeks }

type MonthOption func() Week

func NewMonth(mo MonthOption) Month {
	week := mo()
	currMo := week.TailMonth()
	month := Month{Weeks: append(make(Weeks, 0, 5), week)}

	for {
		week = week.Next()

		if week.HeadMonth() == currMo {
			month.Weeks = append(month.Weeks, week)
		}

		if week.TailMonth() != currMo {
			break
		}
	}

	return month
}
