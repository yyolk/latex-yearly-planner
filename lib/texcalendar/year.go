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
