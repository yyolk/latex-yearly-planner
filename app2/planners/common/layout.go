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
				Width:  "15.6cm",
				Height: "23cm",
			},
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
	Width  string
	Height string
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

type Sizes struct {
	TwoColumnsSeparatorSize   string
	ThreeColumnsSeparatorSize string
}
