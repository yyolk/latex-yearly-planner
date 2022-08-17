package types

import "bytes"

type NamedBuffer struct {
	Name string

	Buffer *bytes.Buffer
}

type NamedBuffers []NamedBuffer
