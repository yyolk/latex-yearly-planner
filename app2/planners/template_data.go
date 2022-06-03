package planners

import (
	"errors"
	"strings"
	"time"

	"github.com/kudrykv/latex-yearly-planner/app2/devices"
)

const (
	TitleSection       = "title"
	AnnualSection      = "annual"
	QuarterliesSection = "quarterlies"
	MonthliesSection   = "monthlies"
)

type TemplateData struct {
	year     int
	files    []string
	device   devices.Device
	layout   Layout
	sections []string
}

var UnknownDeviceTypeErr = errors.New("unknown device type")

type ApplyToTemplateData func(*TemplateData)

func (r *TemplateData) Apply(options ...ApplyToTemplateData) {
	for _, option := range options {
		option(r)
	}
}

func WithFiles(files []string) ApplyToTemplateData {
	return func(data *TemplateData) {
		data.files = files
	}
}

func WithLayout(layout Layout) ApplyToTemplateData {
	return func(data *TemplateData) {
		data.layout = layout
	}
}

func WithYear(year int) ApplyToTemplateData {
	return func(data *TemplateData) {
		data.year = year
	}
}

func WithSections(sections []string) ApplyToTemplateData {
	return func(data *TemplateData) {
		data.sections = sections
	}
}

func WithDevice(device devices.Device) ApplyToTemplateData {
	return func(data *TemplateData) {
		data.device = device
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

func (r TemplateData) Files() string {
	include := func(str string) string { return `\include{` + str + `}` }

	var wrappedPages []string
	for _, page := range r.files {
		wrappedPages = append(wrappedPages, include(page))
	}

	return strings.Join(wrappedPages, "\n")
}

func (r *TemplateData) Weekday() time.Weekday {
	return time.Monday
}
