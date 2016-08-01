package strcat

import (
	"strings"
	"testing"
	"testing/quick"
)

func TestJoinStrings(t *testing.T) {
	f := func(s []string) bool {
		out := joinStrings(s)
		return strings.Join(s, "") == out
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func BenchmarkJoinStrings(b *testing.B) {
	input := []string{"1", "2", "3", "4", "5"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		joinStrings(input)
	}
}
