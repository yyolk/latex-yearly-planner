package components

import (
	"fmt"
	"strings"

	"github.com/kudrykv/latex-yearly-planner/app/calendar"
	"github.com/kudrykv/latex-yearly-planner/app/types"
)

type FullPageCalendarParameters struct {
	WeekNumberToTheRight bool
	LineHeight           types.Millimeters

	Dotted Dotted
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
` + r.matrix() + ` \\ \hline
\end{tabularx}
\vskip 5mm

\myUnderline{Notes}
  ` + r.notes()
}

func (r FullPageCalendar) notes() string {
	if r.parameters.LineHeight > 0 {
		return fmt.Sprintf(`\vbox to \dimexpr\textheight-\pagetotal\relax {%%
    \leaders\hbox to \linewidth{\textcolor{\myColorGray}{\rule{0pt}{%s}\hrulefill}}\vfil
}`, r.parameters.LineHeight)
	}

	return fmt.Sprintf(
		`\vbox to 0pt{\vskip%s\leavevmode\multido{\dC=0mm+%s}{%s}{\multido{\dR=0mm+%s}{%s}{\put(\dR,\dC){\scriptsize.}}}}`,
		r.parameters.Dotted.Distance,
		r.parameters.Dotted.Distance,
		r.parameters.Dotted.Rows,
		r.parameters.Dotted.Distance,
		r.parameters.Dotted.Columns,
	)
}

func (r FullPageCalendar) weekdays() string {
	var weekdays []string

	for _, weekday := range r.month.Weekdays() {
		weekdays = append(weekdays, `\hfil{}`+weekday.String())
	}

	if r.parameters.WeekNumberToTheRight {
		weekdays = append(weekdays, "W")
	} else {
		weekdays = append([]string{"W"}, weekdays...)
	}

	return strings.Join(weekdays, " & ")
}

func (r FullPageCalendar) matrix() string {
	pieces := make([]string, 0, len(r.month.Weeks))

	for _, week := range r.month.Weeks {
		weekRow := r.weekRow(week)

		pieces = append(pieces, strings.Join(weekRow, " & "))
	}

	return strings.Join(pieces, " \\\\ \\hline\n")
}

func (r FullPageCalendar) weekRow(week calendar.Week) []string {
	row := make([]string, 0, 8)

	for _, day := range week.Days() {
		if day.IsZero() {
			row = append(row, "")

			continue
		}

		value := fmt.Sprintf(`\hfill%s \parbox{0pt}{\vskip5mm}`, NewDayLink(day))

		row = append(row, value)
	}

	row = r.addWeekNumber(row, week)

	return row
}

func (r FullPageCalendar) addWeekNumber(row []string, week calendar.Week) []string {
	value := fmt.Sprintf(`\rotatebox[origin=tr]{90}{\makebox[1.85cm][c]{%s}}`, NewWeekLink(week).Format("Week %d"))

	if r.parameters.WeekNumberToTheRight {
		return append(row, value)
	}

	return append([]string{value}, row...)
}

func (r FullPageCalendar) tableRule() string {
	if r.parameters.WeekNumberToTheRight {
		return `*{7}{|@{}X@{}}|c|`
	}

	return "|c|*{7}{@{}X@{}|}"
}
