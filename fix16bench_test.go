package fix16

import (
	"math"
	"testing"
)

func BenchmarkFloat32Add(b *testing.B) {
	x, y := float32(1), float32(1)
	for n := 0; n < b.N; n++ {
		x = x + y
	}
}

func BenchmarkFix16Add(b *testing.B) {
	x, y := One, One
	for n := 0; n < b.N; n++ {
		x = x.Add(y)
	}
}

func BenchmarkFloat32Mul(b *testing.B) {
	x, y := float32(1), float32(2)
	for n := 0; n < b.N; n++ {
		x = x * y
	}
}

func BenchmarkFix16Mul(b *testing.B) {
	x, y := One, Int(2)
	for n := 0; n < b.N; n++ {
		x = x.Mul(y)
	}
}

func BenchmarkFloat32Div(b *testing.B) {
	x, y := float32(1), float32(2)
	for n := 0; n < b.N; n++ {
		x = x / y
	}
}

func BenchmarkFix16Div(b *testing.B) {
	x, y := One, Int(2)
	for n := 0; n < b.N; n++ {
		x = x.Div(y)
	}
}

func BenchmarkFloat64Sqrt(b *testing.B) {
	i := Maximum.Float64()
	for n := 0; n < b.N; n++ {
		math.Sqrt(i)
	}
}

func BenchmarkFix16Sqrt(b *testing.B) {
	i := Maximum
	for n := 0; n < b.N; n++ {
		i.Sqrt()
	}
}
