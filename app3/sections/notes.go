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
	index      Index
	page       int
}

func NewNotes(index Index, parameters NotesParameters) (Notes, error) {
	notes, err := components.NewNotes(parameters.NotesParameters)
	if err != nil {
		return Notes{}, fmt.Errorf("new notes: %w", err)
	}

	return Notes{
		parameters: parameters,
		index:      index,
		notes:      notes,
	}, nil
}

func (r Notes) Build() ([]string, error) {
	return []string{r.notes.Build()}, nil
}

func (r Notes) CurrentPage(page int) Notes {
	r.page = page

	return r
}

func (r Notes) Title() string {
	return fmt.Sprintf("Note %d", r.page)
}
