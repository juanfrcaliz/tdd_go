package arraysslices

import (
	"fmt"
	"testing"
)

const (
	Result int = 15
)

var (
	Slice = []int{1, 2, 3, 4, 5}
)

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		got := Sum(Slice)
		want := Result

		if got != want {
			t.Errorf("wanted %d but got %d, given %v", want, got, Slice)
		}
	})
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(Slice)
	}
}

func ExampleSum() {
	sum := Sum(Slice)
	fmt.Println(sum)

	// Output: 15
}
