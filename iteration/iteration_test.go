package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat("a", 6)
	want := "aaaaaa"
	if got != want {
		t.Errorf("Got %q, want %q", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 1)
	}
}

// Repeat function takes a string(s), a number(n) and outputs s repeated n times
func ExampleRepeat() {
	fmt.Println(Repeat("a", 6))
	// Output: aaaaaa
}
