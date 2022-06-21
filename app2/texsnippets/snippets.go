package texsnippets

const Document = "document"
const document = `\documentclass[9pt]{extarticle}

\usepackage[showframe]{geometry}
\usepackage[table]{xcolor}
\usepackage{tabularx}
\usepackage{hyperref}
\usepackage{marginnote}
\usepackage{adjustbox}
\usepackage{multido}
\usepackage{amssymb}

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

\newlength{\myLengthTwoColumnsSeparatorWidth}
\setlength{\myLengthTwoColumnsSeparatorWidth}{ {{- .Layout.Sizes.TwoColumnsSeparatorSize -}} }

\newlength{\myLengthTwoColumnWidth}
\setlength{\myLengthTwoColumnWidth}{\dimexpr.5\textwidth-.5\myLengthTwoColumnsSeparatorWidth}

\newlength{\myLengthThreeColumnsSeparatorWidth}
\setlength{\myLengthThreeColumnsSeparatorWidth}{ {{- .Layout.Sizes.ThreeColumnsSeparatorSize -}} }

\newlength{\myLengthThreeColumnWidth}
\setlength{\myLengthThreeColumnWidth}{\dimexpr.333\textwidth-.667\myLengthThreeColumnsSeparatorWidth}

\newlength{\myLengthTwoThirdsColumnWidth}
\setlength{\myLengthTwoThirdsColumnWidth}{\dimexpr2\myLengthThreeColumnWidth+\myLengthThreeColumnsSeparatorWidth}

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

const Title = "title"
const title = `\hspace{0pt}\vfil
\hfill\resizebox{.7\linewidth}{!}{ {{- .Title -}} }%
\pagebreak`
