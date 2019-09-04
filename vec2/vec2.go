package vec2

import (
	"fmt"

	"github.com/andreas-jonsson/fix16"
)

type T [2]fix16.T

func (v T) String() string {
	return fmt.Sprintf("[%s, %s]", v[0].String(), v[1].String())
}

func (v T) X() fix16.T {
	return v[0]
}

func (v T) Y() fix16.T {
	return v[1]
}

func (v T) Add(b T) T {
	return T{v[0].Add(b[0]), v[1].Add(b[1])}
}

func (v T) Sub(b T) T {
	return T{v[0].Sub(b[0]), v[1].Sub(b[1])}
}

func (v T) Mul(b T) T {
	return T{v[0].Mul(b[0]), v[1].Mul(b[1])}
}

func (v T) Div(b T) T {
	return T{v[0].Div(b[0]), v[1].Div(b[1])}
}

func (v T) Length() fix16.T {
	return v.LengthSqr().Sqrt()
}

func (v T) LengthSqr() fix16.T {
	return v[0].Mul(v[0]).Add(v[1].Mul(v[1]))
}

func (v T) Scale(s fix16.T) T {
	return T{v[0].Mul(s), v[1].Mul(s)}
}

func (v T) Invert() T {
	return T{-v[0], -v[1]}
}

func (v T) Normalize() T {
	sl := v.LengthSqr()
	if sl == 0 || sl == 1 {
		return v
	}
	s := fix16.One.Div(sl.Sqrt())
	return v.Scale(s)
}

func (v T) Dot(b T) fix16.T {
	return v[0].Mul(b[0]).Add(v[1].Mul(b[1]))
}
