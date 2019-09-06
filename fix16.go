package fix16

import "math/bits"

var (
	Minimum = T{^0x7FFFFFFF}
	Maximum = T{0x7FFFFFFF}
)

var (
	Pi   = T{205887}
	E    = T{178145}
	Zero = T{}
	One  = T{0x00010000}
)

var Overflow = T{^0x7FFFFFFF}

type T struct {
	f int32
}

func Binary(a uint32) T {
	return T{int32(a)}
}

func Int(a int) T {
	return Int32(int32(a))
}

func Int32(a int32) T {
	return T{a * int32(One.f)}
}

func Uint32(a uint32) T {
	return Int32(int32(a))
}

func Float32(a float32) T {
	return Float64(float64(a))
}

func Float64(a float64) T {
	tmp := a * float64(One.f)
	if tmp >= 0 {
		return T{int32(tmp + 0.5)}
	}
	return T{int32(tmp - 0.5)}
}

func (a T) Float32() float32 {
	return float32(a.Float64())
}

func (a T) Float64() float64 {
	return float64(a.f) / float64(One.f)
}

func (a T) Binary() uint32 {
	return uint32(a.f)
}

func (a T) Int() int {
	return int(a.Int32())
}

func (a T) Int32() int32 {
	if a.f >= 0 {
		return int32((a.f + (One.f >> 1)) / One.f)
	}
	return int32((a.f - (One.f >> 1)) / One.f)
}

func (a T) Uint32() uint32 {
	return uint32(a.Int32())
}

func (a T) Split() (int32, T) {
	b := a.Binary()
	f := Binary(b & 0x0000FFFF)
	if a.Negative() {
		return -int32((b & 0x7FFFFFFF) >> 16), f
	}
	return int32(b >> 16), f
}

func (a T) Add(b T) T {
	ua := uint32(a.f)
	ub := uint32(b.f)
	sum := ua + ub

	// Overflow can only happen if sign of a == sign of b, and then
	// it causes sign of sum != sign of a.
	if ((ua^ub)&0x80000000) == 0 && ((ua^sum)&0x80000000) != 0 {
		return Overflow
	}
	return Binary(sum)
}

func (a T) AddSaturate(b T) T {
	r := a.Add(b)
	if r == Overflow {
		if a.Negative() {
			return Minimum
		} else {
			return Maximum
		}
	}
	return r
}

func (a T) Sub(b T) T {
	ua := uint32(a.f)
	ub := uint32(b.f)
	diff := ua - ub

	// Overflow can only happen if sign of a == sign of b, and then
	// it causes sign of sum != sign of a.
	if ((ua^ub)&0x80000000) != 0 && ((ua^diff)&0x80000000) != 0 {
		return Overflow
	}
	return Binary(diff)
}

func (a T) SubSaturate(b T) T {
	r := a.Sub(b)
	if r == Overflow {
		if a.Negative() {
			return Minimum
		} else {
			return Maximum
		}
	}
	return r
}

func (a T) Mul(b T) T {
	product := int64(a.f) * int64(b.f)

	// The upper 17 bits should all be the same (the sign).
	upper := uint32(product >> 47)

	if product < 0 {
		if ^upper != 0 {
			return Overflow
		}

		// This adjustment is required in order to round -1/2 correctly.
		product--
	} else if upper != 0 {
		return Overflow
	}

	result := int32(product >> 16)
	return T{result + int32((product&0x8000)>>15)}
}

func (a T) MulSaturate(b T) T {
	r := a.Mul(b)
	if r == Overflow {
		if a.Negative() == b.Negative() {
			return Minimum
		} else {
			return Maximum
		}
	}
	return r
}

func (a T) Div(b T) T {
	if b.f == 0 {
		return Minimum
	}

	remainder := uint32(-a.f)
	if a.f >= 0 {
		remainder = uint32(a.f)
	}

	divider := uint32(-b.f)
	if b.f >= 0 {
		divider = uint32(b.f)
	}

	quotient := uint32(0)
	bitPos := 17

	// Kick-start the division a bit.
	// This improves speed in the worst-case scenarios where N and D are large
	// It gets a lower estimate for the result by N/(D >> 17 + 1).
	if divider&0xFFF00000 != 0 {
		shiftedDiv := (divider >> 17) + 1
		quotient = remainder / shiftedDiv
		remainder -= uint32((uint64(quotient) * uint64(divider)) >> 17)
	}

	// If the divider is divisible by 2^n, take advantage of it.
	for divider&0xF == 0 && bitPos >= 4 {
		divider >>= 4
		bitPos -= 4
	}

	for remainder != 0 && bitPos >= 0 {
		// Shift remainder as much as we can without overflowing.
		shift := bits.LeadingZeros32(remainder)
		if shift > bitPos {
			shift = bitPos
		}
		remainder <<= uint32(shift)
		bitPos -= shift

		div := remainder / divider
		remainder = remainder % divider
		quotient += div << uint32(bitPos)

		if (div & (^(0xFFFFFFFF >> uint32(bitPos)))) != 0 {
			return Overflow
		}

		remainder <<= 1
		bitPos--
	}

	// Quotient is always positive so rounding is easy.
	quotient++

	result := quotient >> 1

	// Figure out the sign of the result.
	if (a.f^b.f)&(^0x7FFFFFFF) != 0 {
		if result == 0x80000000 {
			return Overflow
		}
		result = -result
	}
	return Binary(result)
}

func (a T) DivSaturate(b T) T {
	r := a.Div(b)
	if r == Overflow {
		if a.Negative() == b.Negative() {
			return Minimum
		} else {
			return Maximum
		}
	}
	return r
}

func (a T) Zero() bool {
	return a == Zero
}

func (a T) Negative() bool {
	return a.f < 0
}

func (a T) Less(b T) bool {
	return a.f < b.f
}

func (a T) LEqual(b T) bool {
	return a.f <= b.f
}

func (a T) Inv() T {
	return T{-a.f}
}

func (a T) Mod(b T) T {
	return T{a.f % b.f}
}

func (a T) Abs() T {
	if a.f < 0 {
		return a.Inv()
	}
	return a
}

func (a T) Floor() T {
	return T{a.f & (^0x0000FFFF)}
}

func (a T) Ceil() T {
	var n int32
	if a.f&0x0000FFFF != 0 {
		n = One.f
	}
	return T{(a.f & (^0x0000FFFF)) + n}
}

func (a T) Min(b T) T {
	if a.f < b.f {
		return a
	}
	return b
}

func (a T) Max(b T) T {
	if a.f > b.f {
		return a
	}
	return b
}

func (a T) Clamp(low, high T) T {
	return a.Max(low).Min(high)
}
