package sections

import (
	"fmt"

	"github.com/kudrykv/latex-yearly-planner/app3/components"
	"github.com/kudrykv/latex-yearly-planner/lib/calendar"
)

type QuarterlyParameters struct {
	CalendarsToTheRight      bool
	LittleCalendarParameters components.LittleCalendarParameters
	NotesParameters          components.NotesParameters
}

type Quarterly struct {
	parameters QuarterlyParameters
	quarter    calendar.Quarter
	notes      components.Notes
}

func NewQuarterly(quarter calendar.Quarter, parameters QuarterlyParameters) (Quarterly, error) {
	notes, err := components.NewNotes(parameters.NotesParameters)
	if err != nil {
		return Quarterly{}, fmt.Errorf("new notes: %w", err)
	}

	return Quarterly{
		quarter:    quarter,
		parameters: parameters,
		notes:      notes,
	}, nil
}

func (r Quarterly) Title() string {
	return r.quarter.Name()
}

func (r Quarterly) Reference() string {
	return r.quarter.Name()
}

func (r Quarterly) Build() ([]string, error) {
	leftColumn := r.calendarColumn()
	rightColumn := r.notesColumn()

	if r.parameters.CalendarsToTheRight {
		leftColumn, rightColumn = rightColumn, leftColumn
	}

	return []string{leftColumn + `\hspace{5mm}` + rightColumn}, nil
}

func (r Quarterly) calendarColumn() string {
	mon1 := components.NewLittleCalendarFromMonth(r.quarter.Months[0], r.parameters.LittleCalendarParameters)
	mon2 := components.NewLittleCalendarFromMonth(r.quarter.Months[1], r.parameters.LittleCalendarParameters)
	mon3 := components.NewLittleCalendarFromMonth(r.quarter.Months[2], r.parameters.LittleCalendarParameters)

	return fmt.Sprintf(
		quarterlyCalendarTemplate,
		mon1.Build(),
		mon2.Build(),
		mon3.Build(),
	)
}

const quarterlyCalendarTemplate = `\begin{minipage}[t][18cm]{5cm}
%s
\vfill
%s
\vfill
%s
\end{minipage}`

func (r Quarterly) notesColumn() string {
	return fmt.Sprintf(quarterlyNotesTemplate, r.notes.Build())
}

const quarterlyNotesTemplate = `\begin{minipage}[t][18cm]{8cm}
%s
\end{minipage}`
