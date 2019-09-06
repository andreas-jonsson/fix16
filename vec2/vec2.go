package vec2

import (
	"fmt"
	"image"
	"math"

	"github.com/andreas-jonsson/fix16"
)

func Rectangle(r image.Rectangle) T {
	if r.Min != image.ZP {
		panic("rectangle min is not zero")
	}
	return Point(r.Max)
}

func Point(pt image.Point) T {
	return Int(pt.X, pt.Y)
}

func Int(x, y int) T {
	return T{fix16.Int(x), fix16.Int(y)}
}

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

/*
func (v T) Length() fix16.T {
	sl := v.X().Mul(v.X()).Add(v.Y().Mul(v.Y()))
	return sl.Sqrt()
}
*/

func (v T) Length() fix16.T {
	xi, xf := v.X().Split()
	yi, yf := v.Y().Split()

	// TODO: Remove math package.
	il := int32(math.Sqrt(float64(int64(yi)*int64(yi) + int64(xi)*int64(xi))))
	fl := xf.Mul(xf).Add(yf.Mul(yf)).Sqrt()

	return fix16.Int32(il).Add(fl)
}

func (v T) Scale(s fix16.T) T {
	return T{v.X().Mul(s), v.Y().Mul(s)}
}

func (v T) Invert() T {
	return T{v.X().Inv(), v.Y().Inv()}
}

func (v T) Normalize() T {
	l := v.Length()
	if l == fix16.Zero || l == fix16.Binary(1) {
		return v
	}
	s := fix16.One.Div(l)
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
