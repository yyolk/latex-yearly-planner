package ref

import (
	"github.com/kudrykv/latex-yearly-planner/app/tex"
)

type Link struct {
	text string
	ref  string
}

func NewLinkWithRef(text, ref string) Link {
	return Link{text: text, ref: ref}
}

func (l Link) Build() string {
	text, ref := l.text, l.ref
	if len(ref) == 0 {
		ref = l.text
	}

	return tex.Hyperlink(ref, text)
}
