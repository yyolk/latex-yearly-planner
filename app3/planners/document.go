package planners

import (
	"bytes"
	"fmt"
	"text/template"
)

type document struct {
	template *template.Template
	planner  *Planner
}

func newDocument(planner *Planner) document {
	return document{
		template: template.Must(template.New("document").Parse(documentTex)),
		planner:  planner,
	}
}

func (r document) build() (*bytes.Buffer, error) {
	buffer := bytes.NewBuffer(make([]byte, 0, len(documentTex)))

	data := map[string]any{
		"Files":    r.planner.Files(),
		"Document": r.planner.builder.Document(),
	}

	if err := r.template.ExecuteTemplate(buffer, "document", data); err != nil {
		return nil, fmt.Errorf("executing template: %w", err)
	}

	return buffer, nil
}

const documentTex = `\documentclass[9pt]{extarticle}

\usepackage{{ if .Document.Debug.ShowFrames }}[showframe]{{end}}{geometry}
\usepackage[table]{xcolor}
\usepackage{tabularx}
\usepackage{hyperref}
\usepackage{marginnote}
\usepackage{adjustbox}
\usepackage{multido}
\usepackage{amssymb}

\hypersetup{ {{- if not .Document.Debug.ShowLinks}}hidelinks=true{{end -}} }

\geometry{paperwidth={{.Document.Screen.Width}}, paperheight={{.Document.Screen.Height}}}
\geometry{
             top={{ .Document.Margin.Top }},
          bottom={{ .Document.Margin.Bottom }},
            left={{ .Document.Margin.Left }},
           right={{ .Document.Margin.Right }},
  marginparwidth={{ .Document.MarginNotes.Width }},
    marginparsep={{ .Document.MarginNotes.Separator }}
}

{{ if .Document.MarginNotes.Reverse -}}
	\reversemarginpar
{{- end }}

\pagestyle{empty}
\newcolumntype{Y}{>{\centering\arraybackslash}X}
\parindent=0pt
\fboxsep0pt

\newcommand{\remainingHeight}{%
  \ifdim\pagegoal=\maxdimen
  \dimexpr\textheight\relax
  \else
  \dimexpr\pagegoal-\pagetotal-\lineskip\relax
  \fi%
}

\newcommand{\myDotGrid}[2]{\vskip5mm\leavevmode\multido{\dC=0mm+5mm}{#1}{\multido{\dR=0mm+5mm}{#2}{\put(\dR,\dC){\scriptsize.}}}}

\newlength{\myLenLineThicknessDefault}
\newlength{\myLenLineThicknessThick}
\setlength{\myLenLineThicknessDefault}{.4pt}
\setlength{\myLenLineThicknessThick}{.8pt}

\newcommand{\myLinePlain}{\hrule width \linewidth height \myLenLineThicknessDefault}
\newcommand{\myLineThick}{\hrule width \linewidth height \myLenLineThicknessThick}

\newcommand{\myLineColor}[1]{\textcolor{#1}{\myLinePlain}}

\newcommand{\myColorGray}{gray}
\newcommand{\myColorLightGray}{gray!50}

\newcommand{\myLineGray}{\myLineColor{\myColorGray}}
\newcommand{\myLineLightGray}{\myLineColor{\myColorLightGray}}

\newcommand{\myUnderline}[1]{#1\vskip1mm\myLineThick\par}

\begin{document}

{{ .Files }}

\end{document}`
