package vec2

import (
	"testing"

	"github.com/andreas-jonsson/fix16"
)

func TestLength(t *testing.T) {
	v := T{fix16.Float64(2.2), fix16.Float64(1.5)}

	if v.Length().Format(4) != "2.5385" {
		t.Fail()
	}
}

func TestNormalize(t *testing.T) {
	v := T{fix16.Float64(2), fix16.Float64(1.5)}
	v = v.Normalize()

	if v.Length().Int() != 1 {
		t.Fail()
	}
}
