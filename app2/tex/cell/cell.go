package cell

import "github.com/kudrykv/latex-yearly-planner/app/tex"

type Cell struct {
	text  string
	refAs string
	ref   bool
}

func New(text string) Cell {
	return Cell{text: text}
}

func (c Cell) Ref() Cell {
	c.ref = true

	return c
}

func (c Cell) Build() string {
	refAs := c.refAs
	if len(refAs) == 0 {
		refAs = c.text
	}

	if !c.ref {
		return tex.Hyperlink(refAs, c.text)
	}

	return tex.CellColor("black", tex.TextColor("white", tex.Hypertarget(refAs, c.text)))
}

func (c Cell) RefAs(refAs string) Cell {
	c.refAs = refAs

	return c
}
