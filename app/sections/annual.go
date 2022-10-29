package sections

import (
	"fmt"
	"strings"

	"github.com/kudrykv/latex-yearly-planner/app/calendar"
	"github.com/kudrykv/latex-yearly-planner/app/components"
	"github.com/kudrykv/latex-yearly-planner/app/types"
)

type AnnualParameters struct {
	ColumnWidth          types.Millimeters
	ColumnSeparatorWidth types.Millimeters

	LittleCalendarParameters components.LittleCalendarParameters
}

type Annual struct {
	parameters AnnualParameters
	year       calendar.Year
}

func NewAnnual(year calendar.Year, parameters AnnualParameters) (Annual, error) {
	return Annual{
		year:       year,
		parameters: parameters,
	}, nil
}

func (r Annual) Build() ([]string, error) {
	calendars := make([]components.LittleCalendar, 0, 12)

	for _, month := range r.year.Months() {
		calendars = append(calendars, components.NewLittleCalendarFromMonth(month, r.parameters.LittleCalendarParameters))
	}

	return []string{
		fmt.Sprintf(r.template(),
			calendars[0].Build(), calendars[1].Build(), calendars[2].Build(),
			calendars[3].Build(), calendars[4].Build(), calendars[5].Build(),
			calendars[6].Build(), calendars[7].Build(), calendars[8].Build(),
			calendars[9].Build(), calendars[10].Build(), calendars[11].Build(),
		),
	}, nil
}

func (r Annual) template() string {
	row := fmt.Sprintf(`\begin{minipage}[t]{%s}%%s\end{minipage}\hspace{%s}%%%%
\begin{minipage}[t]{%s}%%s\end{minipage}\hspace{%s}%%%%
\begin{minipage}[t]{%s}%%s\end{minipage}`, r.parameters.ColumnWidth, r.parameters.ColumnSeparatorWidth, r.parameters.ColumnWidth, r.parameters.ColumnSeparatorWidth, r.parameters.ColumnWidth)

	return strings.Join([]string{row, row, row, row}, "\n\\vfill\n")
}

func (r Annual) Title() string {
	return fmt.Sprintf("%d", r.year.Year())
}

func (r Annual) Reference() string {
	return "Calendar"
}
