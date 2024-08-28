package sliceutil

import "sort"

// compareFunction 是一个比较函数，用于比较两个元素。
type compareFunction func(i, j int) bool

// Sorter 是一个泛型结构体，它允许我们为任何类型的切片创建一个可排序的对象。
type Sorter[T any] struct {
	slice []T
	cmp   compareFunction
}

// Len 是 sort.Interface 的一部分。
func (s *Sorter[T]) Len() int { return len(s.slice) }

// Swap 是 sort.Interface 的一部分。
func (s *Sorter[T]) Swap(i, j int) { s.slice[i], s.slice[j] = s.slice[j], s.slice[i] }

// Less 是 sort.Interface 的一部分。
func (s *Sorter[T]) Less(i, j int) bool {
	return s.cmp(i, j)
}

// MakeSorter 创建一个新的泛型排序器，并接收一个比较函数。
func MakeSorter[T any](s []T, cmp compareFunction) *Sorter[T] {
	aa := Sorter[T]{slice: s, cmp: cmp}
	return &aa
}

// Sort 排序方法，调用 sort.Sort 进行排序。
func (s *Sorter[T]) Sort() {
	sort.Sort(s)
}
