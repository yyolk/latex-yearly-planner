package texsnippets

import "io"

func Execute(writer io.Writer, name string, data any) error {
	return rootTpl.ExecuteTemplate(writer, name, data)
}
