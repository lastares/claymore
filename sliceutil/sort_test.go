package sliceutil

import (
	"reflect"
	"testing"
)

// 测试 MakeSorter 函数和 Sort 方法
func TestNewSorterWithInt(t *testing.T) {
	s := []int{9, 2, 8, 1, 6, 0, 3, 7, 4, 5, 10}
	tests := []struct {
		name     string
		slice    []int
		cmp      compareFunction
		expected []int
	}{
		{
			name:     "Empty slice",
			slice:    []int{},
			cmp:      func(i, j int) bool { return i < j },
			expected: []int{},
		},
		//{
		//	name:     "Non-empty slice",
		//	slice:    s,
		//	cmp:      func(i, j int) bool { return s[i] < s[j] },
		//	expected: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		//},
		{
			name:     "int slice",
			slice:    s,
			cmp:      func(i, j int) bool { return s[i] > s[j] },
			expected: []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MakeSorter(tt.slice, tt.cmp).Sort()
			t.Log("排序后的slice: ", tt.slice)
			if !reflect.DeepEqual(tt.slice, tt.expected) {
				t.Errorf("After sorting: got %v, want %v", tt.slice, tt.expected)
			}
		})
	}
}

func TestNewSorterWithString(t *testing.T) {
	s := []string{"b", "c", "a", "d"}
	tests := []struct {
		name     string
		slice    []string
		cmp      compareFunction
		expected []string
	}{
		{
			name:     "Empty slice",
			slice:    []string{},
			cmp:      func(i, j int) bool { return s[i] < s[j] },
			expected: []string{},
		},
		{
			name:     "string slice",
			slice:    s,
			cmp:      func(i, j int) bool { return s[i] < s[j] },
			expected: []string{"a", "b", "c", "d"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MakeSorter(tt.slice, tt.cmp).Sort()
			t.Log("排序后的slice: ", tt.slice)
			if !reflect.DeepEqual(tt.slice, tt.expected) {
				t.Errorf("After sorting: got %v, want %v", tt.slice, tt.expected)
			}
		})
	}
}

func TestNewSorterWithStruct(t *testing.T) {
	type Exam struct {
		Name  string
		Score int
	}
	s := []*Exam{
		{Name: "aaa", Score: 90},
		{Name: "ddd", Score: 80},
		{Name: "ccc", Score: 70},
		{Name: "bbb", Score: 60},
	}
	tests := []struct {
		name     string
		slice    []*Exam
		cmp      compareFunction
		expected []*Exam
	}{
		{
			name:     "Empty slice",
			slice:    []*Exam{},
			cmp:      func(i, j int) bool { return s[i].Score < s[j].Score },
			expected: []*Exam{},
		},
		//{
		//	name:  "struct slice",
		//	slice: s,
		//	cmp:   func(i, j int) bool { return s[i].Score < s[j].Score },
		//	expected: []*Exam{
		//		{Name: "bbb", Score: 60},
		//		{Name: "ccc", Score: 70},
		//		{Name: "ddd", Score: 80},
		//		{Name: "aaa", Score: 90},
		//	},
		//},
		{
			name:  "struct slice",
			slice: s,
			cmp:   func(i, j int) bool { return s[i].Name < s[j].Name },
			expected: []*Exam{
				{Name: "aaa", Score: 90},
				{Name: "bbb", Score: 60},
				{Name: "ccc", Score: 70},
				{Name: "ddd", Score: 80},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MakeSorter(tt.slice, tt.cmp).Sort()
			for _, ss := range tt.slice {
				t.Logf("排序后的slice: name: %s, score: %d", ss.Name, ss.Score)
			}
			if !reflect.DeepEqual(tt.slice, tt.expected) {
				t.Errorf("After sorting: got %v, want %v", tt.slice, tt.expected)
			}
		})
	}
}
