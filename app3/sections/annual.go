package sections

import (
	"fmt"

	"github.com/kudrykv/latex-yearly-planner/app3/components"
	"github.com/kudrykv/latex-yearly-planner/lib/calendar"
)

type AnnualParameters struct {
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
		fmt.Sprintf(annualTemplate,
			calendars[0].Build(), calendars[1].Build(), calendars[2].Build(),
			calendars[3].Build(), calendars[4].Build(), calendars[5].Build(),
			calendars[6].Build(), calendars[7].Build(), calendars[8].Build(),
			calendars[9].Build(), calendars[10].Build(), calendars[11].Build(),
		),
	}, nil
}

const annualTemplate = `\begin{minipage}[t]{4cm}%s\end{minipage}\hspace{5mm}%%
\begin{minipage}[t]{4cm}%s\end{minipage}\hspace{5mm}%%
\begin{minipage}[t]{4cm}%s\end{minipage}
\vfill
\begin{minipage}[t]{4cm}%s\end{minipage}\hspace{5mm}%%
\begin{minipage}[t]{4cm}%s\end{minipage}\hspace{5mm}%%
\begin{minipage}[t]{4cm}%s\end{minipage}
\vfill
\begin{minipage}[t]{4cm}%s\end{minipage}\hspace{5mm}%%
\begin{minipage}[t]{4cm}%s\end{minipage}\hspace{5mm}%%
\begin{minipage}[t]{4cm}%s\end{minipage}
\vfill
\begin{minipage}[t]{4cm}%s\end{minipage}\hspace{5mm}%%
\begin{minipage}[t]{4cm}%s\end{minipage}\hspace{5mm}%%
\begin{minipage}[t]{4cm}%s\end{minipage}`
