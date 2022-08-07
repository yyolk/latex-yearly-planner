package pages2

import (
	"errors"
	"fmt"
	"strings"
)

type Block interface {
	Build() ([][]byte, error)
}

type Blocks []Block

var ErrBadSliceLen = errors.New("bad slice len")

type Page struct {
	Blocks Blocks
}

func New(blocks ...Block) Page {
	return Page{Blocks: blocks}
}

func (r Page) Build() ([][]byte, error) {
	if len(r.Blocks) == 0 {
		return nil, nil
	}

	slicesOfBuiltBlocks := make([][][]byte, 0, len(r.Blocks))

	for _, block := range r.Blocks {
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

	out := make([][]byte, 0, len(slicesOfBuiltBlocks))

	for pageIndex := 0; pageIndex < len(slicesOfBuiltBlocks[0]); pageIndex++ {
		var bts []byte

		for blockIndex := 0; blockIndex < len(slicesOfBuiltBlocks); blockIndex++ {
			bts = append(bts, slicesOfBuiltBlocks[blockIndex][pageIndex]...)
		}

		out = append(out, bts)
	}

	return out, nil
}

func (r Page) blockTypes() string {
	if len(r.Blocks) == 0 {
		return "<no blocks>"
	}

	out := make([]string, 0, len(r.Blocks))

	for _, block := range r.Blocks {
		out = append(out, fmt.Sprintf("%T", block))
	}

	return strings.Join(out, ", ")
}
