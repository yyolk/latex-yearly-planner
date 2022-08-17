package types

import "bytes"

type NamedBuffer struct {
	bytes.Buffer

	Name string
}

type NamedBuffers []NamedBuffer
