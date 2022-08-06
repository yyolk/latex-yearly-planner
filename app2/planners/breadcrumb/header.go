package breadcrumb

import (
	"strings"

	"github.com/kudrykv/latex-yearly-planner/app/tex"
	"github.com/kudrykv/latex-yearly-planner/app2/tex/cell"
)

type header struct {
	left  cell.Cells
	right cell.Cells
}

func (r header) Build() ([]string, error) {
	leftTabular := tex.Tabular(strings.Repeat("l", len(r.left)), strings.Join(r.left.Slice(), " & "))
	if len(r.left) == 0 {
		leftTabular = ""
	}

	rightTabular := tex.Tabular(strings.Repeat("r", len(r.right)), strings.Join(r.right.Slice(), " & "))
	if len(r.right) == 0 {
		rightTabular = ""
	}

	return []string{
		`{\LARGE{}` + leftTabular + `\hfill{}` + rightTabular + `}\myLinePlain` + "\n\n",
	}, nil
}
