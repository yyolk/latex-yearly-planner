package tex

import (
	"fmt"

	"github.com/kudrykv/latex-yearly-planner/app2/planners/common"
)

const nl = "\n"

func CellColor(color, text string) string {
	return fmt.Sprintf(`\cellcolor{%s}{%s}`, color, text)
}

func TextColor(color, text string) string {
	return fmt.Sprintf(`\textcolor{%s}{%s}`, color, text)
}

func Hyperlink(ref, text string) string {
	return fmt.Sprintf(`\hyperlink{%s}{%s}`, ref, text)
}

func Hypertarget(ref, text string) string {
	return fmt.Sprintf(`\hypertarget{%s}{%s}`, ref, text)
}

func Tabular(format, text string) string {
	return `\begin{tabular}{` + format + `}` + nl + text + nl + `\end{tabular}`
}

func TabularXAlignTopLineWidth(format, text string) string {
	return TabularXLineWidth(`t`, format, text)
}

func TabularXLineWidth(align, format, text string) string {
	return TabularX(`\linewidth`, align, format, text)
}

func TabularX(width, align, format, text string) string {
	return `\begin{tabularx}{` + width + `}[` + align + `]{` + format + `}` + nl + text + nl + `\end{tabularx}`
}

func ResizeBoxW(width, text string) string {
	return fmt.Sprintf(`\resizebox{!}{%s}{%s}`, width, text)
}

func Multirow(rows int, text string) string {
	return fmt.Sprintf(`\multirow{%d}{*}{%s}`, rows, text)
}

func Bold(text string) string {
	return fmt.Sprintf(`\textbf{%s}`, text)
}

func AdjustBox(text string) string {
	return `\adjustbox{valign=t}{` + text + `}`
}

func RenewArrayStretch(value string) string {
	return RenewCommand(`\arraystretch`, value)
}

func RenewCommand(command, value string) string {
	return fmt.Sprintf(`\renewcommand{%s}{%s}`, command, value)
}

func LineHeight(value common.Millimeters) string {
	return Parbox(`0pt`, `\vskip`+value.String())
}

func Parbox(width, text string) string {
	return fmt.Sprintf(`\parbox{%s}{%s}`, width, text)
}
