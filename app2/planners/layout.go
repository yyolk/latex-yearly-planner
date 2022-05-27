package planners

import (
	"fmt"

	"github.com/kudrykv/latex-yearly-planner/app2/devices"
)

func newLayout(device devices.Device) (Layout, error) {
	switch device.(type) {
	case *devices.SupernoteA5X:
		return Layout{
			Margin: Margin{
				Top:    "1cm",
				Right:  "5mm",
				Bottom: "5mm",
				Left:   "1cm",
			},
			MarginNotes: MarginNotes{
				Margin: "3mm",
				Width:  "1cm",
			},
		}, nil
	default:
		return Layout{}, fmt.Errorf("unknown device type %T: %w", device, UnknownDeviceTypeErr)
	}
}

type Layout struct {
	Margin      Margin
	MarginNotes MarginNotes
}

type Margin struct {
	Top    string
	Right  string
	Bottom string
	Left   string
}

type MarginNotes struct {
	Margin string
	Width  string
}
