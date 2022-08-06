package mos

import (
	"fmt"

	"github.com/kudrykv/latex-yearly-planner/app2/types"
)

type Planner struct {
	layout     types.Layout
	parameters Parameters
}

func New(layout types.Layout) (*Planner, error) {
	parameters, ok := layout.Misc.(Parameters)
	if !ok {
		return nil, fmt.Errorf("expected Parameters, got %T", layout.Misc)
	}

	return &Planner{
		layout:     layout,
		parameters: parameters,
	}, nil
}
