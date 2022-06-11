package planners

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.com/kudrykv/latex-yearly-planner/app2/devices"
	"github.com/kudrykv/latex-yearly-planner/app2/texsnippets"
)

const (
	MonthsOnSidesTemplate = "mos"
)

var UnknownTemplateName = errors.New("unknown planner name")
var UnknownSectionErr = errors.New("unknown section")

type Planner struct {
	params      Params
	futureFiles []futureFile
	dir         string
	builder     MonthsOnSides
}

func New(params Params) (*Planner, error) {
	var builder MonthsOnSides

	switch params.Name {
	case MonthsOnSidesTemplate:
		builder = newMonthOnSides(params.TemplateData)
	default:
		return nil, fmt.Errorf("%s: %w", params.Name, UnknownTemplateName)
	}

	return &Planner{
		params:      params,
		futureFiles: nil,
		dir:         "",
		builder:     builder,
	}, nil
}

func (r *Planner) GenerateFor(device devices.Device, hand MainHand) error {
	layout, err := newLayout(device, hand)
	if err != nil {
		return fmt.Errorf("new layout: %w", err)
	}

	r.params.TemplateData.Apply(WithLayout(layout), WithDevice(device))

	sections := r.builder.Sections()

	for _, name := range r.params.TemplateData.sections {
		sectionFunc, ok := sections[name]
		if !ok {
			return fmt.Errorf("%v: %w", name, UnknownSectionErr)
		}

		buffer, err := sectionFunc()
		if err != nil {
			return fmt.Errorf("%v: %w", name, err)
		}

		r.futureFiles = append(r.futureFiles, futureFile{
			name:   name,
			buffer: buffer,
		})
	}

	if err := r.createRootDocument(); err != nil {
		return fmt.Errorf("create root document: %w", err)
	}

	return nil
}

func (r *Planner) WriteTeXTo(dir string) error {
	for _, future := range r.futureFiles {
		if err := os.WriteFile(path.Join(dir, future.name+".tex"), future.buffer.Bytes(), 0600); err != nil {
			return fmt.Errorf("write file %s: %w", future.name, err)
		}
	}

	r.dir = dir

	return nil
}

func (r *Planner) Compile(ctx context.Context) error {
	for i := 0; i < 2; i++ {
		cmd := exec.CommandContext(ctx, "pdflatex", "./document.tex")
		cmd.Dir = r.dir

		if err := cmd.Run(); err != nil {
			return fmt.Errorf("run pdflatex: %w", err)
		}
	}

	return nil
}

func (r *Planner) createRootDocument() error {
	files := make([]string, 0, len(r.futureFiles))

	for _, file := range r.futureFiles {
		files = append(files, file.name)
	}

	r.params.TemplateData.Apply(WithFiles(files))

	buffer := &bytes.Buffer{}
	if err := texsnippets.Execute(buffer, texsnippets.Document, r.params.TemplateData); err != nil {
		return fmt.Errorf("execute template root-document: %w", err)
	}

	r.futureFiles = append(r.futureFiles, futureFile{name: "document", buffer: buffer})

	return nil
}
