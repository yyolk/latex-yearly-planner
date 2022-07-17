package common

import (
	"errors"
	"fmt"
)

type MainHand int

var UnknownDeviceErr = errors.New("unknown device")

const (
	LeftHand MainHand = iota + 1
	RightHand
)

func NewLayout(deviceName string, hand MainHand) (Layout, error) {
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
				Margin:  "2mm",
				Width:   "8mm",
				Reverse: `\reversemarginpar`,
			},
			Sizes: Sizes{
				TwoColumnsSeparatorSize:   "5mm",
				ThreeColumnsSeparatorSize: "5mm",
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
	Width  Millimeters
	Height Millimeters
}

type Millimeters float64

func (r Millimeters) String() string {
	return fmt.Sprintf("%.4fmm", r)
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
	Top    Millimeters
	Right  Millimeters
	Bottom Millimeters
	Left   Millimeters
}

type MarginNotes struct {
	Margin  string
	Width   string
	Reverse string
}

type Sizes struct {
	TwoColumnsSeparatorSize   string
	ThreeColumnsSeparatorSize string
}
