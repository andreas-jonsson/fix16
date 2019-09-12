package vec2

import (
	"fmt"
	"image"

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

func floorSqrt(x int32) int32 {
	if x == 0 || x == 1 {
		return x
	}

	var (
		start int32 = 1
		end         = x
		res   int32
	)

	for start <= end {
		mid := (start + end) / 2

		// If x is a perfect square.
		sprt := mid * mid
		if sprt == x {
			return mid
		}

		// Since we need floor, we update answer when mid*mid is
		// smaller than x, and move closer to sqrt(x).
		if sprt < x {
			start = mid + 1
			res = mid
		} else { // If mid*mid is greater than x.
			end = mid - 1
		}
	}
	return res
}

func (v T) Length() fix16.T {
	x, xf := v.X().Split()
	y, yf := v.Y().Split()

	xi, yi := x.Int32(), y.Int32()

	il := floorSqrt(yi*yi + xi*xi)
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
