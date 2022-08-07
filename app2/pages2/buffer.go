package pages2

import (
	"bytes"
	"fmt"
)

type Buffer struct {
	buffer *bytes.Buffer
}

func NewBuffer() *Buffer {
	return &Buffer{
		buffer: bytes.NewBuffer(nil),
	}
}

func (r *Buffer) WriteBlocks(blocks ...Block) error {
	built, err := New(blocks...).Build()
	if err != nil {
		return fmt.Errorf("build pages from blocks: %w", err)
	}

	for _, page := range built {
		_, _ = r.buffer.Write(page)
		_, _ = r.buffer.WriteString("\n\n")
		_, _ = r.buffer.WriteString(`\pagebreak{}`)
		_, _ = r.buffer.WriteString("\n")
	}

	return nil
}

func (r *Buffer) Bytes() []byte {
	return r.buffer.Bytes()
}

func (r *Buffer) Raw() *bytes.Buffer {
	return r.buffer
}
