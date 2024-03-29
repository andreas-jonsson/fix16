package fix16

func (a T) Sqrt() T {
	var result uint32
	negative := a.Negative()

	num := uint32(a.f)
	if negative {
		num = uint32(-a.f)
	}

	var bit uint32 = 1 << 30
	if num&0xFFF00000 == 0 {
		bit = 1 << 18
	}

	for bit > num {
		bit >>= 2
	}

	// The main part is executed twice, in order to avoid
	// using 64 bit values in computations.
	for n := 0; n < 2; n++ {
		// First we get the top 24 bits of the answer.
		for bit != 0 {
			if num >= result+bit {
				num -= result + bit
				result = (result >> 1) + bit
			} else {
				result = result >> 1
			}
			bit >>= 2
		}

		if n == 0 {
			// Then process it again to get the lowest 8 bits.
			if num > 65535 {
				// The remainder 'num' is too large to be shifted left
				// by 16, so we have to add 1 to result manually and
				// adjust 'num' accordingly.
				// num = a - (result + 0.5)^2
				//	 = num + result^2 - (result + 0.5)^2
				//	 = num - result - 0.5
				num -= result
				num = (num << 16) - 0x8000
				result = (result << 16) + 0x8000
			} else {
				num <<= 16
				result <<= 16
			}

			bit = 1 << 14
		}
	}

	// Finally, if next bit would have been 1, round the result upwards.
	if num > result {
		result++
	}

	if negative {
		return Binary(result).Inv()
	}
	return Binary(result)
}
