package texsnippets

const Document = "document"
const document = `\documentclass[9pt]{extarticle}

\usepackage{geometry}
\usepackage[table]{xcolor}
\usepackage{showframe}
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
\usepackage{easytable}


\geometry{paperwidth={{.Device.Paper.Width}}, paperheight={{.Device.Paper.Height}}}
\geometry{
             top={{ .Layout.Margin.Top }},
          bottom={{ .Layout.Margin.Bottom }},
            left={{ .Layout.Margin.Left }},
           right={{ .Layout.Margin.Right }},
  marginparwidth={{ .Layout.MarginNotes.Width }},
    marginparsep={{ .Layout.MarginNotes.Margin }}
}
{{ .Layout.MarginNotes.Reverse }}

\pagestyle{empty}
\newcolumntype{Y}{>{\centering\arraybackslash}X}
\parindent=0pt
\fboxsep0pt

\begin{document}

{{ .Files }}

\end{document}`

const Title = "title"
const title = `\hspace{0pt}\vfil
\hfill\resizebox{.7\linewidth}{!}{ {{- .Year -}} }%
\pagebreak`

const MOSHeader = "mos-header"
const mosHeader = `\marginnote{\rotatebox[origin=tr]{90}{%
{{ .MarginNotes }}%
}}%
{{ .Header }}

`
