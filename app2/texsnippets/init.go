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

var templatesToCompile = []item{
	{Document, document},
}
