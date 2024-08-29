package stringutil

import (
	"fmt"
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

// 测试 Md5 函数
func TestMd5(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello", "5d41402abc4b2a76b9719d911017c592"},
		// 添加更多测试用例
	}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			actual := Md5(tc.input)
			fmt.Println(actual)
			if actual != tc.expected {
				t.Errorf("Md5(%q) = %q; expected %q", tc.input, actual, tc.expected)
			}
		})
	}
}

// TestConcatStrings 测试 ConcatStrings 函数
func TestConcatStrings(t *testing.T) {
	tests := []struct {
		name     string
		elems    []string
		expected string
	}{
		{
			name:     "single element",
			elems:    []string{"hello"},
			expected: "hello",
		},
		{
			name:     "multiple elements",
			elems:    []string{"hello", " ", "world"},
			expected: "hello world",
		},
		{
			name:     "empty elements",
			elems:    []string{"", ""},
			expected: "",
		},
		{
			name:     "中文拼接",
			elems:    []string{"你好", " ", "桂林"},
			expected: "你好 桂林",
		},
		{
			name:     "nil elements",
			elems:    nil,
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := ConcatStrings(tt.elems...)
			if actual != tt.expected {
				t.Errorf("ConcatStrings(%v) = %q, expected %q", tt.elems, actual, tt.expected)
			}
		})
	}
}

// TestStrToByte is a test function for StrToByte.
func TestStrToByte(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []byte
	}{
		{
			name:     "Empty string",
			input:    "",
			expected: []byte{},
		},
		{
			name:     "Simple string",
			input:    "hello",
			expected: []byte("hello"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := StrToByte(tt.input)
			fmt.Println(string(actual))
			if len(actual) != len(tt.expected) {
				t.Errorf("StrToByte(%q) expected length %d, actual length %d", tt.input, len(tt.expected), len(actual))
				return
			}

			for i, v := range actual {
				if v != tt.expected[i] {
					t.Errorf("StrToByte(%q) expected byte %v at index %d, actual byte %v", tt.input, tt.expected[i], i, v)
					return
				}
			}
		})
	}
}

func TestByteToStr(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected string
	}{
		{
			name:     "Empty slice",
			input:    []byte{},
			expected: "",
		},
		{
			name:     "Valid string",
			input:    []byte("Hello, World!"),
			expected: "Hello, World!",
		},
		{
			name:     "Special characters",
			input:    []byte{0xFF, 0xFE, 0xFD},
			expected: string([]byte{0xFF, 0xFE, 0xFD}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := ByteToStr(tt.input)
			if actual != tt.expected {
				t.Errorf("ByteToStr(%v): expected %q, actual %q", tt.input, tt.expected, actual)
			}
		})
	}
}
