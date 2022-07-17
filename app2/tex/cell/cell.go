package cell

import "github.com/kudrykv/latex-yearly-planner/app/tex"

type Cell struct {
	text     string
	refAs    string
	ref      bool
	noTarget bool
}

func New(text string) Cell {
	return Cell{text: text}
}

func (r Cell) Ref() Cell {
	r.ref = true

	return r
}

func (r Cell) Build() string {
	refAs := r.refAs
	if len(refAs) == 0 {
		refAs = r.text
	}

	if !r.ref {
		return tex.Hyperlink(refAs, r.text)
	}

	if r.noTarget {
		return tex.CellColor("black", tex.TextColor("white", r.text))
	}

	return tex.CellColor("black", tex.TextColor("white", tex.Hypertarget(refAs, r.text)))
}

func (r Cell) RefAs(refAs string) Cell {
	r.refAs = refAs

	return r
}

func (r Cell) NoTarget() Cell {
	r.noTarget = true
	r.ref = true

	return r
}
