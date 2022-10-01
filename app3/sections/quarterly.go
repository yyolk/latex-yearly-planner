package sections

import (
	"fmt"

	"github.com/kudrykv/latex-yearly-planner/app3/components"
	"github.com/kudrykv/latex-yearly-planner/lib/calendar"
)

type QuarterlyParameters struct {
	LittleCalendarParameters components.LittleCalendarParameters
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
	mon1 := components.NewLittleCalendarFromMonth(r.quarter.Months[0], r.parameters.LittleCalendarParameters)
	mon2 := components.NewLittleCalendarFromMonth(r.quarter.Months[1], r.parameters.LittleCalendarParameters)
	mon3 := components.NewLittleCalendarFromMonth(r.quarter.Months[2], r.parameters.LittleCalendarParameters)

	return []string{fmt.Sprintf(
		quarterlyTemplate,
		mon1.Build(),
		mon2.Build(),
		mon3.Build(),
	)}, nil
}

const quarterlyTemplate = `\begin{minipage}[t][\remainingHeight]{5cm}
%s
\vfill
%s
\vfill
%s
\end{minipage}\hspace{5mm}\begin{minipage}[t][\remainingHeight]{5cm}
hello world
\end{minipage}`
