package components

import (
	"strings"

	"github.com/kudrykv/latex-yearly-planner/lib/calendar"
)

type FullPageCalendarParameters struct {
	WeekNumberToTheRight bool
}

type FullPageCalendar struct {
	month      calendar.Month
	parameters FullPageCalendarParameters
}

func NewFullPageCalendar(month calendar.Month, parameters FullPageCalendarParameters) FullPageCalendar {
	return FullPageCalendar{
		month:      month,
		parameters: parameters,
	}
}

func (r FullPageCalendar) Build() string {
	return `\begin{tabularx}{\linewidth}{` + r.tableRule() + `}
` + r.weekdays() + ` \\ \hline
\end{tabularx}`
}

func (r FullPageCalendar) weekdays() string {
	var weekdays []string

	for _, weekday := range r.month.Weekdays() {
		weekdays = append(weekdays, weekday.String())
	}

	if r.parameters.WeekNumberToTheRight {
		weekdays = append(weekdays, "W")
	} else {
		weekdays = append([]string{"W"}, weekdays...)
	}

	return strings.Join(weekdays, " & ")
}

func (r FullPageCalendar) tableRule() string {
	if r.parameters.WeekNumberToTheRight {
		return `*{7}{|@{}Y@{}}|c|`
	}

	return "|c|*{7}{@{}Y@{}|}"
}
