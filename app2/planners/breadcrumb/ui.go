package breadcrumb

import (
	"fmt"

	"github.com/imdario/mergo"
	"github.com/kudrykv/latex-yearly-planner/app2/planners/common"
	"github.com/kudrykv/latex-yearly-planner/app2/types"
)

type UI struct {
	LittleCalArrayStretch types.Spring
	LargeCalHeaderHeight  types.Millimeters
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
