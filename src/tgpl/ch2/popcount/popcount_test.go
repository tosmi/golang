package popcount

import "testing"

func TestPopcount(t *testing.T) {
	if Popcount(1) != 1 {
		t.Error(`Popcount(1) = 1`)
	}
}

func BenchmarkPopcountWithoutLoop(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		Popcount(983745)
	}
}
