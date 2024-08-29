package generalutil

import (
	"testing"
)

func TestBeautifyPrintStruct(t *testing.T) {
	type User struct {
		Name string
		Age  int
	}
	u := User{Name: "John", Age: 30}
	PrettyPrintStruct(u)
}
