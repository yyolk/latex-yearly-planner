package texcalendar

import (
	"strings"

	"github.com/kudrykv/latex-yearly-planner/app2/planners/common"
	"github.com/kudrykv/latex-yearly-planner/lib/calendar"
)

type Month struct {
	month calendar.Month
	hand  common.MainHand
}

func NewMonth(month calendar.Month, hand common.MainHand) Month {
	return Month{month: month, hand: hand}
}

func (m Month) Tabular() string {
	monthName := m.month.Month().String()
	weekdays := strings.Join(append([]string{"W"}, m.weekdaysShort()...), ` & `)
	weeksMatrix := m.tabulate(NewWeeks(common.RightHand, m.month.Weeks).Matrix(), `\\`)

	return `\renewcommand{\arraystretch}{1.5}%` + "\n" +
		`%\setlength{\tabcolsep}{3.5pt}%` + "\n" +
		`\begin{tabularx}{\linewidth}[t]{c|*{7}{Y}}` + "\n" +
		`\multicolumn{8}{c}{` + monthName + `} \\ \hline` + "\n" +
		weekdays + `\\ \hline` + "\n" +
		weeksMatrix + "\n" +
		`\end{tabularx}`
}

func (m Month) ForPage() string {
	weekdaysSlice := m.weekdays()
	if m.hand == common.RightHand {
		weekdaysSlice = append([]string{"W"}, weekdaysSlice...)
	} else {
		weekdaysSlice = append(weekdaysSlice, "W")
	}

	weekdays := strings.Join(weekdaysSlice, ` & `)
	weeksMatrix := m.tabulate(m.cornerDays(m.renameRotateWeek(NewWeeks(m.hand, m.month.Weeks).Matrix())), `\\ \hline`)

	tableRule := `|@{ }c@{ }|*{7}{@{}X@{}|}`
	if m.hand == common.LeftHand {
		tableRule = `*{7}{|@{}X@{}}|@{ }c@{ }|`
	}

	return `\renewcommand{\arraystretch}{0}%` + "\n" +
		`%\setlength{\tabcolsep}{0pt}%` + "\n" +
		`\begin{tabularx}{\linewidth}[t]{` + tableRule + `}` + "\n" +
		weekdays + ` \raisebox{4mm}{} \\[2mm] \hline` + "\n" +
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
	idx := 0
	if m.hand == common.LeftHand {
		idx = 7
	}

	for i := range matrix {
		matrix[i][idx] = `\rotatebox[origin=tr]{90}{\makebox[2cm][c]{` + "Week " + matrix[i][idx] + `}}`
	}

	return matrix
}

func (m Month) cornerDays(matrix [][]string) [][]string {
	left, right := 1, len(matrix[0])
	if m.hand == common.LeftHand {
		left, right = left-1, right-1
	}

	for _, week := range matrix {
		for i := left; i < right; i++ {
			if len(week[i]) == 0 {
				continue
			}

			week[i] = `{\renewcommand{\arraystretch}{1.2}\begin{tabular}{@{}p{5mm}@{}|}\hfil{}` + week[i] + `\\ \hline\end{tabular}}`
		}
	}

	return matrix
}
