package planners

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"path"
	"text/template"

	"github.com/kudrykv/latex-yearly-planner/app2/devices"
)

type MonthsOnSides struct {
	params      Params
	futureFiles []futureFile
	templates   *template.Template
}

type futureFile struct {
	name   string
	buffer *bytes.Buffer
}

func newMonthsOnSides(params Params) (*MonthsOnSides, error) {
	mos := &MonthsOnSides{params: params}

	if err := mos.init(); err != nil {
		return nil, fmt.Errorf("init: %w", err)
	}

	return mos, nil
}

func (r *MonthsOnSides) GenerateFor(device devices.Device) error {
	layout, err := newLayout(device)
	if err != nil {
		return fmt.Errorf("new layout: %w", err)
	}

	r.params.TemplateData.Apply(WithLayout(layout), WithDevice(device))

	if err := r.createTitle(); err != nil {
		return fmt.Errorf("create title: %w", err)
	}

	if err := r.createRootDocument(); err != nil {
		return fmt.Errorf("create root document: %w", err)
	}

	return nil
}

func WithDevice(device devices.Device) ApplyToTemplateData {
	return func(data *TemplateData) {
		data.device = device
	}
}

func (r *MonthsOnSides) WriteTo(dir string) error {
	for _, future := range r.futureFiles {
		if err := os.WriteFile(path.Join(dir, future.name), []byte(future.buffer.String()), 0600); err != nil {
			return fmt.Errorf("write file %s: %w", future.name, err)
		}
	}

	return nil
}

func (r *MonthsOnSides) Compile(ctx context.Context) error {
	cmd := exec.CommandContext(ctx, "pdflatex", "./document.tex")
	cmd.Dir = "./out"

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("run pdflatex: %w", err)
	}

	return nil
}

func (r *MonthsOnSides) init() error {
	var err error

	if r.templates, err = bigTemplateStuff(); err != nil {
		return fmt.Errorf("big template stuff: %w", err)
	}

	return nil
}

func (r *MonthsOnSides) createTitle() error {
	buffer := &bytes.Buffer{}

	if err := r.templates.ExecuteTemplate(buffer, "title", r.params.TemplateData); err != nil {
		return fmt.Errorf("execute template title: %w", err)
	}

	r.futureFiles = append(r.futureFiles, futureFile{name: "title.tex", buffer: buffer})

	return nil
}

func (r *MonthsOnSides) createRootDocument() error {
	files := make([]string, 0, len(r.futureFiles))

	for _, file := range r.futureFiles {
		files = append(files, file.name)
	}

	buffer := &bytes.Buffer{}

	if err := r.templates.ExecuteTemplate(buffer, "root-document", r.params.TemplateData.Apply(WithFiles(files...))); err != nil {
		return fmt.Errorf("execute template root-document: %w", err)
	}

	r.futureFiles = append(r.futureFiles, futureFile{name: "document.tex", buffer: buffer})

	return nil
}

func bigTemplateStuff() (*template.Template, error) {
	tpls := template.New("")

	var err error

	for _, row := range [][]string{{"title", titleTex}, {"root-document", rootDocumentTex}} {
		tpls = tpls.New(row[0])

		if tpls, err = tpls.Parse(row[1]); err != nil {
			return nil, fmt.Errorf("parse %s: %w", row[1], err)
		}
	}

	return tpls, nil
}

const titleTex = `\hspace{0pt}\vfil
\hfill\resizebox{.7\linewidth}{!}{ {{- .Year -}} }%
\pagebreak`

const rootDocumentTex = `\documentclass[9pt]{extarticle}

\usepackage{geometry}
\usepackage[table]{xcolor}
\usepackage{calc}
\usepackage{dashrule}
\usepackage{setspace}
\usepackage{array}
\usepackage{tikz}
\usepackage{varwidth}
\usepackage{blindtext}
\usepackage{tabularx}
\usepackage{wrapfig}
\usepackage{makecell}
\usepackage{graphicx}
\usepackage{multirow}
\usepackage{amssymb}
\usepackage{expl3}
\usepackage{leading}
\usepackage{pgffor}
\usepackage{hyperref}
\usepackage{marginnote}
\usepackage{adjustbox}
\usepackage{multido}


\geometry{paperwidth={{.Device.Paper.Width}}, paperheight={{.Device.Paper.Height}}}
\geometry{
             top={{ .TopMargin }},
          bottom={{ .BottomMargin }},
            left={{ .LeftMargin }},
           right={{ .RightMargin }},
  marginparwidth={{ .MarginNotesWidth }},
    marginparsep={{ .MarginNotesMargin }}
}

\pagestyle{empty}
\newcolumntype{Y}{>{\centering\arraybackslash}X}
\parindent=0pt
\fboxsep0pt

\begin{document}

{{ .Pages }}

\end{document}`
