package cell

import "github.com/kudrykv/latex-yearly-planner/app/tex"

type Cell struct {
	text string
	ref  bool
}

func New(text string) Cell {
	return Cell{text: text}
}

func (c Cell) Ref() Cell {
	c.ref = true

	return c
}

func (c Cell) Build() string {
	if !c.ref {
		return tex.Hyperlink(c.text, c.text)
	}

	return tex.CellColor("black", tex.TextColor("white", tex.Hypertarget(c.text, c.text)))
}
