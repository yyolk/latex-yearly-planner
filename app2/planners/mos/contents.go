package mos

type notesContents struct{}

func (r notesContents) Build() ([]string, error) {
	return []string{`\vskip5mm\hspace{0.5mm}\vbox to 0mm{\myDotGrid{41}{29}}`}, nil
}

type titleContents struct {
	title string
}

func (r titleContents) Build() ([]string, error) {
	return []string{`\hspace{0pt}\vfil
\hfill\resizebox{.7\linewidth}{!}{` + r.title + `}%
\pagebreak`}, nil
}
