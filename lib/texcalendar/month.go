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

func (m Month) LittleCalendar() string {
	monthName := m.month.Month().String()
	weekdays := strings.Join(append([]string{"W"}, m.weekdaysShort()...), ` & `)
	weeksMatrix := m.tabulate(NewWeeks(common.RightHand, m.month.Weeks, false).Matrix(), `\\`)

	return `\renewcommand{\arraystretch}{1.5}%` + "\n" +
		`%\setlength{\tabcolsep}{3.5pt}%` + "\n" +
		`\begin{tabularx}{\linewidth}[t]{c|*{7}{Y}}` + "\n" +
		`\multicolumn{8}{c}{` + monthName + `} \\ \hline` + "\n" +
		weekdays + `\\ \hline` + "\n" +
		weeksMatrix + "\n" +
		`\end{tabularx}`
}

func (m Month) LargeCalendar() string {
	weeks := NewWeeks(m.hand, m.month.Weeks, true)
	weekdays := strings.Join(weeks.Weekdays(), ` & `)
	weeksMatrix := m.tabulate(weeks.Matrix(), `\\ \hline`)

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

func (m Month) tabulate(matrix [][]string, join string) string {
	rows := make([]string, 0, len(matrix))

	for _, row := range matrix {
		rows = append(rows, strings.Join(row, " & "))
	}

	return strings.Join(rows, join+"\n")
}
