package generalutil

import (
	"testing"
)

func TestBeautifyPrintStruct(t *testing.T) {
	type User struct {
		Name    string
		Age     int
		Address []int
	}
	u := User{Name: "John", Age: 30, Address: []int{1, 2, 3, 4}}
	PrettyPrintStruct(u)
}
