package types

import "fmt"

type Millimeters float64

func (r Millimeters) String() string {
	return fmt.Sprintf("%.4fmm", r)
}
