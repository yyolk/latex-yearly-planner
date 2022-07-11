package mos

import (
	"fmt"

	"github.com/kudrykv/latex-yearly-planner/app2/tex/components"
)

type index struct {
	typ       int
	pages     int
	perColumn int
}

func (r index) Build() ([]string, error) {
	return r.buildIndexPages(), nil
}

func (r index) buildIndexPages() []string {
	out := make([]string, 0, r.pages)

	for i := 1; i <= r.pages; i++ {
		out = append(out, r.buildIndexPage(i))
	}

	return out
}

func (r index) buildIndexPage(pageNum int) string {
	from := r.perColumn*2*(pageNum-1) + 1
	to := from + r.perColumn - 1

	return fmt.Sprintf(
		indexFormat,
		components.NewHyperlines(r.typ, from, to).Build(),
		components.NewHyperlines(r.typ, to+1, to+r.perColumn).Build(),
	)
}

const indexFormat = `\begin{minipage}[t]{\myLengthTwoColumnWidth}
%s
\end{minipage}%%
\hspace{\myLengthTwoColumnsSeparatorWidth}%%
\begin{minipage}[t]{\myLengthTwoColumnWidth}
%s
\end{minipage}`
