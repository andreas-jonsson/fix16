package fix16

import (
	"testing"
)

func TestSimpleAdd(t *testing.T) {
	a := Float64(2.75)
	b := Float64(3.5)

	if a.Add(b).Float64() != 6.25 {
		t.Fail()
	}
}

func TestSimpleSub(t *testing.T) {
	a := Float64(2.75)
	b := Float64(1.5)

	if a.Sub(b).Float64() != 1.25 {
		t.Fail()
	}

	a = Parse("-1.25")
	b = Parse("1.25")

	if a != b.Inv() {
		t.Fail()
	}
}

func TestSimpleMul(t *testing.T) {
	a := Float64(7.75)
	b := Float64(2)

	if a.Mul(b).Float64() != 15.5 {
		t.Fail()
	}
}

func TestSimpleDiv(t *testing.T) {
	a := Float64(1.5)
	b := Float64(-0.5)

	if a.Div(b).Float64() != -3 {
		t.Fail()
	}
}

func TestParse(t *testing.T) {
	a := Parse("12.25")

	if a.Float64() != 12.25 {
		t.Fail()
	}

	a = Parse("-10.5")

	if a.Float64() != -10.5 {
		t.Fail()
	}
	/*
		a = Parse("0.00002")

		if a != One {
			t.Fail()
		}

		a = Parse("0.99998")

		if a != Binary(65535) {
			t.Fail()
		}

		a = Parse("32767.99998")

		if a != Maximum {
			t.Fail()
		}

		a = Parse("-32768.00000")

		if a != Minimum {
			t.Fail()
		}
	*/
}

func TestSimpleString(t *testing.T) {
	a := Float64(12.25)

	if a.String() != "12.25000" {
		t.Fail()
	}

	a = Float64(-10.5)

	if a.String() != "-10.50000" {
		t.Fail()
	}
}

func TestString(t *testing.T) {
	/*
		value := Minimum
		for value.Less(Maximum) {
			fvalue := math.Round(value.Float64()*100000) / 100000
			fstr := fmt.Sprintf("%0.5f", fvalue)

			str := value.String()
			if str != fstr {
				t.FailNow()
			}

			pvalue := Parse(str)
			if pvalue != value {
				t.FailNow()
			}

			value = Binary(value.Binary() + 0x10001)
		}
	*/
}

func TestSimpleSqrt(t *testing.T) {
	if Int(16).Sqrt() != Int(4) {
		t.Fail()
	}

	if Int(100).Sqrt() != Int(10) {
		t.Fail()
	}

	if Int(1).Sqrt() != Int(1) {
		t.Fail()
	}
}

func TestSqrt(t *testing.T) {
	if Binary(214748302).Sqrt() != Binary(3751499) {
		t.Fail()
	}

	if Binary(214748303).Sqrt() != Binary(3751499) {
		t.Fail()
	}

	if Binary(214748359).Sqrt() != Binary(3751499) {
		t.Fail()
	}

	if Binary(214748360).Sqrt() != Binary(3751500) {
		t.Fail()
	}
}
