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
	page       int
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
	if r.page < 1 || r.page > r.parameters.Pages {
		return nil, fmt.Errorf("invalid page: %d", r.page)
	}

	items := make([]string, 0, r.parameters.ItemsPerPage)

	for j := 1; j <= r.parameters.ItemsPerPage; j++ {
		itemNumber := (r.page-1)*r.parameters.ItemsPerPage + j

		items = append(items, fmt.Sprintf(`%d & \parbox{0pt}{\vskip%s} \\ \hline`, itemNumber, r.parameters.LineHeight))
	}

	return []string{fmt.Sprintf(indexTemplate, strings.Join(items, "\n"))}, nil
}

const indexTemplate = `{\arrayrulecolor{gray}\renewcommand{\arraystretch}{0}\begin{tabularx}{\linewidth}{l|l}
%s
\end{tabularx}}`

func (r Index) IndexPages() int {
	return r.parameters.Pages
}

func (r Index) Title() string {
	postfix := ""
	if r.page > 1 {
		postfix = fmt.Sprintf(" %d", r.page)
	}

	return "Index" + postfix
}

func (r Index) CurrentPage(page int) Index {
	r.page = page

	return r
}

func (r Index) ItemPages() int {
	return r.parameters.ItemsPerPage * r.parameters.Pages
}

func (r Index) IndexPageFromItemPage(page int) Index {
	r.page = (page-1)/r.parameters.ItemsPerPage + 1

	return r
}

func (r Index) Reference() string {
	return r.Title()
}
