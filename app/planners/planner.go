package planners

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/kudrykv/latex-yearly-planner/app/planners/mos"
	"github.com/kudrykv/latex-yearly-planner/app/types"
)

type PlannerBuilder interface {
	Generate() (types.NamedBuffers, error)
	RunTimes() int
	Document() types.Document
}

type Planner struct {
	builder PlannerBuilder

	futureFiles types.NamedBuffers
}

var ErrInvalidParameters = errors.New("invalid parameters")

func New(template string, layout any) (*Planner, error) {
	var builder PlannerBuilder

	switch template {
	case "mos":
		parameters, ok := layout.(mos.Parameters)
		if !ok {
			return nil, fmt.Errorf("%w: want mos.Parameters, got %T: ", ErrInvalidParameters, layout)
		}

		builder = mos.New(parameters)
	}

	return &Planner{
		builder: builder,
	}, nil
}

func (r *Planner) Generate() error {
	var err error
	if r.futureFiles, err = r.builder.Generate(); err != nil {
		return fmt.Errorf("generate %T: %w", r.builder, err)
	}

	if err := r.createRootDocument(); err != nil {
		return fmt.Errorf("create root document: %w", err)
	}

	return nil
}

func (r *Planner) createRootDocument() error {
	buffer, err := newDocument(r).build()
	if err != nil {
		return fmt.Errorf("build document: %w", err)
	}

	r.futureFiles = append(r.futureFiles, types.NamedBuffer{Name: "document", Buffer: buffer})

	return nil
}

func (r *Planner) WriteTeXTo(dir string) error {
	for _, futureFile := range r.futureFiles {
		filepath := path.Join(dir, futureFile.Name+".tex")
		bufferBytes := futureFile.Buffer.Bytes()

		if err := os.WriteFile(filepath, bufferBytes, 0644); err != nil {
			return fmt.Errorf("write file %s: %w", futureFile.Name, err)
		}
	}

	return nil
}

func (r *Planner) Compile(ctx context.Context, dir string) error {
	for i := 0; i < r.builder.RunTimes(); i++ {
		cmd := exec.CommandContext(ctx, "pdflatex", "./document.tex")
		cmd.Dir = dir

		if err := cmd.Run(); err != nil {
			return fmt.Errorf("run pdflatex: %w", err)
		}
	}

	return nil
}

func (r *Planner) Files() string {
	pieces := make([]string, 0, len(r.futureFiles))

	for _, futureFile := range r.futureFiles {
		pieces = append(pieces, fmt.Sprintf(`\include{%s}`, futureFile.Name))
	}

	return strings.Join(pieces, "\n")
}
