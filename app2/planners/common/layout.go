package common

import (
	"errors"
	"fmt"

	"github.com/kudrykv/latex-yearly-planner/app2/types"
)

type MainHand int

var UnknownDeviceErr = errors.New("unknown device")

const (
	LeftHand MainHand = iota + 1
	RightHand
)

func newLayout(deviceName string, hand MainHand) (Layout, error) {
	switch deviceName {
	case "supernote_a5x":
		var layout = Layout{
			Name: deviceName,
			Hand: hand,
			Paper: Paper{
				Width:  156,
				Height: 230,
			},
			Margin: Margin{
				Top:    10,
				Right:  5,
				Bottom: 5,
				Left:   10,
			},
			MarginNotes: MarginNotes{
				Margin:  2,
				Width:   8,
				Reverse: `\reversemarginpar`,
			},
			Sizes: Sizes{
				TwoColumnsSeparatorSize:   5,
				ThreeColumnsSeparatorSize: 5,
			},
		}
		if hand == LeftHand {
			layout.Margin.Left, layout.Margin.Right = layout.Margin.Right, layout.Margin.Left
			layout.MarginNotes.Reverse = ""
		}

		return layout, nil
	default:
		return Layout{}, fmt.Errorf("unknown device %s: %w", deviceName, UnknownDeviceErr)
	}
}

type Paper struct {
	Width  types.Millimeters
	Height types.Millimeters
}

type Layout struct {
	Name string

	Hand        MainHand
	Paper       Paper
	Margin      Margin
	MarginNotes MarginNotes
	Sizes       Sizes
}

type Margin struct {
	Top    types.Millimeters
	Right  types.Millimeters
	Bottom types.Millimeters
	Left   types.Millimeters
}

type MarginNotes struct {
	Margin  types.Millimeters
	Width   types.Millimeters
	Reverse string
}

type Sizes struct {
	TwoColumnsSeparatorSize   types.Millimeters
	ThreeColumnsSeparatorSize types.Millimeters
}
