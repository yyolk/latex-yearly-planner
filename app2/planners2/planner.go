package planners2

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.com/kudrykv/latex-yearly-planner/app2/planners2/mos"
	"github.com/kudrykv/latex-yearly-planner/app2/types"
)

type Planner struct {
	layout      types.Layout
	builder     *mos.Planner
	workdir     string
	futureFiles types.NamedDatas
}

func New(templateName string, layout types.Layout) (*Planner, error) {
	var (
		builder *mos.Planner
		err     error
	)

	switch templateName {
	case "mos":
		builder, err = mos.New(layout)
	}

	if err != nil {
		return nil, fmt.Errorf("new builder %s: %w", templateName, err)
	}

	return &Planner{
		layout:  layout,
		builder: builder,
	}, nil
}

func (r *Planner) Generate() error {
	var err error
	if r.futureFiles, err = r.builder.BuildData(); err != nil {
		return fmt.Errorf("build data: %w", err)
	}

	if err := r.createRootDocument(); err != nil {
		return fmt.Errorf("create root document: %w", err)
	}

	return nil
}

func (r *Planner) WriteTeXTo(workdir string) error {
	for _, future := range r.futureFiles {
		if err := os.WriteFile(path.Join(workdir, future.Name+".tex"), future.Data, 0600); err != nil {
			return fmt.Errorf("write file %s: %w", future.Name, err)
		}
	}

	r.workdir = workdir

	return nil
}

func (r *Planner) Compile(ctx context.Context) error {
	for i := 0; i < r.builder.RunTimes(); i++ {
		cmd := exec.CommandContext(ctx, "pdflatex", "./document.tex")
		cmd.Dir = r.workdir

		if err := cmd.Run(); err != nil {
			return fmt.Errorf("run pdflatex: %w", err)
		}
	}

	return nil
}

func (r *Planner) createRootDocument() error {
	data, err := newDocument(r).build()
	if err != nil {
		return fmt.Errorf("build: %w", err)
	}

	r.futureFiles.Append("document", data)

	return nil
}
