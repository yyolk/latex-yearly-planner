package types

import "fmt"

type Millimeters float64

func (r Millimeters) String() string {
	return fmt.Sprintf("%.4fmm", r)
}

type Spring float64

func (r Spring) String() string {
	return fmt.Sprintf("%.4f", r)
}
