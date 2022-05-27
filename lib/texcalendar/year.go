package texcalendar

import (
	"strings"

	"github.com/kudrykv/latex-yearly-planner/lib/calendar"
)

type Year struct {
	year calendar.Year
}

func NewYear(year calendar.Year) Year {
	return Year{year: year}
}

func (r Year) BuildCalendar() string {
	quarterRows := make([]string, 0, len(r.year.Quarters))

	for _, quarter := range r.year.Quarters {
		quarterRows = append(quarterRows, NewQuarter(quarter).Row())
	}

	return strings.Join(quarterRows, "\n"+`\vfill`+"\n")
}
