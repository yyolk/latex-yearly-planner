package components

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/kudrykv/latex-yearly-planner/lib/calendar"
)

type LittleCalendarParameters struct {
	WeekNumberToTheRight bool
	WeekHighlight        bool
}

type LittleCalendar struct {
	today      calendar.Day
	month      calendar.Month
	parameters LittleCalendarParameters
}

var ErrNoMonth = errors.New("no month")

func NewLittleCalendar(today calendar.Day, parameters LittleCalendarParameters) (LittleCalendar, error) {
	month := today.CalendarMonth()
	if month == nil {
		return LittleCalendar{}, fmt.Errorf("day doesn't return its month: %w", ErrNoMonth)
	}

	return LittleCalendar{
		today:      today,
		month:      *month,
		parameters: parameters,
	}, nil
}

func NewLittleCalendarFromMonth(month calendar.Month, parameters LittleCalendarParameters) LittleCalendar {
	return LittleCalendar{
		month:      month,
		parameters: parameters,
	}
}

func (r LittleCalendar) Build() string {
	heading := strings.Join(r.centerItems(r.weekNumberInHeader(r.weeks())), " & ")
	return fmt.Sprintf(`\setlength{\tabcolsep}{0pt}%%
\renewcommand{\arraystretch}{1.5}%%
\begin{tabularx}{\linewidth}[t]{%s}
	\multicolumn{8}{c}{%s} \\
	%s \\ \hline
	%s
\end{tabularx}
`,
		r.tableRule(),
		r.month.Month().String(),
		heading,
		r.matrix(),
	)
}

func (r LittleCalendar) centerItems(row []string) []string {
	for i, item := range row {
		row[i] = `\hfil{}` + item
	}

	return row
}

func (r LittleCalendar) weekNumberInHeader(weekdays []string) []string {
	if r.parameters.WeekNumberToTheRight {
		return append(weekdays, "W")
	}

	return append([]string{"W"}, weekdays...)
}

func (r LittleCalendar) weeks() []string {
	weekdays := make([]string, 0, 8)

	for _, weekday := range r.month.Weekdays() {
		weekdays = append(weekdays, weekday.String()[:1])
	}

	return weekdays
}

func (r LittleCalendar) matrix() string {
	pieces := make([]string, 0, len(r.month.Weeks))

	for _, week := range r.month.Weeks {
		weekRow := r.weekRow(week)
		highlight := ""

		if r.parameters.WeekHighlight && r.today.Week().WeekNumber() == week.WeekNumber() {
			highlight = `\rowcolor{lightgray}`
		}

		pieces = append(pieces, highlight+strings.Join(weekRow, " & "))
	}

	return strings.Join(pieces, " \\\\ \n")
}

func (r LittleCalendar) weekRow(week calendar.Week) []string {
	row := make([]string, 0, 8)

	for _, day := range week.Days() {
		if day.IsZero() {
			row = append(row, "")

			continue
		}

		if r.today.Day() == day.Day() {
			row = append(row, fmt.Sprintf(`\cellcolor{black}{\textcolor{white}{%d}}`, day.Day()))

			continue
		}

		row = append(row, strconv.Itoa(day.Day()))
	}

	row = r.addWeekNumber(row, week.WeekNumber())

	return row
}

func (r LittleCalendar) addWeekNumber(row []string, number int) []string {
	if r.parameters.WeekNumberToTheRight {
		return append(row, strconv.Itoa(number))
	}

	return append([]string{strconv.Itoa(number)}, row...)
}

func (r LittleCalendar) tableRule() string {
	if r.parameters.WeekNumberToTheRight {
		return `*{7}{@{}Y@{}}|Y`
	}

	return "Y|*{7}{@{}Y@{}}"
}
