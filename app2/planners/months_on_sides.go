package planners

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.com/kudrykv/latex-yearly-planner/app2/devices"
	"github.com/kudrykv/latex-yearly-planner/app2/pages"
	"github.com/kudrykv/latex-yearly-planner/app2/texsnippets"
	"github.com/kudrykv/latex-yearly-planner/lib/calendar"
)

type MonthsOnSides struct {
	params      Params
	futureFiles []futureFile
	dir         string
}

var UnknownSectionErr = errors.New("unknown section")

func newMonthsOnSides(params Params) (*MonthsOnSides, error) {
	mos := &MonthsOnSides{
		params: params,
	}

	return mos, nil
}

func (r *MonthsOnSides) GenerateFor(device devices.Device) error {
	layout, err := newLayout(device)
	if err != nil {
		return fmt.Errorf("new layout: %w", err)
	}

	r.params.TemplateData.Apply(WithLayout(layout), WithDevice(device))

	sections := r.sections()

	for _, name := range r.params.TemplateData.sections {
		sectionFunc, ok := sections[name]
		if !ok {
			return fmt.Errorf("%v: %w", name, UnknownSectionErr)
		}

		buffer, err := sectionFunc()
		if err != nil {
			return fmt.Errorf("%v: %w", name, err)
		}

		r.futureFiles = append(r.futureFiles, futureFile{
			name:   name,
			buffer: buffer,
		})
	}

	if err := r.createRootDocument(); err != nil {
		return fmt.Errorf("create root document: %w", err)
	}

	return nil
}

type sectionFunc func() (*bytes.Buffer, error)

func (r *MonthsOnSides) sections() map[string]sectionFunc {
	return map[string]sectionFunc{
		TitleSection:       r.createTitle,
		AnnualSection:      r.annualSection,
		QuarterliesSection: r.quarterliesSection,
		MonthliesSection:   r.monthliesSection,
		WeekliesSection:    r.weekliesSection,
	}
}

func (r *MonthsOnSides) WriteTeXTo(dir string) error {
	for _, future := range r.futureFiles {
		if err := os.WriteFile(path.Join(dir, future.name+".tex"), future.buffer.Bytes(), 0600); err != nil {
			return fmt.Errorf("write file %s: %w", future.name, err)
		}
	}

	r.dir = dir

	return nil
}

func (r *MonthsOnSides) Compile(ctx context.Context) error {
	for i := 0; i < 2; i++ {
		cmd := exec.CommandContext(ctx, "pdflatex", "./document.tex")
		cmd.Dir = r.dir

		if err := cmd.Run(); err != nil {
			return fmt.Errorf("run pdflatex: %w", err)
		}
	}

	return nil
}

func (r *MonthsOnSides) createTitle() (*bytes.Buffer, error) {
	buffer := &bytes.Buffer{}

	if err := texsnippets.Execute(buffer, texsnippets.Title, r.params.TemplateData); err != nil {
		return nil, fmt.Errorf("execute template title: %w", err)
	}

	return buffer, nil
}

func (r *MonthsOnSides) createRootDocument() error {
	files := make([]string, 0, len(r.futureFiles))

	for _, file := range r.futureFiles {
		files = append(files, file.name)
	}

	r.params.TemplateData.Apply(WithFiles(files))

	buffer := &bytes.Buffer{}
	if err := texsnippets.Execute(buffer, texsnippets.Document, r.params.TemplateData); err != nil {
		return fmt.Errorf("execute template root-document: %w", err)
	}

	r.futureFiles = append(r.futureFiles, futureFile{name: "document", buffer: buffer})

	return nil
}

func (r *MonthsOnSides) annualSection() (*bytes.Buffer, error) {
	year := calendar.NewYear(r.params.TemplateData.Year(), r.params.TemplateData.Weekday())

	buffer, err := writeToBuffer(&bytes.Buffer{}, mosAnnualHeader{year: year}, mosAnnualContents{year: year})
	if err != nil {
		return nil, fmt.Errorf("write to buffer: %w", err)
	}

	return buffer, nil
}

func (r *MonthsOnSides) quarterliesSection() (*bytes.Buffer, error) {
	var (
		buffer = &bytes.Buffer{}
		err    error
	)

	year := calendar.NewYear(r.params.TemplateData.Year(), r.params.TemplateData.Weekday())

	for _, quarter := range year.Quarters {
		buffer, err = writeToBuffer(buffer, mosQuarterlyHeader{year: year}, mosQuarterlyContents{quarter: quarter})
		if err != nil {
			return nil, fmt.Errorf("write to buffer: %w", err)
		}
	}

	return buffer, nil
}

func (r *MonthsOnSides) monthliesSection() (*bytes.Buffer, error) {
	var (
		buffer = &bytes.Buffer{}
		err    error
	)

	year := calendar.NewYear(r.params.TemplateData.Year(), r.params.TemplateData.Weekday())

	for _, quarter := range year.Quarters {
		for _, month := range quarter.Months {
			buffer, err = writeToBuffer(buffer, mosMonthlyHeader{year: year}, mosMonthlyContents{month: month})
			if err != nil {
				return nil, fmt.Errorf("write to buffer: %w", err)
			}
		}
	}

	return buffer, nil
}

func (r *MonthsOnSides) weekliesSection() (*bytes.Buffer, error) {
	var (
		buffer = &bytes.Buffer{}
		err    error
	)

	year := calendar.NewYear(r.params.TemplateData.Year(), r.params.TemplateData.Weekday())
	weeks := year.InWeeks()

	for _, week := range weeks {
		buffer, err = writeToBuffer(buffer, mosWeeklyHeader{year: year}, mosWeeklyContents{week: week})
		if err != nil {
			return nil, fmt.Errorf("write to buffer: %w", err)
		}
	}

	return buffer, nil
}

func writeToBuffer(buffer *bytes.Buffer, blocks ...pages.Block) (*bytes.Buffer, error) {
	compiledPages, err := pages.NewPage(blocks...).Build()
	if err != nil {
		return nil, fmt.Errorf("build new page: %w", err)
	}

	for _, page := range compiledPages {
		buffer.WriteString(page + "\n\n" + `\pagebreak{}` + "\n")
	}

	return buffer, nil
}
