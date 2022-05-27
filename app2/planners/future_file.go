package planners

import "bytes"

type futureFile struct {
	name   string
	buffer *bytes.Buffer
}
