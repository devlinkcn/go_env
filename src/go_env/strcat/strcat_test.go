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
