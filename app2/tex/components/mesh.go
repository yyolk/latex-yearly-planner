package components

import "fmt"

type Mesh struct {
	Rows int
	Cols int
}

func NewMesh(rows, cols int) Mesh {
	return Mesh{Rows: rows, Cols: cols}
}

func (r Mesh) Build() string {
	return fmt.Sprintf(meshFormat, r.Rows, r.Cols)
}

const meshFormat = `\vspace{5mm}\hspace{.5mm}\vbox to 0mm{\myDotGrid{%d}{%d}}`
