package pages

import (
	"fmt"
	"strings"
)

type Block interface {
	Build() (string, error)
}

type Page struct {
	blocks []Block
}

func NewPage(blocks ...Block) Page {
	return Page{
		blocks: blocks,
	}
}

func (r Page) Build() (string, error) {
	if len(r.blocks) == 0 {
		return "", nil
	}

	out := make([]string, 0, len(r.blocks))

	for _, block := range r.blocks {
		text, err := block.Build()
		if err != nil {
			return "", fmt.Errorf("build %T: %w", block, err)
		}

		out = append(out, text)
	}

	return strings.Join(out, ""), nil
}
