package types

import "fmt"

type Spring float64

func (s Spring) String() string {
	return fmt.Sprintf("%.4g", s)
}
