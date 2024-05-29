package arays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Any size array", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6
		if got != want {
			t.Errorf("got %d, want %d, given %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("Two array", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 3})
		want := []int{3, 3}
		checkSums(t, got, want)
	})

	t.Run("Three array", func(t *testing.T) {
		got := SumAll([]int{1}, []int{-1, 2}, []int{1, 2, 3, 4})
		want := []int{1, 1, 10}
		checkSums(t, got, want)
	})
}

func TestSumAllTails(t *testing.T) {
	//checkSums := func(t testing.TB, got, want []int) {
	//	t.Helper()
	//	if !slices.Equal(got, want) {
	//		t.Errorf("got %v, want %v", got, want)
	//	}
	//}
	t.Run("Three arrays", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3, 4}, []int{1}, []int{1, 2})
		want := []int{9, 0, 2}
		checkSums(t, got, want)
	})

	t.Run("Empty array", func(t *testing.T) {
		got := SumAllTails([]int{})
		want := []int{0}
		checkSums(t, got, want)
	})
}

func checkSums(t testing.TB, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
