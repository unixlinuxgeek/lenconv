package lenconv

import (
	"github.com/unixlinuxgeek/floatgen"
	"testing"
)

// делим значение длины на 3.281
func TestFtToMet(t *testing.T) {
	r := floatgen.GenRan(1, 9)
	o := r / 0.3048
	b := float64(FtToMet(Ft(r)))

	t.Log(o, b)
	if o == b {
		t.Logf("Test Passed!!! %v equal %v\n", o, b)
	} else {
		t.Fatalf("Test Failed %v not equal %v\n", o, b)
	}
}

// умножаем значение длины на 3.281
func TestMetToFt(t *testing.T) {
	r := floatgen.GenRan(1, 9)
	o := r * 3.281
	b := float64(MetToFt(Meter(r)))

	if o == b {
		t.Logf("%v equal %v\n", o, b)
	} else {
		t.Fatalf("%v not equal %v\n", o, b)
	}
}
