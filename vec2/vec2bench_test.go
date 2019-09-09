package vec2

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/andreas-jonsson/fix16"
)

func BenchmarkFloorSqrt(b *testing.B) {
	for i := 0; i < 10; i++ {
		r := rand.Int31()
		b.Run(fmt.Sprintf("floorSqrt(%d)", r), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				floorSqrt(r)
			}
		})
	}
}

func BenchmarkLength(b *testing.B) {
	v := T{fix16.Int32(rand.Int31()), fix16.Int32(rand.Int31())}
	for n := 0; n < b.N; n++ {
		v.Length()
	}
}
