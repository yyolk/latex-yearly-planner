package planners

import (
	"fmt"
	"text/template"
)

func createTemplates() (*template.Template, error) {
	var (
		tpls = template.New("")
		err  error
	)

	for _, row := range templatesToCompile {
		tpls = tpls.New(row[0])

		if tpls, err = tpls.Parse(row[1]); err != nil {
			return nil, fmt.Errorf("parse %s: %w", row[1], err)
		}
	}

	return tpls, nil
}

var templatesToCompile = [][]string{
	{titleTpl, titleTex},
	{rootDocumentTpl, rootDocumentTex},
}

const titleTpl = "title"
const titleTex = `\hspace{0pt}\vfil
\hfill\resizebox{.7\linewidth}{!}{ {{- .Year -}} }%
\pagebreak`

const rootDocumentTpl = "root-document"
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
             top={{ .Layout.Margin.Top }},
          bottom={{ .Layout.Margin.Bottom }},
            left={{ .Layout.Margin.Left }},
           right={{ .Layout.Margin.Right }},
  marginparwidth={{ .Layout.MarginNotes.Width }},
    marginparsep={{ .Layout.MarginNotes.Margin }}
}

\pagestyle{empty}
\newcolumntype{Y}{>{\centering\arraybackslash}X}
\parindent=0pt
\fboxsep0pt

\begin{document}

{{ .Files }}

\end{document}`
