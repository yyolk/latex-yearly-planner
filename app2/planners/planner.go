package planners

import (
	"errors"
	"fmt"

	"github.com/kudrykv/latex-yearly-planner/app2/devices"
)

type Planner interface {
	GenerateFor(device devices.Device) error
}

var UnknownTemplateName = errors.New("unknown planner name")

func New(name string) (Planner, error) {
	switch name {
	case "mos":
		return &MonthsOnSides{}, nil
	default:
		return nil, fmt.Errorf("%s: %w", name, UnknownTemplateName)
	}
}

type MonthsOnSides struct{}

func (r *MonthsOnSides) GenerateFor(device devices.Device) error {
	return nil
}
