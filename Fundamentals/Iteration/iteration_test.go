package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 10)
	expected := "aaaaaaaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func TestIndex(t *testing.T) {
	indexed := strings.Index("abcde", "de")
	expected := 3

	if indexed != expected {
		t.Errorf("expected %q but got %q", expected, indexed)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}

func ExampleRepeat() {
	string := Repeat("a", 10)
	fmt.Println(string)
	// Output: aaaaaaaaaa
}
