package planners

import (
	"fmt"
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/kudrykv/latex-yearly-planner/app2/devices"
)

type MonthsOnSides struct {
	params      Params
	futureFiles []futureFile
	templates   *template.Template
}

type futureFile struct {
	name    string
	builder *strings.Builder
}

func newMonthsOnSides(params Params) (*MonthsOnSides, error) {
	mos := &MonthsOnSides{params: params}

	if err := mos.init(); err != nil {
		return nil, fmt.Errorf("init: %w", err)
	}

	return mos, nil
}

func (r *MonthsOnSides) GenerateFor(device devices.Device) error {
	if err := r.createTitle(); err != nil {
		return fmt.Errorf("create title: %w", err)
	}

	return nil
}

func (r *MonthsOnSides) WriteTo(dir string) error {
	for _, future := range r.futureFiles {
		if err := os.WriteFile(path.Join(dir, future.name), []byte(future.builder.String()), 0600); err != nil {
			return fmt.Errorf("write file %s: %w", future.name, err)
		}
	}

	return nil
}

func (r *MonthsOnSides) init() error {
	var err error

	if r.templates, err = bigTemplateStuff(); err != nil {
		return fmt.Errorf("big template stuff: %w", err)
	}

	return nil
}

func (r *MonthsOnSides) createTitle() error {
	builder := &strings.Builder{}

	if err := r.templates.ExecuteTemplate(builder, "title", r.params.TemplateData); err != nil {
		return fmt.Errorf("execute template title: %w", err)
	}

	r.futureFiles = append(r.futureFiles, futureFile{name: "title.tex", builder: builder})

	return nil
}

func (r *MonthsOnSides) writeTo(dir string) {

}

func bigTemplateStuff() (*template.Template, error) {
	tpls := template.New("")

	var err error

	for _, row := range [][]string{{"title", titleTex}} {
		tpls = tpls.New(row[0])

		if tpls, err = tpls.Parse(row[1]); err != nil {
			return nil, fmt.Errorf("parse %s: %w", row[1], err)
		}
	}

	return tpls, nil
}

const titleTex = `\hspace{0pt}\vfil
\hfill\resizebox{.7\linewidth}{!}{ {{- .Year -}} }%
\pagebreak`
