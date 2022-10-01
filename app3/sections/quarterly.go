package sections

import "github.com/kudrykv/latex-yearly-planner/lib/calendar"

type QuarterlyParameters struct {
}

type Quarterly struct {
	parameters QuarterlyParameters
	quarter    calendar.Quarter
}

func NewQuarterly(quarter calendar.Quarter, parameters QuarterlyParameters) (Quarterly, error) {
	return Quarterly{
		quarter:    quarter,
		parameters: parameters,
	}, nil
}

func (r Quarterly) Title() string {
	return r.quarter.Name()
}

func (r Quarterly) Build() ([]string, error) {
	return []string{r.quarter.Name()}, nil
}
