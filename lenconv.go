package lenconv

import (
	"fmt"
)

type Meter float64
type Ft float64

func (m Meter) String() string {
	return fmt.Sprintf("%g Meter", m)
}

func (f Ft) String() string {
	return fmt.Sprintf("%g Ft.", f)
}
