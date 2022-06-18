package texcalendar

import (
	"strings"

	"github.com/kudrykv/latex-yearly-planner/app2/planners/common"
	"github.com/kudrykv/latex-yearly-planner/lib/calendar"
)

type Year struct {
	year calendar.Year
	hand common.MainHand
}

func NewYear(hand common.MainHand, year calendar.Year) Year {
	return Year{hand: hand, year: year}
}

func (r Year) BuildCalendar() string {
	quarterRows := make([]string, 0, len(r.year.Quarters))

	tabular := func(str string) string {
		return `\begin{tabularx}{\linewidth}{@{}*{3}{X}@{}}` + "\n" +
			str +
			"\n" + `\end{tabularx}`
	}

	for _, quarter := range r.year.Quarters {
		quarterRows = append(quarterRows, tabular(NewQuarter(r.hand, quarter).Row()))
	}

	return strings.Join(quarterRows, "\n"+`\vfill`+"\n")
}

func (r Year) Quarters() string {
	quarters := make([]string, 0, 4)

	for i := len(r.year.Quarters) - 1; i >= 0; i-- {
		quarters = append(quarters, r.year.Quarters[i].Name())
	}

	return `\begin{tabularx}{5cm}{*{4}{|Y}|}
	` + strings.Join(quarters, " & ") + ` \\ \hline
\end{tabularx}`
}

func (r Year) Months() string {
	months := make([]string, 0, 12)

	for i := len(r.year.Quarters) - 1; i >= 0; i-- {
		for j := 2; j >= 0; j-- {
			months = append(months, r.year.Quarters[i].Months[j].Month().String()[:3])
		}
	}

	return `\begin{tabularx}{15.5cm}{*{12}{|Y}|}
	` + strings.Join(months, " & ") + `\\ \hline
\end{tabularx}`
}
