package planners

import (
	"bytes"
	"strings"
)

type futureFiles []futureFile

func (r futureFiles) buildAsTexIncludes() string {
	if len(r) == 0 {
		return ""
	}

	out := make([]string, 0, len(r))

	for _, file := range r {
		out = append(out, `\include{`+file.name+`}`)
	}

	return strings.Join(out, "\n")
}

type futureFile struct {
	name   string
	buffer *bytes.Buffer
}
