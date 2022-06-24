package mos

import "strings"

type todoContents struct {
}

func (t todoContents) Build() ([]string, error) {
	var block []string

	for i := 0; i < 8; i++ {
		block = append(block, `\parbox{0pt}{\vskip7mm}$\square$\myLineGray`)
	}

	blockBlob := `\myUnderline{\textcolor{white}{Q}}` + "\n" + strings.Join(block, "\n")

	var colled []string

	for i := 0; i < 3; i++ {
		colled = append(colled, blockBlob)
	}

	col := strings.Join(colled, `\vfill`+"\n")

	return []string{
		`\vskip5mm\begin{minipage}[t][\remainingHeight]{\myLengthTwoColumnWidth}
` + col + `
\end{minipage}%
\hspace{\myLengthTwoColumnsSeparatorWidth}%
\begin{minipage}[t][\remainingHeight]{\myLengthTwoColumnWidth}
` + col + `
\end{minipage}`,
	}, nil
}
