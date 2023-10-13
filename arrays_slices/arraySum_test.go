package arraysslices

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		got := Sum([]int{1, 2, 3, 4, 5})
		want := 15

		if got != want {
			t.Errorf("wanted %d but got %d, given %v", want, got, []int{1, 2, 3, 4, 5})
		}
	})
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum([]int{1, 2, 3, 4, 5})
	}
}

func ExampleSum() {
	sum := Sum([]int{1, 2, 3, 4, 5})
	fmt.Println(sum)

	// Output: 15
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2, 3, 4, 5}, []int{3, 4, 5})
	want := []int{15, 12}

	assertEqualSlice(t, got, want)
}

func TestSumAllTails(t *testing.T) {
	t.Run("non-empty slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3, 4, 5}, []int{3, 4, 5})
		want := []int{14, 9}

		assertEqualSlice(t, got, want)
	})
	t.Run("one empty slice", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{1, 2, 3, 4, 5})
		want := []int{0, 14}

		assertEqualSlice(t, got, want)
	})

}

func assertEqualSlice(t testing.TB, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Wanted %v but got %v", want, got)
	}
}
