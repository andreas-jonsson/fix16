package fix16

import "math/bits"

const (
	Minimum T = ^0x7FFFFFFF
	Maximum T = 0x7FFFFFFF
)

const (
	Pi  T = 205887
	E   T = 178145
	One T = 0x00010000
)

const Overflow T = ^0x7FFFFFFF

type T int32

func Int(a int) T {
	return Int32(int32(a))
}

func Int32(a int32) T {
	return T(a * int32(One))
}

func Uint32(a uint32) T {
	return Int32(int32(a))
}

func Float32(a float32) T {
	return Float64(float64(a))
}

func Float64(a float64) T {
	tmp := a * float64(One)
	if tmp >= 0 {
		return T(tmp + 0.5)
	}
	return T(tmp - 0.5)
}

func (a T) Float32() float32 {
	return float32(a.Float64())
}

func (a T) Float64() float64 {
	return float64(a) / float64(One)
}

func (a T) Int() int {
	return int(a.Int32())
}

func (a T) Int32() int32 {
	if a >= 0 {
		return int32((a + (One >> 1)) / One)
	}
	return int32((a - (One >> 1)) / One)
}

func (a T) Uint32() uint32 {
	return uint32(a.Int32())
}

func (a T) Add(b T) T {
	ua := uint32(a)
	ub := uint32(b)
	sum := ua + ub

	// Overflow can only happen if sign of a == sign of b, and then
	// it causes sign of sum != sign of a.
	if ((ua^ub)&0x80000000) == 0 && ((ua^sum)&0x80000000) != 0 {
		return Overflow
	}
	return T(sum)
}

func (a T) AddSaturate(b T) T {
	r := a.Add(b)
	if r == Overflow {
		if a >= 0 {
			return Maximum
		} else {
			return Minimum
		}
	}
	return r
}

func (a T) Sub(b T) T {
	ua := uint32(a)
	ub := uint32(b)
	diff := ua - ub

	// Overflow can only happen if sign of a == sign of b, and then
	// it causes sign of sum != sign of a.
	if ((ua^ub)&0x80000000) != 0 && ((ua^diff)&0x80000000) != 0 {
		return Overflow
	}
	return T(diff)
}

func (a T) SubSaturate(b T) T {
	r := a.Sub(b)
	if r == Overflow {
		if a >= 0 {
			return Maximum
		} else {
			return Minimum
		}
	}
	return r
}

func (a T) Mul(b T) T {
	product := int64(a) * int64(b)

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

	result := T(product >> 16)
	return result + T((product&0x8000)>>15)
}

func (a T) MulSaturate(b T) T {
	r := a.Mul(b)
	if r == Overflow {
		if (a >= 0) == (b >= 0) {
			return Maximum
		} else {
			return Minimum
		}
	}
	return r
}

func (a T) Div(b T) T {
	if b == 0 {
		return Minimum
	}

	remainder := uint32(-a)
	if a >= 0 {
		remainder = uint32(a)
	}

	divider := uint32(-b)
	if b >= 0 {
		divider = uint32(b)
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
	if (a^b)&(^0x7FFFFFFF) != 0 {
		if result == 0x80000000 {
			return Overflow
		}
		result = -result
	}
	return T(result)
}

func (a T) DivSaturate(b T) T {
	r := a.Div(b)
	if r == Overflow {
		if (a >= 0) == (b >= 0) {
			return Maximum
		} else {
			return Minimum
		}
	}
	return r
}

func (a T) Mod(b T) T {
	return a % b
}

func (a T) Abs() T {
	if a < 0 {
		return -a
	}
	return a
}

func (a T) Floor() T {
	return a & (^0x0000FFFF)
}

func (a T) Ceil() T {
	var n T
	if a&0x0000FFFF != 0 {
		n = One
	}
	return (a & (^0x0000FFFF)) + n
}

func (a T) Min(b T) T {
	if a < b {
		return a
	}
	return b
}

func (a T) Max(b T) T {
	if a > b {
		return a
	}
	return b
}

func (a T) Clamp(low, high T) T {
	return a.Max(low).Min(high)
}
