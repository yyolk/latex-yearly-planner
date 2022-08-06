package types

import "bytes"

type SectionFunc func() (*bytes.Buffer, error)
