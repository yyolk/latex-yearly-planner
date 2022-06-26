package planners

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/kudrykv/latex-yearly-planner/app2/planners/common"
	"github.com/kudrykv/latex-yearly-planner/app2/planners/mos"
	"github.com/kudrykv/latex-yearly-planner/app2/texsnippets"
)

const (
	MonthsOnSidesTemplate = "mos"
)

var UnknownTemplateName = errors.New("unknown planner name")
var UnknownSectionErr = errors.New("unknown section")

type Planner struct {
	params      common.Params
	futureFiles []futureFile
	dir         string
	builder     mos.MonthsOnSides
}

func New(template string, params common.Params) (*Planner, error) {
	var builder mos.MonthsOnSides

	switch template {
	case MonthsOnSidesTemplate:
		builder = mos.New(params)
	default:
		return nil, fmt.Errorf("%s: %w", template, UnknownTemplateName)
	}

	return &Planner{
		params:      params,
		futureFiles: nil,
		dir:         "",
		builder:     builder,
	}, nil
}

func (r *Planner) Generate() error {
	layout, err := common.NewLayout(r.params.Device, r.params.Hand)
	if err != nil {
		return fmt.Errorf("new layout: %w", err)
	}

	if err = r.builder.PrepareDetails(r.params.Device, layout); err != nil {
		return fmt.Errorf("prepare details: %w", err)
	}

	sections := r.builder.Sections()

	for _, name := range r.params.Sections {
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
	for i := 0; i < r.builder.RunTimes(); i++ {
		cmd := exec.CommandContext(ctx, "pdflatex", "./document.tex")
		cmd.Dir = r.dir

		if err := cmd.Run(); err != nil {
			return fmt.Errorf("run pdflatex: %w", err)
		}
	}

	return nil
}

func (r *Planner) createRootDocument() error {
	futureFiles := make([]string, 0, len(r.futureFiles))

	for _, file := range r.futureFiles {
		futureFiles = append(futureFiles, `\include{`+file.name+`}`)
	}

	buffer := &bytes.Buffer{}
	if err := texsnippets.Execute(buffer, texsnippets.Document, map[string]interface{}{
		"Device":     r.params.Device,
		"Layout":     r.builder.Layout(),
		"Files":      strings.Join(futureFiles, "\n"),
		"ShowFrames": r.params.ShowFrames,
		"ShowLinks":  r.params.ShowLinks,
	}); err != nil {
		return fmt.Errorf("execute template root-document: %w", err)
	}

	r.futureFiles = append(r.futureFiles, futureFile{name: "document", buffer: buffer})

	return nil
}
