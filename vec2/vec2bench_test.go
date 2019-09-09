package vec2

import (
	"fmt"
	"math/rand"
	"testing"
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
