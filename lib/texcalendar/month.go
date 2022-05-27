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
	weekdays := make([]string, 0, 7)
	for _, weekday := range m.month.Weekdays() {
		weekdays = append(weekdays, weekday.String()[:1])
	}

	return `\renewcommand{\arraystretch}{1.5}%` + "\n" +
		`%\setlength{\tabcolsep}{3.5pt}%` + "\n" +
		`\begin{tabularx}{\linewidth}{Y*{7}{Y}}` + "\n" +
		`\multicolumn{8}{c}{` + m.month.Month().String() + `} \\ \hline` + "\n" +
		`\hline` + "\n" +
		strings.Join(append([]string{"W"}, weekdays...), ` & `) + `\\` + "\n" +
		m.month.TabularWeeks() + "\n" +
		`\end{tabularx}`
}
