package contents

import (
	"bytes"
	"fmt"

	"github.com/kudrykv/latex-yearly-planner/app2/tex/components"
	"github.com/kudrykv/latex-yearly-planner/app2/tex/ref"
	"github.com/kudrykv/latex-yearly-planner/lib/texcalendar"
)

type Daily struct {
	day           texcalendar.Day
	nearNotesLine string

	parameters DailyParameters
}

type DailyParameters struct {
	Schedule Schedule

	ShowCalendar bool

	TodosNumber int
	NotesNumber int
}

type Schedule struct {
	FromHour   int
	ToHour     int
	HourFormat string
}

func NewDaily(day texcalendar.Day, parameters DailyParameters) Daily {
	return Daily{
		day:        day,
		parameters: parameters,
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
	return components.NewSchedule(
		r.parameters.Schedule.FromHour, r.parameters.Schedule.ToHour, r.parameters.Schedule.HourFormat,
	).Build()
}

func (r Daily) optionalCalendar() string {
	if r.parameters.ShowCalendar {
		return `\vspace{5mm}` + r.day.CalendarMonth().Selected(r.day).LittleCalendar()
	}

	return ""
}

func (r Daily) prioritiesAndNotesColumn() string {
	return fmt.Sprintf(
		prioritiesAndNotesColumnFormat,
		components.NewTodos(r.parameters.TodosNumber).Build(),
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
