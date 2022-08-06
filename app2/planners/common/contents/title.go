package contents

import "fmt"

type Title struct {
	name string
}

func NewTitle(name string) Title {
	return Title{name: name}
}

func (r Title) Build() ([]string, error) {
	return []string{fmt.Sprintf(titleTemplate, r.name)}, nil
}

const titleTemplate = `\hspace{0pt}\vfil
\hfill\resizebox{.7\linewidth}{!}{%s}%%
\pagebreak`
