package sections

import (
	"fmt"

	"github.com/kudrykv/latex-yearly-planner/app/calendar"
	"github.com/kudrykv/latex-yearly-planner/app/components"
)

type DailyNotesParameters struct {
	Pages           int
	NotesParameters components.NotesParameters
}

type DailyNotes struct {
	parameters DailyNotesParameters
	day        calendar.Day
	notes      components.Notes
}

func NewDailyNotes(day calendar.Day, parameters DailyNotesParameters) (DailyNotes, error) {
	notes, err := components.NewNotes(parameters.NotesParameters)
	if err != nil {
		return DailyNotes{}, fmt.Errorf("new notes: %w", err)
	}

	return DailyNotes{
		day:        day,
		parameters: parameters,
		notes:      notes,
	}, nil
}

func (r DailyNotes) Build() ([]string, error) {
	pages := make([]string, 0, r.parameters.Pages)

	for i := 0; i < r.parameters.Pages; i++ {
		pages = append(pages, r.notes.Build())
	}

	return pages, nil
}

func (r DailyNotes) Link() string {
	return fmt.Sprintf(`\hyperlink{daily-notes-%s}{%s}`, r.day.Format("2006-01-02"), "More")
}

func (r DailyNotes) Reference() string {
	return fmt.Sprintf(`daily-notes-%s`, r.day.Format("2006-01-02"))
}

func (r DailyNotes) Repeat() int {
	return r.parameters.Pages
}
