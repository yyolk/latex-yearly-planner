package cell

import "github.com/kudrykv/latex-yearly-planner/app/tex"

type Cell struct {
	text string
	ref  bool
}

type Cells []Cell

func (r Cells) Slice() []string {
	if len(r) == 0 {
		return nil
	}

	out := make([]string, 0, len(r))

	for _, cell := range r {
		out = append(out, cell.Build())
	}

	return out
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
