package common

import (
	"errors"
	"fmt"

	"github.com/kudrykv/latex-yearly-planner/app2/devices"
)

type MainHand int

var UnknownDeviceTypeErr = errors.New("unknown device type")

const (
	LeftHand MainHand = iota + 1
	RightHand
)

func NewLayout(device devices.Device, hand MainHand) (Layout, error) {
	switch device.(type) {
	case *devices.SupernoteA5X:
		layout := Layout{
			Hand: hand,
			Margin: Margin{
				Top:    "1cm",
				Right:  "5mm",
				Bottom: "5mm",
				Left:   "1cm",
			},
			MarginNotes: MarginNotes{
				Margin:  "2mm",
				Width:   "8mm",
				Reverse: `\reversemarginpar`,
			},
		}

		if hand == LeftHand {
			layout.Margin.Left, layout.Margin.Right = layout.Margin.Right, layout.Margin.Left
			layout.MarginNotes.Reverse = ""
		}

		return layout, nil
	default:
		return Layout{}, fmt.Errorf("unknown device type %T: %w", device, UnknownDeviceTypeErr)
	}
}

type Layout struct {
	Hand        MainHand
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
	Margin  string
	Width   string
	Reverse string
}
