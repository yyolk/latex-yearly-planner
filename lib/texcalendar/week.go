package texcalendar

import "github.com/kudrykv/latex-yearly-planner/lib/calendar"

type Weeks struct {
	weeks calendar.Weeks
}

func NewWeek(weeks calendar.Weeks) Weeks {
	return Weeks{weeks: weeks}
}
