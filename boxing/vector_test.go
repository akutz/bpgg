package boxing_test

import (
	"math/rand"
	"testing"

	vector "bpgg/boxing"
)

func BenchmarkInterfaceVector(b *testing.B) {
	var v vector.IVector
	for i := 0; i < b.N; i++ {
		v = append(v, rand.Int())
	}
}

func BenchmarkGenericVector(b *testing.B) {
	var v vector.GVector[int]
	for i := 0; i < b.N; i++ {
		v = append(v, rand.Int())
	}
}
