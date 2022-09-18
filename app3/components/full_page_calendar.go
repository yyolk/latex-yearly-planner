package components

import "github.com/kudrykv/latex-yearly-planner/lib/calendar"

type FullPageCalendarParameters struct {
}

type FullPageCalendar struct {
	month      calendar.Month
	parameters FullPageCalendarParameters
}

func NewFullPageCalendar(month calendar.Month, parameters FullPageCalendarParameters) FullPageCalendar {
	return FullPageCalendar{
		month:      month,
		parameters: parameters,
	}
}

func (r FullPageCalendar) Build() string {
	return r.month.Month().String()
}
