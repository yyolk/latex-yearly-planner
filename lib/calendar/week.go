package calendar

type Weeks []Week

type Week struct {
	Days [7]Day
}

type WeekOption func() Day

func NewWeek(wo WeekOption) Week {
	week := Week{}
	week.Days[0] = wo()

	for i := 1; i < 7; i++ {
		week.Days[i] = week.Days[i-1].Add(1)
	}

	return week
}
