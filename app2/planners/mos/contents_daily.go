package mos

import (
	"fmt"
	"strings"
	"time"

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
	)
}

func (r dailyContents) newTodos() string {
	return components.NewTodos(r.ui.TodosNumber).Build()
}

func (r dailyContents) scheduleColumn() string {
	hoursSchedule := r.hoursSchedule()

	return fmt.Sprintf(scheduleColumnFormat, hoursSchedule, r.optionalCalendar())
}

func (r dailyContents) hoursSchedule() string {
	var schedule []string

	for i := r.ui.FromScheduleHour; i <= r.ui.ToScheduleHour; i++ {
		formattedHour := time.Date(0, 0, 0, i, 0, 0, 0, time.Local).Format(r.ui.HourFormat)
		schedule = append(schedule, r.scheduleHour(formattedHour))
	}

	return strings.Join(schedule, "\n")
}

func (r dailyContents) scheduleHour(strHour string) string {
	return r.height() + strHour + `\myLineLightGray
\vskip5mm\myLineGray`
}

func (r dailyContents) height() string {
	return `\parbox{0pt}{\vskip5mm}`
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
\vspace{5mm}\hspace{.5mm}\vbox to 0mm{\myDotGrid{30}{19}}
\end{minipage}`

const scheduleColumnFormat = `\begin{minipage}[t]{\myLengthThreeColumnWidth}
\myUnderline{Schedule\textcolor{white}{g}}
%s
%s
\end{minipage}`
