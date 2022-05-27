package planners

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"
	"text/template"

	"github.com/kudrykv/latex-yearly-planner/app2/devices"
)

type MonthsOnSides struct {
	params      Params
	futureFiles []futureFile
	templates   *template.Template
	dir         string
}

var UnknownSectionErr = errors.New("unknown section")

func newMonthsOnSides(params Params) (*MonthsOnSides, error) {
	mos := &MonthsOnSides{
		params: params,
	}

	if err := mos.init(); err != nil {
		return nil, fmt.Errorf("init: %w", err)
	}

	return mos, nil
}

func (r *MonthsOnSides) GenerateFor(device devices.Device) error {
	layout, err := newLayout(device)
	if err != nil {
		return fmt.Errorf("new layout: %w", err)
	}

	r.params.TemplateData.Apply(WithLayout(layout), WithDevice(device))

	sections := r.sections()

	for _, section := range r.params.TemplateData.sections {
		sectionFunc, ok := sections[section]
		if !ok {
			return fmt.Errorf("%v: %w", section, UnknownSectionErr)
		}

		if err = sectionFunc(); err != nil {
			return fmt.Errorf("%v: %w", section, err)
		}
	}

	if err := r.createRootDocument(); err != nil {
		return fmt.Errorf("create root document: %w", err)
	}

	return nil
}

func (r *MonthsOnSides) sections() map[string]func() error {
	return map[string]func() error{
		TitleSection:  r.createTitle,
		AnnualSection: r.annualSection,
	}
}

func (r *MonthsOnSides) WriteTeXTo(dir string) error {
	for _, future := range r.futureFiles {
		if err := os.WriteFile(path.Join(dir, future.name), []byte(future.buffer.String()), 0600); err != nil {
			return fmt.Errorf("write file %s: %w", future.name, err)
		}
	}

	r.dir = dir

	return nil
}

func (r *MonthsOnSides) Compile(ctx context.Context) error {
	cmd := exec.CommandContext(ctx, "pdflatex", "./document.tex")
	cmd.Dir = r.dir

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("run pdflatex: %w", err)
	}

	return nil
}

func (r *MonthsOnSides) init() error {
	var err error

	if r.templates, err = createTemplates(); err != nil {
		return fmt.Errorf("big template stuff: %w", err)
	}

	return nil
}

func (r *MonthsOnSides) createTitle() error {
	buffer := &bytes.Buffer{}

	if err := r.templates.ExecuteTemplate(buffer, titleTpl, r.params.TemplateData); err != nil {
		return fmt.Errorf("execute template title: %w", err)
	}

	r.futureFiles = append(r.futureFiles, futureFile{name: "title.tex", buffer: buffer})

	return nil
}

func (r *MonthsOnSides) createRootDocument() error {
	files := make([]string, 0, len(r.futureFiles))

	for _, file := range r.futureFiles {
		files = append(files, file.name)
	}

	r.params.TemplateData.Apply(WithFiles(files))

	buffer := &bytes.Buffer{}
	if err := r.templates.ExecuteTemplate(buffer, rootDocumentTpl, r.params.TemplateData); err != nil {
		return fmt.Errorf("execute template root-document: %w", err)
	}

	r.futureFiles = append(r.futureFiles, futureFile{name: "document.tex", buffer: buffer})

	return nil
}

func (r *MonthsOnSides) annualSection() error {
	buffer := &bytes.Buffer{}

	buffer.Write([]byte("hello world"))

	r.futureFiles = append(r.futureFiles, futureFile{name: "annual.tex", buffer: buffer})

	return nil
}
