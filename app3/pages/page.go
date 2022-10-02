package pages

import (
	"errors"
	"fmt"
	"strings"
)

type Block interface {
	Build() ([]string, error)
}

var ErrBadSliceLen = errors.New("bad slice len")

type Page struct {
	blocks []Block
}

func NewPage(blocks ...Block) Page {
	return Page{
		blocks: blocks,
	}
}

func (r Page) Build() ([]string, error) {
	if len(r.blocks) == 0 {
		return nil, nil
	}

	slicesOfBuiltBlocks := make([][]string, 0, len(r.blocks))

	for _, block := range r.blocks {
		piece, err := block.Build()
		if err != nil {
			return nil, fmt.Errorf("build %T: %w", block, err)
		}

		slicesOfBuiltBlocks = append(slicesOfBuiltBlocks, piece)
	}

	itemsOnPage := len(slicesOfBuiltBlocks[0])

	for i := 1; i < len(slicesOfBuiltBlocks); i++ {
		if len(slicesOfBuiltBlocks[i]) != itemsOnPage {
			return nil, fmt.Errorf("%s :%w", r.blockTypes(), ErrBadSliceLen)
		}
	}

	out := make([]string, 0, len(slicesOfBuiltBlocks))

	for pageIndex := 0; pageIndex < len(slicesOfBuiltBlocks[0]); pageIndex++ {
		str := ""

		for blockIndex := 0; blockIndex < len(slicesOfBuiltBlocks); blockIndex++ {
			str += slicesOfBuiltBlocks[blockIndex][pageIndex]
		}

		out = append(out, str)
	}

	return out, nil
}

func (r Page) blockTypes() string {
	if len(r.blocks) == 0 {
		return "<no blocks>"
	}

	out := make([]string, 0, len(r.blocks))

	for _, block := range r.blocks {
		out = append(out, fmt.Sprintf("%T", block))
	}

	return strings.Join(out, ", ")
}
