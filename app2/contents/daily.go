package contents

import (
	"bytes"
	"fmt"

	"github.com/kudrykv/latex-yearly-planner/app2/tex/components"
	"github.com/kudrykv/latex-yearly-planner/app2/tex/ref"
	"github.com/kudrykv/latex-yearly-planner/app2/types"
	"github.com/kudrykv/latex-yearly-planner/lib/texcalendar"
)

type Daily struct {
	day texcalendar.Day

	schedule types.Schedule

	showCalendar bool

	nearNotesLine string
	todosNumber   int
	notesNumber   int
}

type DailyOption func(*Daily)

func NewDaily(options ...DailyOption) Daily {
	daily := Daily{}

	for _, option := range options {
		option(&daily)
	}

	return daily
}

func DailyWithDay(day texcalendar.Day) DailyOption {
	return func(r *Daily) {
		r.day = day
	}
}

func DailyWithSchedule(schedule types.Schedule) DailyOption {
	return func(r *Daily) {
		r.schedule = schedule
	}
}

func DailyWithCalendar(show bool) DailyOption {
	return func(r *Daily) {
		r.showCalendar = show
	}
}

func DailyWithPlanning(todosNumber, notesNumber int) DailyOption {
	return func(r *Daily) {
		r.todosNumber = todosNumber
		r.notesNumber = notesNumber
	}
}

func DailyWithNearNotesLine(line string) DailyOption {
	return func(r *Daily) {
		r.nearNotesLine = line
	}
}

func (r Daily) Build() ([][]byte, error) {
	buff := &bytes.Buffer{}

	buff.WriteString(`\noindent\vskip1mm`)
	buff.WriteString(r.scheduleColumn())
	buff.WriteString(`\hspace{\myLengthThreeColumnsSeparatorWidth}`)
	buff.WriteString(r.prioritiesAndNotesColumn())

	return [][]byte{buff.Bytes()}, nil
}

func (r Daily) scheduleColumn() string {
	return fmt.Sprintf(scheduleColumnFormat, r.hoursSchedule(), r.optionalCalendar())
}

func (r Daily) hoursSchedule() string {
	return components.NewSchedule(r.schedule.FromHour, r.schedule.ToHour, r.schedule.HourFormat).Build()
}

func (r Daily) optionalCalendar() string {
	if r.showCalendar {
		return `\vspace{5mm}` + r.day.CalendarMonth().Selected(r.day).LittleCalendar()
	}

	return ""
}

func (r Daily) prioritiesAndNotesColumn() string {
	return fmt.Sprintf(
		prioritiesAndNotesColumnFormat,
		components.NewTodos(r.todosNumber).Build(),
		ref.NewNote("More", r.day.Ref()).Build(),
		ref.NewReflect("Reflect", r.day.Ref()).Build(),
		components.NewMesh(30, 19).Build(),
	)
}

const scheduleColumnFormat = `\begin{minipage}[t]{\myLengthThreeColumnWidth}
\myUnderline{Schedule\textcolor{white}{g}}
%s
%s
\end{minipage}`

const prioritiesAndNotesColumnFormat = `\begin{minipage}[t]{\dimexpr2\myLengthThreeColumnWidth+\myLengthThreeColumnsSeparatorWidth}
\myUnderline{Top Priorities}
%s
\vskip7mm\myUnderline{Notes | %s \hfill{}%s}
%s
\end{minipage}`
