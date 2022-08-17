package planners

import (
	"context"
)

type Planner struct {
}

func New(template string, layout any) (*Planner, error) {
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
