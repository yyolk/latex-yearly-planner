package texcalendar

import (
	"strings"

	"github.com/kudrykv/latex-yearly-planner/lib/calendar"
)

type Month struct {
	month calendar.Month
}

func NewMonth(month calendar.Month) Month {
	return Month{month: month}
}

func (m Month) Tabular() string {
	monthName := m.month.Month().String()
	weekdays := strings.Join(append([]string{"W"}, m.weekdays()...), ` & `)
	weeksMatrix := m.tabulate(NewWeeks(m.month.Weeks).Matrix())

	return `\renewcommand{\arraystretch}{1.5}%` + "\n" +
		`%\setlength{\tabcolsep}{3.5pt}%` + "\n" +
		`\begin{tabularx}{\linewidth}[t]{c|*{7}{Y}}` + "\n" +
		`\multicolumn{8}{c}{` + monthName + `} \\ \hline` + "\n" +
		weekdays + `\\ \hline` + "\n" +
		weeksMatrix + "\n" +
		`\end{tabularx}`
}

func (m Month) weekdays() []string {
	weekdays := make([]string, 0, 7)
	for _, weekday := range m.month.Weekdays() {
		weekdays = append(weekdays, weekday.String()[:1])
	}
	return weekdays
}

func (m Month) tabulate(matrix [][]string) string {
	rows := make([]string, 0, len(matrix))

	for _, row := range matrix {
		rows = append(rows, strings.Join(row, " & "))
	}

	return strings.Join(rows, `\\`+"\n")
}

func (m Month) Tabloid() string {
	return m.month.Month().String()
}
