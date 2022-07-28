package planners

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.com/kudrykv/latex-yearly-planner/app2/planners/common"
	"github.com/kudrykv/latex-yearly-planner/app2/planners/mos"
)

const (
	MonthsOnSidesTemplate = "mos"
)

var UnknownTemplateName = errors.New("unknown planner name")
var UnknownSectionErr = errors.New("unknown section")

type Planner[T any] struct {
	params      common.Params[T]
	futureFiles futureFiles
	dir         string
	builder     mos.MonthsOnSides
	layout      common.Layout
}

func New[T any](template string, params common.Params[T]) (*Planner[T], error) {
	var builder mos.MonthsOnSides

	layout, err := params.Layout()
	if err != nil {
		return nil, fmt.Errorf("layout: %w", err)
	}

	switch template {
	case MonthsOnSidesTemplate:
		builder, _ = mos.New(params)
	default:
		return nil, fmt.Errorf("%s: %w", template, UnknownTemplateName)
	}

	return &Planner[T]{
		params:      params,
		layout:      layout,
		futureFiles: nil,
		dir:         "",
		builder:     builder,
	}, nil
}

func (r *Planner[T]) Generate() error {
	if err := r.builder.PrepareDetails(r.layout); err != nil {
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

func (r *Planner[T]) WriteTeXTo(dir string) error {
	for _, future := range r.futureFiles {
		if err := os.WriteFile(path.Join(dir, future.name+".tex"), future.buffer.Bytes(), 0600); err != nil {
			return fmt.Errorf("write file %s: %w", future.name, err)
		}
	}

	r.dir = dir

	return nil
}

func (r *Planner[T]) Compile(ctx context.Context) error {
	for i := 0; i < r.builder.RunTimes(); i++ {
		cmd := exec.CommandContext(ctx, "pdflatex", "./document.tex")
		cmd.Dir = r.dir

		if err := cmd.Run(); err != nil {
			return fmt.Errorf("run pdflatex: %w", err)
		}
	}

	return nil
}

func (r *Planner[T]) createRootDocument() error {
	buffer, err := newDocument(r).createBuffer()
	if err != nil {
		return fmt.Errorf("create buffer: %w", err)
	}

	r.futureFiles = append(r.futureFiles, futureFile{name: "document", buffer: buffer})

	return nil
}
