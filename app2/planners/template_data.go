package planners

import (
	"strings"

	"github.com/kudrykv/latex-yearly-planner/app2/devices"
)

type TemplateData struct {
	Year   int
	files  []string
	device devices.Device
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

func (r TemplateData) Device() devices.Device {
	return r.device
}

func (r TemplateData) TopMargin() string {
	return "1.4cm"
}

func (r TemplateData) BottomMargin() string {
	return "5mm"
}

func (r TemplateData) LeftMargin() string {
	return "5mm"
}

func (r TemplateData) RightMargin() string {
	return "5mm"
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
