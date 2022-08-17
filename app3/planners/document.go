package planners

import (
	"bytes"
	"text/template"
)

type document struct {
	template *template.Template
}

func newDocument(planner *Planner) document {
	return document{
		template: template.Must(template.New("document").Parse(documentTex)),
	}
}

func (r document) build() (*bytes.Buffer, error) {
	buffer := bytes.NewBuffer(make([]byte, 0, len(documentTex)))

	if err := r.template.ExecuteTemplate(buffer, "document", r); err != nil {
		return nil, err
	}

	return buffer, nil
}

const documentTex = `\documentclass[9pt]{extarticle}

\usepackage{{ if .Layout.Debug.ShowFrames }}[showframe]{{end}}{geometry}
\usepackage[table]{xcolor}
\usepackage{tabularx}
\usepackage{hyperref}
\usepackage{marginnote}
\usepackage{adjustbox}
\usepackage{multido}
\usepackage{amssymb}

\hypersetup{
    {{- if not .Layout.Debug.ShowLinks}}hidelinks=true{{end -}}
}

\geometry{paperwidth={{.Layout.Paper.Width}}, paperheight={{.Layout.Paper.Height}}}
\geometry{
             top={{ .Layout.Margin.Top }},
          bottom={{ .Layout.Margin.Bottom }},
            left={{ .Layout.Margin.Left }},
           right={{ .Layout.Margin.Right }},
  marginparwidth={{ .Layout.MarginNotes.Width }},
    marginparsep={{ .Layout.MarginNotes.Margin }}
}

{{ if .Layout.MarginNotes.Reverse -}}
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

\newcommand{\myDotGrid}[2]{\leavevmode\multido{\dC=0mm+5mm}{#1}{\multido{\dR=0mm+5mm}{#2}{\put(\dR,\dC){\circle*{0.1}}}}}

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
