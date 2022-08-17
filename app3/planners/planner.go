package planners

import (
	"context"

	"github.com/kudrykv/latex-yearly-planner/app2/types"
)

type Planner struct {
}

func New(template string, layout types.Layout) (*Planner, error) {
	return &Planner{}, nil
}

func (r *Planner) Generate() error {
	panic("not implemented")
}

func (r *Planner) WriteTeXTo(dir string) error {
	panic("not implemented")
}

func (r *Planner) Compile(ctx context.Context) error {
	panic("not implemented")
}
