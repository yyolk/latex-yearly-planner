package mos

import (
	"strconv"
	"strings"
)

type index struct {
	refPrefix string
}

func (i index) Build() ([]string, error) {
	return []string{
		`\vskip1mm\begin{minipage}[t]{\myLengthTwoColumnWidth}
` + strings.Join(i.todoCol(1, 29), "\n") + `
\end{minipage}%
\hspace{\myLengthTwoColumnsSeparatorWidth}%
\begin{minipage}[t]{\myLengthTwoColumnWidth}
` + strings.Join(i.todoCol(30, 58), "\n") + `
\end{minipage}`,
		`\vskip1mm\begin{minipage}[t]{\myLengthTwoColumnWidth}
` + strings.Join(i.todoCol(59, 87), "\n") + `
\end{minipage}%
\hspace{\myLengthTwoColumnsSeparatorWidth}%
\begin{minipage}[t]{\myLengthTwoColumnWidth}
` + strings.Join(i.todoCol(87, 115), "\n") + `
\end{minipage}`,
	}, nil
}

func (r index) todoCol(from, to int) []string {
	var col []string

	for i := from; i <= to; i++ {
		itoa := strconv.Itoa(i)
		linked := `\hyperlink{` + r.refPrefix + `-` + itoa + `}{\parbox{1cm}{` + itoa + `.}}`
		col = append(col, `\parbox{0pt}{\vskip7mm}`+linked+`\myLineGray`)
	}
	return col
}
