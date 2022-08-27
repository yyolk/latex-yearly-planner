package sections

import (
	"fmt"

	"github.com/kudrykv/latex-yearly-planner/app3/components"
	"github.com/kudrykv/latex-yearly-planner/app3/types"
	"github.com/kudrykv/latex-yearly-planner/lib/calendar"
)

type DailyParameters struct {
	ScheduleColumnWidth types.Millimeters
	ScheduleParameters  components.ScheduleParameters
}

type Daily struct {
	day        calendar.Day
	parameters DailyParameters
	schedule   components.Schedule
}

func NewDaily(day calendar.Day, parameters DailyParameters) (Daily, error) {
	schedule, err := components.NewSchedule(day.Moment(), parameters.ScheduleParameters)
	if err != nil {
		return Daily{}, fmt.Errorf("new schedule: %w", err)
	}

	return Daily{
		day:        day,
		parameters: parameters,
		schedule:   schedule,
	}, nil
}

func (r Daily) Build() ([]string, error) {
	return []string{r.scheduleColumn()}, nil
}

func (r Daily) scheduleColumn() string {
	return fmt.Sprintf(
		scheduleColumnFormat,
		r.parameters.ScheduleColumnWidth,
		r.schedule.Build(),
	)
}

const scheduleColumnFormat = `\begin{minipage}[t]{%s}
\myUnderline{Schedule\textcolor{white}{g}}
%s
\end{minipage}`
