package stringutil

import (
	"testing"
)

func TestSubstr(t *testing.T) {
	testCases := []struct {
		str    string
		start  int
		length int
		want   string
	}{
		{"hello", 0, 5, "hello"},
		{"hello", 1, 4, "ello"},
		{"hello", 2, 3, "llo"},
		{"hello", -1, 1, "o"},
		{"hello", -2, 2, "lo"},
		{"hello", -5, 5, "hello"},
		{"hello", 0, 0, ""},
		{"hello", 1, 0, ""},
		{"hello", 1, 1, "e"},
		{"", 0, 1, ""},
		{"", -1, 1, ""},
		{"a", 0, 1, "a"},
		{"a", -1, 1, "a"},
		{"a", 1, 1, ""},
		{"a最后的战神", 0, 2, "a最"},
	}

	for _, tc := range testCases {
		got := Substr(tc.str, tc.start, tc.length)
		if got != tc.want {
			t.Errorf("Substr(%q, %d, %d) = %q; want %q", tc.str, tc.start, tc.length, got, tc.want)
		}
	}
}
