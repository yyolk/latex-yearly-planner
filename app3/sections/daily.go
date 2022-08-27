package sections

import (
	"fmt"
	"strings"

	"github.com/kudrykv/latex-yearly-planner/app3/components"
	"github.com/kudrykv/latex-yearly-planner/app3/types"
	"github.com/kudrykv/latex-yearly-planner/lib/calendar"
)

type DailyParameters struct {
	ScheduleColumnWidth           types.Millimeters
	PrioritiesAndNotesColumnWidth types.Millimeters
	ScheduleParameters            components.ScheduleParameters
	TodosParameters               components.TodosParameters
	PrioritiesAndNotesSkip        types.Millimeters
	ColumnsSeparatorWidth         types.Millimeters
	ScheduleToTheRight            bool
}

type Daily struct {
	day        calendar.Day
	parameters DailyParameters
	schedule   components.Schedule
	todos      components.Todos
}

func NewDaily(day calendar.Day, parameters DailyParameters) (Daily, error) {
	schedule, err := components.NewSchedule(day.Moment(), parameters.ScheduleParameters)
	if err != nil {
		return Daily{}, fmt.Errorf("new schedule: %w", err)
	}

	todos, err := components.NewTodos(parameters.TodosParameters)
	if err != nil {
		return Daily{}, fmt.Errorf("new todos: %w", err)
	}

	return Daily{
		day:        day,
		parameters: parameters,
		schedule:   schedule,
		todos:      todos,
	}, nil
}

func (r Daily) Build() ([]string, error) {
	pieces := []string{
		r.scheduleColumn(),
		fmt.Sprintf(`\hspace{%s}`, r.parameters.ColumnsSeparatorWidth),
		r.prioritiesAndNotesColumn(),
	}

	if r.parameters.ScheduleToTheRight {
		pieces[0], pieces[2] = pieces[2], pieces[0]
	}

	return []string{strings.Join(pieces, "")}, nil
}

func (r Daily) scheduleColumn() string {
	return fmt.Sprintf(
		scheduleColumnFormat,
		r.parameters.ScheduleColumnWidth,
		r.schedule.Build(),
	)
}

func (r Daily) prioritiesAndNotesColumn() string {
	return fmt.Sprintf(
		prioritiesAndNotesColumnFormat,
		r.parameters.PrioritiesAndNotesColumnWidth,
		r.todos.Build(),
		r.parameters.PrioritiesAndNotesSkip,
	)
}

const scheduleColumnFormat = `\begin{minipage}[t]{%s}
\myUnderline{Schedule\textcolor{white}{g}}
%s
\end{minipage}`

const prioritiesAndNotesColumnFormat = `\begin{minipage}[t]{%s}
\myUnderline{Top Priorities}
%s
\vskip%s\myUnderline{Notes | later here \hfill{}later there}
notes box
\end{minipage}`
