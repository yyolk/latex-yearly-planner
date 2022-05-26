package planners

import (
	"errors"
	"fmt"

	"github.com/kudrykv/latex-yearly-planner/app2/devices"
)

const (
	MonthsOnSidesTemplate = "mos"
)

type Planner interface {
	GenerateFor(device devices.Device) error
}

var UnknownTemplateName = errors.New("unknown planner name")

func New(params Params) (Planner, error) {
	switch params.Name {
	case MonthsOnSidesTemplate:
		return &MonthsOnSides{params: params}, nil
	default:
		return nil, fmt.Errorf("%s: %w", params.Name, UnknownTemplateName)
	}
}
