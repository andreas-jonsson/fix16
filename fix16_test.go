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

	if a != -b {
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
}

func TestString(t *testing.T) {
	a := Float64(12.25)

	if a.String() != "12.25000" {
		t.Fail()
	}

	a = Float64(-10.5)

	if a.String() != "-10.50000" {
		t.Fail()
	}
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
	if T(214748302).Sqrt() != T(3751499) {
		t.Fail()
	}

	if T(214748303).Sqrt() != T(3751499) {
		t.Fail()
	}

	if T(214748359).Sqrt() != T(3751499) {
		t.Fail()
	}

	if T(214748360).Sqrt() != T(3751500) {
		t.Fail()
	}
}
