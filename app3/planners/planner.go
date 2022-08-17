package planners

import (
	"context"
	"errors"
	"fmt"

	"github.com/kudrykv/latex-yearly-planner/app3/planners/mos"
	"github.com/kudrykv/latex-yearly-planner/app3/types"
)

type Planner struct {
	builder *mos.Planner

	futureFiles types.NamedBuffers
}

var ErrInvalidParameters = errors.New("invalid parameters")

func New(template string, layout any) (*Planner, error) {
	var builder *mos.Planner

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
	panic("not implemented")
}

func (r *Planner) WriteTeXTo(dir string) error {
	panic("not implemented")
}

func (r *Planner) Compile(ctx context.Context) error {
	panic("not implemented")
}
