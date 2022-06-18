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
	weekdays := strings.Join(append([]string{"W"}, m.weekdaysShort()...), ` & `)
	weeksMatrix := m.tabulate(NewWeeks(m.month.Weeks).Matrix(), `\\`)

	return `\renewcommand{\arraystretch}{1.5}%` + "\n" +
		`%\setlength{\tabcolsep}{3.5pt}%` + "\n" +
		`\begin{tabularx}{\linewidth}[t]{c|*{7}{Y}}` + "\n" +
		`\multicolumn{8}{c}{` + monthName + `} \\ \hline` + "\n" +
		weekdays + `\\ \hline` + "\n" +
		weeksMatrix + "\n" +
		`\end{tabularx}`
}

func (m Month) ForPage() string {
	weekdays := strings.Join(append([]string{"W"}, m.weekdays()...), ` & `)
	weeksMatrix := m.tabulate(m.cornerDays(m.renameRotateWeek(NewWeeks(m.month.Weeks).Matrix())), `\\ \hline`)

	return `\renewcommand{\arraystretch}{0}%` + "\n" +
		`%\setlength{\tabcolsep}{0pt}%` + "\n" +
		`\begin{tabularx}{\linewidth}[t]{@{ }c@{ }|*{7}{@{}X@{}|}}` + "\n" +
		weekdays + `\\[3mm] \hline` + "\n" +
		weeksMatrix + "\\\\ \\hline\n" +
		`\end{tabularx}`
}

func (m Month) weekdaysShort() []string {
	weekdays := make([]string, 0, 7)

	for _, weekday := range m.month.Weekdays() {
		weekdays = append(weekdays, weekday.String()[:1])
	}

	return weekdays
}

func (m Month) weekdays() []string {
	weekdays := make([]string, 0, 7)

	for _, weekday := range m.month.Weekdays() {
		weekdays = append(weekdays, `\hfil{}`+weekday.String())
	}

	return weekdays
}

func (m Month) tabulate(matrix [][]string, join string) string {
	rows := make([]string, 0, len(matrix))

	for _, row := range matrix {
		rows = append(rows, strings.Join(row, " & "))
	}

	return strings.Join(rows, join+"\n")
}

func (m Month) Tabloid() string {
	return m.month.Month().String()
}

func (m Month) renameRotateWeek(matrix [][]string) [][]string {
	for i := range matrix {
		matrix[i][0] = `\rotatebox[origin=tr]{90}{\makebox[2cm][c]{` + "Week " + matrix[i][0] + `}}`
	}

	return matrix
}

func (m Month) cornerDays(matrix [][]string) [][]string {
	for _, week := range matrix {
		for i := 1; i < len(week); i++ {
			if len(week[i]) == 0 {
				continue
			}

			week[i] = `{\renewcommand{\arraystretch}{1.2}\begin{tabular}{@{}p{5mm}@{}|}\hfil{}` + week[i] + `\\ \hline\end{tabular}}`
		}
	}

	return matrix
}
