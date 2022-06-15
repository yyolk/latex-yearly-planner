package mos

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/kudrykv/latex-yearly-planner/app2/devices"
	"github.com/kudrykv/latex-yearly-planner/app2/pages"
	"github.com/kudrykv/latex-yearly-planner/app2/planners/common"
	"github.com/kudrykv/latex-yearly-planner/app2/tex/cell"
	"github.com/kudrykv/latex-yearly-planner/app2/texsnippets"
	"github.com/kudrykv/latex-yearly-planner/lib/calendar"
)

type MonthsOnSides struct {
	parameters Parameters
}

func New(params common.Params) MonthsOnSides {
	return MonthsOnSides{
		parameters: Parameters{
			year:    params.Year,
			weekday: params.Weekday,
		},
	}
}

func (r *MonthsOnSides) SetLayout(layout common.Layout) {
	r.parameters.layout = layout
}

func (r MonthsOnSides) Layout() common.Layout {
	return r.parameters.layout
}

func (r *MonthsOnSides) PrepareDetails(device devices.Device) error {
	r.parameters.device = device

	switch device.(type) {
	case *devices.SupernoteA5X:
		r.parameters.ui = mosUI{
			HeaderMarginNotesArrayStretch:  "2.042",
			HeaderMarginNotesMonthsWidth:   "15.7cm",
			HeaderMarginNotesQuartersWidth: "5cm",
			HeaderArrayStretch:             "1.8185",
		}
	default:
		return fmt.Errorf("%T: %w", device, common.UnknownDeviceTypeErr)
	}

	return nil
}

type sectionFunc func() (*bytes.Buffer, error)

func (r MonthsOnSides) Sections() map[string]sectionFunc {
	return map[string]sectionFunc{
		common.TitleSection:       r.titleSection,
		common.AnnualSection:      r.annualSection,
		common.QuarterliesSection: r.quarterliesSection,
		common.MonthliesSection:   r.monthliesSection,
		common.WeekliesSection:    r.weekliesSection,
		common.DailiesSection:     r.dailiesSection,
	}
}

func (r *MonthsOnSides) titleSection() (*bytes.Buffer, error) {
	buffer := &bytes.Buffer{}

	title := strconv.Itoa(r.parameters.year)
	if err := texsnippets.Execute(buffer, texsnippets.Title, map[string]string{"Title": title}); err != nil {
		return nil, fmt.Errorf("execute template title: %w", err)
	}

	return buffer, nil
}

func (r *MonthsOnSides) annualSection() (*bytes.Buffer, error) {
	year := calendar.NewYear(r.parameters.year, r.parameters.weekday)

	rightItems := cell.Cells{cell.New("Calendar").Ref(), cell.New("To Do"), cell.New("Notes")}
	header := newMOSAnnualHeader(
		r.parameters.layout,
		r.parameters.ui,
		headerWithYear(year),
		headerWithLeft(strconv.Itoa(year.Year())),
		headerWithRight(rightItems.Slice()),
	)

	buffer, err := writeToBuffer(&bytes.Buffer{}, header, annualContents{year: year})
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

	year := calendar.NewYear(r.parameters.year, r.parameters.weekday)

	rightItems := cell.Cells{cell.New("Calendar"), cell.New("To Do"), cell.New("Notes")}

	for _, quarter := range year.Quarters {
		header := newMOSAnnualHeader(
			r.parameters.layout,
			r.parameters.ui,
			headerWithYear(year),
			headerWithLeft(strconv.Itoa(year.Year())),
			headerWithRight(rightItems.Slice()),
			headerSelectQuarter(quarter),
		)

		buffer, err = writeToBuffer(buffer, header, quarterlyContents{quarter: quarter})
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

	year := calendar.NewYear(r.parameters.year, r.parameters.weekday)
	rightItems := cell.Cells{cell.New("Calendar"), cell.New("To Do"), cell.New("Notes")}

	for _, month := range year.Months() {
		header := newMOSAnnualHeader(
			r.parameters.layout,
			r.parameters.ui,
			headerWithYear(year),
			headerWithLeft(strconv.Itoa(year.Year())),
			headerWithRight(rightItems.Slice()),
			headerSelectMonth(month),
		)

		buffer, err = writeToBuffer(buffer, header, monthlyContents{month: month})
		if err != nil {
			return nil, fmt.Errorf("write to buffer: %w", err)
		}
	}

	return buffer, nil
}

func (r *MonthsOnSides) weekliesSection() (*bytes.Buffer, error) {
	var (
		buffer = &bytes.Buffer{}
		err    error
	)

	year := calendar.NewYear(r.parameters.year, r.parameters.weekday)
	weeks := year.InWeeks()

	for _, week := range weeks {
		buffer, err = writeToBuffer(buffer, mosWeeklyHeader{year: year}, weeklyContents{week: week})
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

	year := calendar.NewYear(r.parameters.year, r.parameters.weekday)

	for _, day := range year.Days() {
		buffer, err = writeToBuffer(buffer, mosDailyHeader{year: year}, dailyContents{day: day})
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
