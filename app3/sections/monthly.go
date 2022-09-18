package sections

import "github.com/kudrykv/latex-yearly-planner/lib/calendar"

type MonthlyParameters struct {
}

type Monthly struct {
	month      calendar.Month
	parameters MonthlyParameters
}

func NewMonthly(month calendar.Month, parameters MonthlyParameters) (Monthly, error) {
	return Monthly{
		month:      month,
		parameters: parameters,
	}, nil
}

func (m Monthly) Build() ([]string, error) {
	return []string{"hello world"}, nil
}
