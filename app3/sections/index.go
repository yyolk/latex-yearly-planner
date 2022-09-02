package sections

import (
	"errors"
	"fmt"
	"strings"

	"github.com/kudrykv/latex-yearly-planner/app3/types"
)

type IndexParameters struct {
	ItemsPerPage int
	Pages        int
	LineHeight   types.Millimeters
}

var (
	ErrNonPositiveItemsPerPage = errors.New("negative index per page")
	ErrNonPositivePages        = errors.New("negative pages")
	ErrNonPositiveLineHeight   = errors.New("negative line height")
)

func (r IndexParameters) Test() error {
	if r.ItemsPerPage <= 0 {
		return ErrNonPositiveItemsPerPage
	}

	if r.Pages <= 0 {
		return ErrNonPositivePages
	}

	if r.LineHeight <= 0 {
		return ErrNonPositiveLineHeight
	}

	return nil
}

type Index struct {
	parameters IndexParameters
}

func NewIndex(parameters IndexParameters) (Index, error) {
	if err := parameters.Test(); err != nil {
		return Index{}, fmt.Errorf("test index parameters: %w", err)
	}

	return Index{
		parameters: parameters,
	}, nil
}

func (r Index) Build() ([]string, error) {
	pages := make([]string, 0, r.parameters.Pages)

	for i := 1; i <= r.parameters.Pages; i++ {
		items := make([]string, 0, r.parameters.ItemsPerPage)

		for j := 1; j <= r.parameters.ItemsPerPage; j++ {
			itemNumber := (i-1)*r.parameters.ItemsPerPage + j

			items = append(items, fmt.Sprintf(`%d & \parbox{0pt}{\vskip%s} \\ \hline`, itemNumber, r.parameters.LineHeight))
		}

		pages = append(pages, fmt.Sprintf(indexTemplate, strings.Join(items, "\n")))
	}

	return pages, nil
}

const indexTemplate = `{\arrayrulecolor{gray}\begin{tabularx}{\linewidth}{l|l}
%s
\end{tabularx}}`

func (r Index) Repeat() int {
	return r.parameters.Pages
}
