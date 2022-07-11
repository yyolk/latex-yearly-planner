package mos

import (
	"fmt"

	"github.com/kudrykv/latex-yearly-planner/app2/planners/common"
	"github.com/kudrykv/latex-yearly-planner/app2/tex/components"
	"github.com/kudrykv/latex-yearly-planner/app2/tex/ref"
	"github.com/kudrykv/latex-yearly-planner/lib/texcalendar"
)

type dailyContents struct {
	day  texcalendar.Day
	hand common.MainHand
	ui   ui
}

func (r dailyContents) Build() ([]string, error) {
	leftColumn := r.scheduleColumn()
	rightColumn := r.prioritiesAndNotesColumn()

	if r.hand == common.LeftHand {
		leftColumn, rightColumn = rightColumn, leftColumn
	}

	return []string{
		`\noindent\vskip1mm` + leftColumn + `\hspace{\myLengthThreeColumnsSeparatorWidth}` + rightColumn,
	}, nil
}

func (r dailyContents) prioritiesAndNotesColumn() string {
	return fmt.Sprintf(
		prioritiesAndNotesColumnFormat,
		r.newTodos(),
		ref.NewNote("More", r.day.Ref()).Build(),
		ref.NewReflect("Reflect", r.day.Ref()).Build(),
		components.NewMesh(30, 19).Build(),
	)
}

func (r dailyContents) newTodos() string {
	return components.NewTodos(r.ui.TodosNumber).Build()
}

func (r dailyContents) scheduleColumn() string {
	return fmt.Sprintf(scheduleColumnFormat, r.hoursSchedule(), r.optionalCalendar())
}

func (r dailyContents) hoursSchedule() string {
	return components.NewSchedule(r.ui.FromScheduleHour, r.ui.ToScheduleHour, r.ui.HourFormat).Build()
}

func (r dailyContents) optionalCalendar() string {
	if !r.ui.EnableCalendarOnDailyPages {
		return ""
	}

	return r.day.CalendarMonth().Selected(r.day).LittleCalendar()
}

const prioritiesAndNotesColumnFormat = `\begin{minipage}[t]{\dimexpr2\myLengthThreeColumnWidth+\myLengthThreeColumnsSeparatorWidth}
\myUnderline{Top Priorities}
%s
\vskip7mm\myUnderline{Notes | %s \hfill{}%s}
%s
\end{minipage}`

const scheduleColumnFormat = `\begin{minipage}[t]{\myLengthThreeColumnWidth}
\myUnderline{Schedule\textcolor{white}{g}}
%s
%s
\end{minipage}`
