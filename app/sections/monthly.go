package sections

import (
	"github.com/kudrykv/latex-yearly-planner/app/calendar"
	"github.com/kudrykv/latex-yearly-planner/app/components"
)

type MonthlyParameters struct {
	FullPageCalendarParameters components.FullPageCalendarParameters
}

type Monthly struct {
	month            calendar.Month
	parameters       MonthlyParameters
	fullPageCalendar components.FullPageCalendar
}

func NewMonthly(month calendar.Month, parameters MonthlyParameters) (Monthly, error) {
	fullPageCalendar := components.NewFullPageCalendar(month, parameters.FullPageCalendarParameters)

	return Monthly{
		month:      month,
		parameters: parameters,

		fullPageCalendar: fullPageCalendar,
	}, nil
}

func (r Monthly) Title() string {
	return r.month.Month().String()
}

func (r Monthly) Reference() string {
	return r.month.Month().String()
}

func (r Monthly) Build() ([]string, error) {
	return []string{r.fullPageCalendar.Build()}, nil
}
