package pages

import (
	"bytes"
	"fmt"
)

type Buffer struct {
	*bytes.Buffer
}

func NewBuffer() *Buffer {
	return &Buffer{&bytes.Buffer{}}
}

func (r *Buffer) WriteBlocks(blocks ...Block) error {
	compiledPages, err := NewPage(blocks...).Build()
	if err != nil {
		return fmt.Errorf("build pages from blocks: %w", err)
	}

	for _, page := range compiledPages {
		_, _ = r.WriteString(page + "\n\n" + `\pagebreak{}` + "\n")
	}

	return nil
}
