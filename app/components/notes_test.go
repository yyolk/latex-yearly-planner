package components_test

import (
	"testing"

	"github.com/kudrykv/latex-yearly-planner/app/components"
	. "github.com/kudrykv/latex-yearly-planner/app/test"
	"github.com/stretchr/testify/assert"
)

func TestNotes_Build(t *testing.T) {
	t.Parallel()

	dotted := components.Dotted{
		Distance: 5,
		Columns:  6,
		Rows:     7,
	}

	lined := components.Lined{
		Number: 5,
		Height: 6,
	}

	dottedConfig := components.NotesParameters{Dotted: dotted}
	linedConfig := components.NotesParameters{Lined: lined}

	t.Run("dotted notes", func(t *testing.T) {
		t.Parallel()

		notes, err := components.NewNotes(dottedConfig)

		assert.NoError(t, err)
		assert.Equal(t, Fixture("notes_dotted"), notes.Build())
	})

	t.Run("lined notes", func(t *testing.T) {
		t.Parallel()

		notes, err := components.NewNotes(linedConfig)

		assert.NoError(t, err)
		assert.Equal(t, Fixture("notes_lined"), notes.Build())
	})

	t.Run("error if both configs are present", func(t *testing.T) {
		t.Parallel()

		_, err := components.NewNotes(components.NotesParameters{Dotted: dotted, Lined: lined})

		assert.ErrorIs(t, err, components.ErrBothLinedAndDottedSet)
	})

	t.Run("error if both configs are missing", func(t *testing.T) {
		t.Parallel()

		_, err := components.NewNotes(components.NotesParameters{})

		assert.ErrorIs(t, err, components.ErrBothLinedAndDottedMissing)
	})
}
