package planners

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
  marginparwidth={{ .MarginNotesWidth }},
    marginparsep={{ .MarginNotesMargin }}
}

\pagestyle{empty}
\newcolumntype{Y}{>{\centering\arraybackslash}X}
\parindent=0pt
\fboxsep0pt

\begin{document}

{{ .Files }}

\end{document}`
