package cell

type Cells []Cell

func NewCells(text ...string) Cells {
	if len(text) == 0 {
		return nil
	}

	cells := make(Cells, 0, len(text))

	for _, item := range text {
		cells = append(cells, New(item))
	}

	return cells
}

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

func (r Cells) Push(cell Cell) Cells {
	cells := make(Cells, len(r))
	copy(cells, r)

	return append(cells, cell)
}

func (r Cells) Shift(cell Cell) Cells {
	return append([]Cell{cell}, r...)
}

func (r Cells) Ref(text string) Cells {
	cells := make(Cells, len(r))
	copy(cells, r)

	for i, cell := range cells {
		if cell.text == text {
			cells[i] = cell.Ref()
		}
	}

	return cells
}
