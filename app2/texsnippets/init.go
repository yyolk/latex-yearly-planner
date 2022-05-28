package texsnippets

import "text/template"

var rootTpl *template.Template

func init() {
	rootTpl = template.New("")

	compileSnippets()
}

func compileSnippets() {
	for _, item := range templatesToCompile {
		rootTpl = template.Must(rootTpl.New(item.name).Parse(item.body))
	}
}

type item struct {
	name string
	body string
}

const (
	Document = "document"
	Title    = "title"
)

var templatesToCompile = []item{
	{Document, document},
	{Title, title},
}

const title = `\hspace{0pt}\vfil
\hfill\resizebox{.7\linewidth}{!}{ {{- .Year -}} }%
\pagebreak`
