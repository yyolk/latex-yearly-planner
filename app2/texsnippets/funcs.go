package texsnippets

import (
	"io"
	"strings"
)

func Execute(writer io.Writer, name string, data any) error {
	return rootTpl.ExecuteTemplate(writer, name, data)
}

func Build(name string, data any) (string, error) {
	builder := &strings.Builder{}

	if err := Execute(builder, name, data); err != nil {
		return "", err
	}

	return builder.String(), nil
}
