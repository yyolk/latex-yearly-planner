package components

import (
	"fmt"
	"strings"
)

type Tabs []Tab

func (r Tabs) String() string {
	if len(r) == 0 {
		return ""
	}

	pieces := make([]string, 0, len(r))

	for _, tab := range r {
		pieces = append(pieces, tab.String())
	}

	return strings.Join(pieces, " & ")
}

type Tab struct {
	Text      string
	Reference string
	Target    bool
}

func (r Tab) String() string {
	text := r.Text
	ref := r.Reference

	if len(ref) == 0 {
		ref = text
	}

	if r.Target {
		return fmt.Sprintf(`\cellcolor{black}{\textcolor{white}{%s}}`, text)
	}

	return fmt.Sprintf(`\hyperlink{%s}{%s}`, ref, text)
}
