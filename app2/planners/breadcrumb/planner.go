package breadcrumb

import (
	"bytes"
	"fmt"
	"time"

	"github.com/imdario/mergo"
	"github.com/kudrykv/latex-yearly-planner/app2/pages"
	"github.com/kudrykv/latex-yearly-planner/app2/planners/common"
	"github.com/kudrykv/latex-yearly-planner/app2/planners/common/contents"
	"github.com/kudrykv/latex-yearly-planner/app2/types"
	"github.com/kudrykv/latex-yearly-planner/lib/texcalendar"
)

type Planner struct {
	layout common.Layout

	ui      UI
	weekday time.Weekday
	year    texcalendar.Year
}

func New[T any](params common.Params[T]) (*Planner, error) {
	ui, ok := any(params.UI).(UI)
	if !ok {
		return nil, fmt.Errorf("invalid UI: %T", params.UI)
	}

	return &Planner{
		ui: ui,

		weekday: params.Weekday,

		year: texcalendar.NewYear(params.Year, texcalendar.WithParameters(texcalendar.Parameters{
			Hand:       params.Hand,
			Weekday:    params.Weekday,
			FirstMonth: time.January,
		})),
	}, nil
}

func (r *Planner) PrepareDetails(layout common.Layout) error {
	r.layout = layout

	var err error
	if r.ui, err = newUI(r.layout, r.ui); err != nil {
		return fmt.Errorf("new UI: %w", err)
	}

	r.year.Apply(
		texcalendar.WithLittleCalArrayStretch(r.ui.LittleCalArrayStretch),
		texcalendar.WithLargeCalHeaderHeight(r.ui.LargeCalHeaderHeight),
	)

	return nil
}

func newUI(layout common.Layout, ui UI) (UI, error) {
	var finalUI UI

	switch layout.Name {
	case "supernote_a5x":
		finalUI = UI{
			LittleCalArrayStretch: 1.6,
			LargeCalHeaderHeight:  5,
		}

		if err := mergo.Merge(&finalUI, ui); err != nil {
			return finalUI, fmt.Errorf("merge: %w", err)
		}

	default:
		return finalUI, fmt.Errorf("unknown layout: %s", layout.Name)
	}

	return finalUI, nil
}

func (r *Planner) Sections() map[string]types.SectionFunc {
	return map[string]types.SectionFunc{
		common.TitleSection:  r.titleSection,
		common.AnnualSection: r.annualSection,
	}
}

func (r *Planner) RunTimes() int {
	return 1
}

func (r *Planner) titleSection() (*bytes.Buffer, error) {
	buffer := pages.NewBuffer()

	if err := buffer.WriteBlocks(contents.NewTitle(r.year.Name())); err != nil {
		return nil, fmt.Errorf("write to buffer: %w", err)
	}

	return buffer.Buffer, nil
}

func (r *Planner) annualSection() (*bytes.Buffer, error) {
	buffer := pages.NewBuffer()

	if err := buffer.WriteBlocks(contents.NewAnnual(r.year)); err != nil {
		return nil, fmt.Errorf("write to buffer: %w", err)
	}

	return buffer.Buffer, nil
}
