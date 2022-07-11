package components

import (
	"strconv"
	"strings"

	"github.com/kudrykv/latex-yearly-planner/app2/tex/ref"
)

type Hyperlines []Hyperline

func NewHyperlines(typ, from, to int) Hyperlines {
	lines := make(Hyperlines, 0, to-from+1)

	for i := from; i <= to; i++ {
		lines = append(lines, NewHyperline(typ, i))
	}

	return lines
}

func (r Hyperlines) Build() string {
	out := make([]string, 0, len(r))

	for _, hyperline := range r {
		out = append(out, hyperline.Build())
	}

	return strings.Join(out, "\n")
}

type Hyperline struct {
	item ref.Item
}

func NewHyperline(typ, pos int) Hyperline {
	itoa := strconv.Itoa(pos)

	return Hyperline{
		item: ref.NewItem(typ, `\parbox{1cm}{`+itoa+`.}`, itoa),
	}
}

func (r Hyperline) Build() string {
	return `\parbox{0pt}{\vskip7mm}` + r.item.Build() + `\myLineGray`
}
