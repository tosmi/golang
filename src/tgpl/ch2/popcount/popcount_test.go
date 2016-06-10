package popcount

import (
	"testing"
	//"fmt"
)

func TestPopcount(t *testing.T) {
	if Popcount(1) != 1 {
		t.Error(`Popcount(1) = 1`)
	}
}

func BenchmarkPopcountShift(b *testing.B) {
	//fmt.Printf("Loop %v\n", b.N)
	for i := 0; i <= b.N; i++ {
		PopcountShift(983745)
	}
}

func BenchmarkPopcountLoop(b *testing.B) {
	//fmt.Printf("Loop %v\n", b.N)
	for i := 0; i <= b.N; i++ {
		PopcountLoop(983745)
	}
}

func BenchmarkPopcountWithoutLoop(b *testing.B) {
	//fmt.Printf("Without Loop %v\n", b.N)
	for i := 0; i <= b.N; i++ {
		Popcount(983745)
	}
}
