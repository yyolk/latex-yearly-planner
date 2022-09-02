package sections

import (
	"fmt"

	"github.com/kudrykv/latex-yearly-planner/app3/components"
)

type NotesParameters struct {
	IndexParameters IndexParameters
	NotesParameters components.NotesParameters
}

type Notes struct {
	parameters NotesParameters
	notes      components.Notes
}

func NewNotes(parameters NotesParameters) (Notes, error) {
	notes, err := components.NewNotes(parameters.NotesParameters)
	if err != nil {
		return Notes{}, fmt.Errorf("new notes: %w", err)
	}

	return Notes{
		parameters: parameters,
		notes:      notes,
	}, nil
}

func (r Notes) Build() ([]string, error) {
	return []string{r.notes.Build()}, nil
}
