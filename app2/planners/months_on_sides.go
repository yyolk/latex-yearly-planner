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

type MonthsOnSides struct {
	params      Params
	futureFiles []futureFile
	dir         string
	builder     monthsOnSidesBuilder
}

var UnknownSectionErr = errors.New("unknown section")

func (r *MonthsOnSides) GenerateFor(device devices.Device) error {
	layout, err := newLayout(device)
	if err != nil {
		return fmt.Errorf("new layout: %w", err)
	}

	r.params.TemplateData.Apply(WithLayout(layout), WithDevice(device))

	sections := r.builder.sections()

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

func (r *MonthsOnSides) WriteTeXTo(dir string) error {
	for _, future := range r.futureFiles {
		if err := os.WriteFile(path.Join(dir, future.name+".tex"), future.buffer.Bytes(), 0600); err != nil {
			return fmt.Errorf("write file %s: %w", future.name, err)
		}
	}

	r.dir = dir

	return nil
}

func (r *MonthsOnSides) Compile(ctx context.Context) error {
	for i := 0; i < 2; i++ {
		cmd := exec.CommandContext(ctx, "pdflatex", "./document.tex")
		cmd.Dir = r.dir

		if err := cmd.Run(); err != nil {
			return fmt.Errorf("run pdflatex: %w", err)
		}
	}

	return nil
}

func (r *MonthsOnSides) createRootDocument() error {
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
