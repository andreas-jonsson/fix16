package vec2

import (
	"fmt"
	"image"

	"github.com/andreas-jonsson/fix16"
)

type T [2]fix16.T

func (v T) String() string {
	return fmt.Sprintf("[%s, %s]", v.X().String(), v.Y().String())
}

func (v T) X() fix16.T {
	return v[0]
}

func (v T) Y() fix16.T {
	return v[1]
}

func (v T) XY() (fix16.T, fix16.T) {
	return v.X(), v.Y()
}

func (v T) YX() (fix16.T, fix16.T) {
	return v.Y(), v.X()
}

func (v T) Add(b T) T {
	return T{v.X().Add(b.X()), v.Y().Add(b.Y())}
}

func (v T) Sub(b T) T {
	return T{v.X().Sub(b.X()), v.Y().Sub(b.Y())}
}

func (v T) Mul(b T) T {
	return T{v.X().Mul(b.X()), v.Y().Mul(b.Y())}
}

func (v T) Div(b T) T {
	return T{v.X().Div(b.X()), v.Y().Div(b.Y())}
}

func (v T) Length() fix16.T {
	return v.LengthSqr().Sqrt()
}

func (v T) LengthSqr() fix16.T {
	return v.X().Mul(v.X()).Add(v.Y().Mul(v.Y()))
}

func (v T) Scale(s fix16.T) T {
	return T{v.X().Mul(s), v.Y().Mul(s)}
}

func (v T) Invert() T {
	return T{-v.X(), -v.Y()}
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
	return v.X().Mul(b.X()).Add(v.Y().Mul(b.Y()))
}

func (v T) Point() image.Point {
	return image.Pt(v.X().Int(), v.Y().Int())
}

func (v T) Rectangle() image.Rectangle {
	return image.Rect(0, 0, v.X().Int(), v.Y().Int())
}
