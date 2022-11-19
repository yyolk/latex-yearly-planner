package components

import (
	"errors"
	"fmt"
	"strings"

	"github.com/kudrykv/latex-yearly-planner/app/types"
)

type NotesParameters struct {
	Lined  Lined
	Dotted Dotted
}

type Lined struct {
	Number int
	Height types.Millimeters
}

type Dotted struct {
	Distance types.Millimeters
	Columns  int
	Rows     int
}

var (
	ErrBothLinedAndDottedSet     = errors.New("both lined and dotted parameters are set")
	ErrBothLinedAndDottedMissing = errors.New("both lined and dotted parameters are unset")
	ErrNonPositiveLineHeight     = errors.New("line height is not positive")
	ErrNonPositiveRows           = errors.New("rows is not positive")
	ErrNonPositiveColumns        = errors.New("columns is not positive")
)

func (r NotesParameters) Test() error {
	if r.Lined.Number > 0 && r.Dotted.Distance > 0 {
		return ErrBothLinedAndDottedSet
	}

	if r.Lined.Number <= 0 && r.Dotted.Distance <= 0 {
		return ErrBothLinedAndDottedMissing
	}

	if r.Lined.Number > 0 {
		if r.Lined.Height <= 0 {
			return ErrNonPositiveLineHeight
		}
	}

	if r.Dotted.Distance > 0 {
		if r.Dotted.Rows <= 0 {
			return ErrNonPositiveRows
		}

		if r.Dotted.Columns <= 0 {
			return ErrNonPositiveColumns
		}
	}

	return nil
}

type Notes struct {
	parameters NotesParameters
}

func (r Notes) Build() string {
	if r.parameters.Lined.Number > 0 {
		return r.buildLines()
	}

	return r.buildDots()
}

func (r Notes) buildLines() string {
	pieces := make([]string, 0, r.parameters.Lined.Number)

	for i := 0; i < r.parameters.Lined.Number; i++ {
		pieces = append(pieces, fmt.Sprintf(`\parbox{0pt}{\vskip%s}\myLineLightGray`, r.parameters.Lined.Height))
	}

	return strings.Join(pieces, "\n")
}

func (r Notes) buildDots() string {
	return fmt.Sprintf(
		`\vskip%s\leavevmode\multido{\dC=0mm+%s}{%d}{\multido{\dR=0mm+%s}{%d}{\put(\dR,\dC){\scriptsize.}}}`,
		r.parameters.Dotted.Distance,
		r.parameters.Dotted.Distance,
		r.parameters.Dotted.Rows,
		r.parameters.Dotted.Distance,
		r.parameters.Dotted.Columns,
	)
}

func NewNotes(parameters NotesParameters) (Notes, error) {
	if err := parameters.Test(); err != nil {
		return Notes{}, err
	}

	return Notes{parameters}, nil
}
