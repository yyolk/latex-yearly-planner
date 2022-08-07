package planners2

import (
	"fmt"

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

	return nil
}
