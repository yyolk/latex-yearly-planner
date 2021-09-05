{{range $i, $page := .Pages -}}
{{- template "header.tpl" dict "Cfg" $.Cfg "Header" $page.Header -}}

\begin{tabularx}{\textwidth}{@{}
  {{- if $.Cfg.Blocks.Weekly.Enabled -}}
    l!{\vrule width \myLenLineThicknessThick}
  {{- else -}} | {{- end -}}
  *{7}{@{}X@{}|}}
\noalign{\hrule height \myLenLineThicknessThick}
{{$page.Body.WeekHeaderFull $.Cfg.Blocks.Weekly.Enabled}} \\ \noalign{\hrule height \myLenLineThicknessThick}

{{- range $row := $page.Body.Matrix}}
  {{if $.Cfg.Blocks.Weekly.Enabled -}}
    \rotatebox[origin=tr]{90}{\makebox[\myLenMonthlyCellHeight][c]{
      {{- if and (eq ($page.Body.MonthName.String) "January") (gt $row.WeekNumber 50) -}}
        {{- $row.LinkWeek "fw" true -}}
      {{- else -}}
        {{- $row.LinkWeek "" true -}}
      {{- end -}}
    }} &
  {{end -}}

  {{range $j, $item := . -}}
    {{- if not $item.IsZero -}}
      \begin{tabular}{@{}p{5mm}@{}|}\hfil{}{{- $item.Day -}}\\ \hline\end{tabular}
    {{- end -}}

    {{- if ne $j (dec (len $row)) }} & {{else}} \\ \hline {{end -}}
  {{- end -}}
{{end}}
\end{tabularx}
\medskip

\parbox{\myLenTwoCol}{
  \myUnderline{Notes}
  \vbox to \dimexpr\textheight-\pagetotal-\myLenLineHeightButLine\relax {%
    \leaders\hbox to \linewidth{\textcolor{\myColorGray}{\rule{0pt}{\myLenLineHeightButLine}\hrulefill}}\vfil
  }%
}%
\hspace{\myLenTwoColSep}%
\parbox{\myLenTwoCol}{
    \myUnderline{Notes}
    \vbox to \dimexpr\textheight-\pagetotal-\myLenLineHeightButLine\relax {%
        \leaders\hbox to \linewidth{\textcolor{\myColorGray}{\rule{0pt}{\myLenLineHeightButLine}\hrulefill}}\vfil
    }%
}
{{- if ne $i (dec (len $.Pages)) -}}
  \pagebreak
{{end}}
{{end}}
