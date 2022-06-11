package planners

import (
	"context"
	"errors"
	"fmt"

	"github.com/kudrykv/latex-yearly-planner/app2/devices"
)

const (
	MonthsOnSidesTemplate = "mos"
)

type Planner interface {
	GenerateFor(devices.Device) error
	WriteTeXTo(dir string) error
	Compile(context.Context) error
}

var UnknownTemplateName = errors.New("unknown planner name")

func New(params Params) (*MonthsOnSides, error) {
	var builder monthsOnSidesBuilder

	switch params.Name {
	case MonthsOnSidesTemplate:
		builder = newMonthOnSidesBuilder(params.TemplateData)
	default:
		return nil, fmt.Errorf("%s: %w", params.Name, UnknownTemplateName)
	}

	return &MonthsOnSides{
		params:      params,
		futureFiles: nil,
		dir:         "",
		builder:     builder,
	}, nil
}
