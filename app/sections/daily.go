package sections

import (
	"fmt"
	"strings"

	"github.com/kudrykv/latex-yearly-planner/app/calendar"
	"github.com/kudrykv/latex-yearly-planner/app/components"
	"github.com/kudrykv/latex-yearly-planner/app/types"
)

type SafeBuilder interface {
	Build() string
}

type SafeBuilderNoop struct{}

func (s SafeBuilderNoop) Build() string {
	return ""
}

type DailyParameters struct {
	ScheduleColumnWidth           types.Millimeters
	PrioritiesAndNotesColumnWidth types.Millimeters
	PrioritiesAndNotesSkip        types.Millimeters
	ColumnsSeparatorWidth         types.Millimeters
	ScheduleParameters            components.ScheduleParameters
	TodosParameters               components.TodosParameters
	NotesParameters               components.NotesParameters
	LittleCalendarParameters      components.LittleCalendarParameters
	ScheduleToTheRight            bool
	EnableLittleCalendar          bool
}

type Daily struct {
	day            calendar.Day
	parameters     DailyParameters
	schedule       components.Schedule
	todos          components.Todos
	notes          components.Notes
	littleCalendar SafeBuilder
	nearNotesLine  string
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

	notes, err := components.NewNotes(parameters.NotesParameters)
	if err != nil {
		return Daily{}, fmt.Errorf("new notes: %w", err)
	}

	littleCalendar := SafeBuilder(SafeBuilderNoop{})
	if parameters.EnableLittleCalendar {
		if littleCalendar, err = components.NewLittleCalendar(day, parameters.LittleCalendarParameters); err != nil {
			return Daily{}, fmt.Errorf("new little calendar: %w", err)
		}
	}

	return Daily{
		day:            day,
		parameters:     parameters,
		schedule:       schedule,
		todos:          todos,
		notes:          notes,
		littleCalendar: littleCalendar,
	}, nil
}

func (r Daily) AppendNearNotesLine(nearNotesLine string) Daily {
	r.nearNotesLine += nearNotesLine

	return r
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
		r.littleCalendar.Build(),
	)
}

func (r Daily) prioritiesAndNotesColumn() string {
	return fmt.Sprintf(
		prioritiesAndNotesColumnFormat,
		r.parameters.PrioritiesAndNotesColumnWidth,
		r.todos.Build(),
		r.parameters.PrioritiesAndNotesSkip,
		r.nearNotesLine,
		r.notes.Build(),
	)
}

const scheduleColumnFormat = `\begin{minipage}[t]{%s}
\myUnderline{Schedule\textcolor{white}{g}}
%s
%s
\end{minipage}`

const prioritiesAndNotesColumnFormat = `\begin{minipage}[t]{%s}
\myUnderline{Top Priorities}
%s
\vskip%s\myUnderline{Notes %s}
%s
\end{minipage}`

func (r Daily) Reference() string {
	return r.day.Format("2006-01-02")
}

func (r Daily) Title() string {
	return r.day.Format("Monday, 2")
}
