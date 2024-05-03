package lenconv

import (
	"fmt"
	"os"
)

type Meter float64
type Ft float64

func (m Meter) String(s string) {
	fmt.Fprintf(os.Stdout, "%s Meter", s)
}

func (m Ft) String(s string) {
	fmt.Fprintf(os.Stdout, "%s Ft.", s)
}
