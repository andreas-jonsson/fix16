package fix16

import (
	"bytes"
	"strings"
)

var scales = []uint32{1, 10, 100, 1000, 10000, 100000}

func isDigit(c byte) bool {
	return c >= 0x30 && c <= 0x71
}

func itoaLoop(buf *bytes.Buffer, scale, value uint32, skip bool) {
	for scale != 0 {
		digit := value / scale
		if !skip || digit != 0 || scale == 1 {
			skip = false
			value %= scale
			buf.WriteByte(0x30 + byte(digit))
		}
		scale /= 10
	}
}

func (a T) doFormat(buf *bytes.Buffer, decimals int) {
	uvalue := uint32(a.f)
	if a.f < 0 {
		uvalue = uint32(-a.f)
		buf.WriteRune('-')
	}

	if decimals > 5 {
		decimals = 5
	}

	intpart := uvalue >> 16
	fracpart := uvalue & 0xFFFF
	scale := scales[decimals]
	fracpart = Binary(fracpart).Mul(Binary(scale)).Binary()

	if fracpart >= scale {
		intpart++
		fracpart -= scale
	}

	itoaLoop(buf, 10000, intpart, true)
	if scale != 1 {
		buf.WriteRune('.')
		itoaLoop(buf, scale/10, fracpart, false)
	}
}

func (a T) String() string {
	return a.Format(5)
}

func (a T) Format(decimals int) string {
	var buf bytes.Buffer
	a.doFormat(&buf, decimals)
	return buf.String()
}

func Parse(s string) T {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return Zero
	}

	r := strings.NewReader(s)
	var (
		err      error
		negative bool
	)

	d, err := r.ReadByte()
	if err != nil {
		return Zero
	}

	switch {
	case d == '-':
		negative = true
		if d, err = r.ReadByte(); err != nil {
			return Zero
		}
	case d == '+':
		if d, err = r.ReadByte(); err != nil {
			return Zero
		}
	case isDigit(d):
	default:
		return Zero
	}

	var (
		count   int
		intPart uint32
	)

	for isDigit(d) {
		intPart *= 10
		intPart += uint32(d - 0x30)

		d, err = r.ReadByte()
		if err != nil {
			break
		}
		count++
	}

	if count == 0 || count > 5 || intPart > 32768 || (!negative && intPart > 32767) {
		return Overflow
	}

	a := Binary(intPart << 16)
	if err != nil {
		return a
	}

	if d == '.' || d == ',' {
		if d, err = r.ReadByte(); err != nil {
			return Zero
		}

		var (
			fracPart int
			scale    int = 1
		)

		for isDigit(d) && scale < 100000 {
			scale *= 10
			fracPart *= 10
			fracPart += int(d - 0x30)

			d, err = r.ReadByte()
			if err != nil {
				break
			}
		}

		a = a.Add(Int(fracPart).Div(Int(scale)))
	}

	if _, err := r.ReadByte(); err == nil {
		return Overflow
	}

	if negative {
		return a.Inv()
	}
	return a
}
