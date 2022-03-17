package planners

import (
	"errors"
	"fmt"
)

type planner interface {
	GenerateFiles(dir string) error
	Compile(dir string) error
}

type Planner struct {
	planner planner
}

type Params struct {
	BreadcrumbParams BreadcrumbParams
	MOSParams        MOSParams
	Name             string
}

var ErrUnknownPlanner = errors.New("unknown planner")

func New(params Params) (Planner, error) {
	switch params.Name {
	case "breadcrumb":
		return Planner{planner: breadcrumb{params: params.BreadcrumbParams}}, nil
	case "mos":
		return Planner{planner: mos{params: params.MOSParams}}, nil
	}

	return Planner{}, fmt.Errorf("%s: %w", params.Name, ErrUnknownPlanner)
}

func (p Planner) GenerateFiles(dir string) error {
	return p.planner.GenerateFiles(dir) //nolint:wrapcheck
}

func (p Planner) Compile(dir string) error {
	return p.planner.Compile(dir) //nolint:wrapcheck
}
