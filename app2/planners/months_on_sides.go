package planners

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/kudrykv/latex-yearly-planner/app2/pages"
	"github.com/kudrykv/latex-yearly-planner/app2/tex/cell"
	"github.com/kudrykv/latex-yearly-planner/app2/texsnippets"
	"github.com/kudrykv/latex-yearly-planner/lib/calendar"
)

type MonthsOnSides struct {
	templateData TemplateData
}

func newMonthOnSides(templateData TemplateData) MonthsOnSides {
	return MonthsOnSides{
		templateData: templateData,
	}
}

type sectionFunc func() (*bytes.Buffer, error)

func (r MonthsOnSides) Sections() map[string]sectionFunc {
	return map[string]sectionFunc{
		TitleSection:       r.titleSection,
		AnnualSection:      r.annualSection,
		QuarterliesSection: r.quarterliesSection,
		MonthliesSection:   r.monthliesSection,
		WeekliesSection:    r.weekliesSection,
		DailiesSection:     r.dailiesSection,
	}
}

func (r *MonthsOnSides) titleSection() (*bytes.Buffer, error) {
	buffer := &bytes.Buffer{}

	if err := texsnippets.Execute(buffer, texsnippets.Title, r.templateData); err != nil {
		return nil, fmt.Errorf("execute template title: %w", err)
	}

	return buffer, nil
}

func (r *MonthsOnSides) annualSection() (*bytes.Buffer, error) {
	templateData := r.templateData
	year := calendar.NewYear(templateData.Year(), templateData.Weekday())

	rightItems := cell.Cells{cell.New("Calendar").Ref(), cell.New("To Do"), cell.New("Notes")}
	header := newMOSAnnualHeader(
		templateData.Layout(),
		headerWithYear(year),
		headerWithLeft(strconv.Itoa(year.Year())),
		headerWithRight(rightItems.Slice()),
	)

	buffer, err := writeToBuffer(&bytes.Buffer{}, header, mosAnnualContents{year: year})
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

	year := calendar.NewYear(r.templateData.Year(), r.templateData.Weekday())

	rightItems := cell.Cells{cell.New("Calendar"), cell.New("To Do"), cell.New("Notes")}

	for _, quarter := range year.Quarters {
		header := newMOSAnnualHeader(
			r.templateData.Layout(),
			headerWithYear(year),
			headerWithLeft(strconv.Itoa(year.Year())),
			headerWithRight(rightItems.Slice()),

			headerSelectQuarter(quarter),
		)

		buffer, err = writeToBuffer(buffer, header, mosQuarterlyContents{quarter: quarter})
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

	year := calendar.NewYear(r.templateData.Year(), r.templateData.Weekday())

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

	year := calendar.NewYear(r.templateData.Year(), r.templateData.Weekday())
	weeks := year.InWeeks()

	for _, week := range weeks {
		buffer, err = writeToBuffer(buffer, mosWeeklyHeader{year: year}, mosWeeklyContents{week: week})
		if err != nil {
			return nil, fmt.Errorf("write to buffer: %w", err)
		}
	}

	return buffer, nil
}

func (r *MonthsOnSides) dailiesSection() (*bytes.Buffer, error) {
	var (
		buffer = &bytes.Buffer{}
		err    error
	)

	year := calendar.NewYear(r.templateData.Year(), r.templateData.Weekday())

	for _, day := range year.Days() {
		buffer, err = writeToBuffer(buffer, mosDailyHeader{year: year}, mosDailyContents{day: day})
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
