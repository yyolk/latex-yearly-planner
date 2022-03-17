package planners

import (
	"errors"
	"fmt"
)

type Planner interface {
	GenerateFiles(dir string) error
	Compile(dir string) error
}

type Params struct {
	Name string
}

var ErrUnknownPlanner = errors.New("unknown planner")

func New(params Params) (Planner, error) { //nolint:ireturn
	switch params.Name { // nolint:gocritic
	case "breadcrumb":
		return Breadcrumb{}, nil
	}

	return nil, fmt.Errorf("%s: %w", params.Name, ErrUnknownPlanner)
}
