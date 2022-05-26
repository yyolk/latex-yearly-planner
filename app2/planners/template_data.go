package planners

import (
	"errors"
	"fmt"
	"strings"

	"github.com/kudrykv/latex-yearly-planner/app2/devices"
)

type TemplateData struct {
	year   int
	files  []string
	device devices.Device
	layout Layout
}

var UnknownDeviceTypeErr = errors.New("unknown device type")

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
		}, nil
	default:
		return Layout{}, fmt.Errorf("unknown device type %T: %w", device, UnknownDeviceTypeErr)
	}
}

type Layout struct {
	Margin Margin
}

type Margin struct {
	Top    string
	Right  string
	Bottom string
	Left   string
}

type ApplyToTemplateData func(*TemplateData)

func (r *TemplateData) Apply(options ...ApplyToTemplateData) *TemplateData {
	for _, option := range options {
		option(r)
	}

	return r
}

func WithFiles(files ...string) ApplyToTemplateData {
	return func(data *TemplateData) {
		data.files = files
	}
}

func WithLayout(layout Layout) ApplyToTemplateData {
	return func(data *TemplateData) {
		data.layout = layout
	}
}

func (r TemplateData) Year() int {
	return r.year
}

func (r TemplateData) Device() devices.Device {
	return r.device
}

func (r TemplateData) Layout() Layout {
	return r.layout
}

func (r TemplateData) MarginNotesWidth() string {
	return "1cm"
}

func (r TemplateData) MarginNotesMargin() string {
	return "3mm"
}

func (r TemplateData) Pages() string {
	include := func(str string) string { return `\include{` + str + `}` }

	var wrappedPages []string
	for _, page := range r.files {
		wrappedPages = append(wrappedPages, include(page))
	}

	return strings.Join(wrappedPages, "\n")
}
