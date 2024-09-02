package maputil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrderedMap_Set_Get(t *testing.T) {
	om := NewOrderedMap[string, int]()
	om.Set("a", 1)
	om.Set("b", 2)
	om.Set("c", 3)
	val, ok := om.Get("a")
	assert.Equal(t, 1, val)
	assert.Equal(t, true, ok)
	val, ok = om.Get("d")
	assert.Equal(t, false, ok)
	assert.Equal(t, 0, val)
}

func TestOrderedMap_Delete_Clear(t *testing.T) {
	om := NewOrderedMap[string, int]()
	om.Set("a", 1)
	om.Set("b", 2)
	om.Set("c", 3)
	assert.Equal(t, 3, om.Len())
	om.Delete("b")
	assert.Equal(t, 2, om.Len())
	om.Clear()
	assert.Equal(t, 0, om.Len())
}

func TestOrderedMap_Keys_Values(t *testing.T) {
	om := NewOrderedMap[string, int]()
	om.Set("a", 1)
	om.Set("b", 2)
	om.Set("c", 3)
	assert.Equal(t, []string{"a", "b", "c"}, om.Keys())
	assert.Equal(t, []int{1, 2, 3}, om.Values())
}

func TestOrderedMap_Contains(t *testing.T) {
	om := NewOrderedMap[string, int]()
	om.Set("a", 1)
	om.Set("b", 2)
	om.Set("c", 3)
	assert.Equal(t, true, om.ContainKey("a"))
	assert.Equal(t, false, om.ContainKey("d"))
}

func TestOrderedMap_Elements(t *testing.T) {
	om := NewOrderedMap[string, int]()
	om.Set("a", 1)
	om.Set("b", 2)
	om.Set("c", 3)
	elements := []struct {
		Key   string
		Value int
	}{
		{"a", 1},
		{"b", 2},
		{"c", 3},
	}
	assert.Equal(t, elements, om.Elements())
}

func TestOrderedMap_Range(t *testing.T) {
	om := NewOrderedMap[string, int]()
	om.Set("a", 1)
	om.Set("b", 2)
	om.Set("c", 3)
	om.Set("d", 4)
	var keys []string
	om.Range(func(key string, value int) bool {
		keys = append(keys, key)
		return key != "c"
	})
	assert.Equal(t, []string{"a", "b", "c"}, keys)
}

func TestOrderedMap_ReverseIter(t *testing.T) {
	om := NewOrderedMap[string, int]()
	om.Set("a", 1)
	om.Set("b", 2)
	om.Set("c", 3)
	var items []struct {
		Key   string
		Value int
	}
	iterCh := om.ReverseIter()
	for item := range iterCh {
		items = append(items, item)
	}
	expected := []struct {
		Key   string
		Value int
	}{
		{"c", 3},
		{"b", 2},
		{"a", 1},
	}
	assert.Equal(t, expected, items)
}

func TestOrderedMap_SortByKey(t *testing.T) {
	om := NewOrderedMap[string, int]()
	om.Set("d", 4)
	om.Set("b", 2)
	om.Set("c", 3)
	om.Set("a", 1)
	om.SortByKey(func(a, b string) bool {
		return a < b
	})
	assert.Equal(t, []string{"a", "b", "c", "d"}, om.Keys())
}

func TestOrderedMap_MarshalJSON(t *testing.T) {
	om := NewOrderedMap[string, int]()
	om.Set("a", 1)
	om.Set("b", 2)
	om.Set("c", 3)
	jsonBytes, err := om.MarshalJSON()
	if err != nil {
		t.Errorf("MarshalJSON error: %v", err)
	}
	assert.Equal(t, `{"a":1,"b":2,"c":3}`, string(jsonBytes))
}

func TestOrderedMap_UnmarshalJSON(t *testing.T) {
	jsonStr := `{"a":1,"b":2,"c":3}`
	om := NewOrderedMap[string, int]()
	err := om.UnmarshalJSON([]byte(jsonStr))
	if err != nil {
		t.Errorf("MarshalJSON error: %v", err)
	}
	assert.Equal(t, 3, om.Len())
	assert.Equal(t, true, om.ContainKey("a"))
	assert.Equal(t, true, om.ContainKey("b"))
	assert.Equal(t, true, om.ContainKey("c"))
}

func TestOrderedMap_Front_Back(t *testing.T) {
	om := NewOrderedMap[string, int]()
	om.Set("a", 1)
	om.Set("b", 2)
	om.Set("c", 3)
	frontElement, ok := om.Front()
	assert.Equal(t, "a", frontElement.Key)
	assert.Equal(t, 1, frontElement.Value)
	assert.Equal(t, true, ok)
	backElement, ok := om.Back()
	assert.Equal(t, "c", backElement.Key)
	assert.Equal(t, 3, backElement.Value)
	assert.Equal(t, true, ok)
}
