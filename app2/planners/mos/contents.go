package mos

type notesContents struct{}

func (r notesContents) Build() ([]string, error) {
	return []string{`\vskip5mm\hspace{0.5mm}\vbox to 0mm{\myDotGrid{41}{29}}`}, nil
}
