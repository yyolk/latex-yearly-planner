package ref

import "github.com/kudrykv/latex-yearly-planner/app/tex"

type Target struct {
	text string
	ref  string
}

func NewTarget(text string) Target {
	return Target{text: text}
}

func NewTargetWithRef(text, ref string) Target {
	return Target{text: text, ref: ref}
}

func (t Target) Build() string {
	text, ref := t.text, t.ref
	if len(ref) == 0 {
		ref = text
	}

	return tex.Hypertarget(ref, text)
}
