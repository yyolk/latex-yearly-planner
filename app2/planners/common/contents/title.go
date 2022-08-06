package contents

type Title struct {
	name string
}

func NewTitle(name string) Title {
	return Title{name: name}
}

func (r Title) Build() ([]string, error) {
	return []string{`\hspace{0pt}\vfil
\hfill\resizebox{.7\linewidth}{!}{` + r.name + `}%
\pagebreak`}, nil
}
