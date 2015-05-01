package bystring

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type StructuredData struct {
	Field1, Field2 string
}

func (s StructuredData) String() string {
	return s.Field1 + s.Field2
}

func TestHappyPath(t *testing.T) {
	actual := []string{"c", "b", "a", "d"}
	expected := []string{"a", "b", "c", "d"}
	Sort(actual, func(i interface{}) string { return i.(string) })
	assert.Equal(t, expected, actual)
}

func TestStableSort(t *testing.T) {
	actual := []StructuredData{
		{Field1: "a", Field2: "zz"},
		{Field1: "z", Field2: "g"},
		{Field1: "z", Field2: "h"},
		{Field1: "d", Field2: "h"},
		{Field1: "z", Field2: "i"},
		{Field1: "z", Field2: "1"},
		{Field1: "z", Field2: "o"},
		{Field1: "d", Field2: "h"},
		{Field1: "a", Field2: "z"},
		{Field1: "z", Field2: "a"},
		{Field1: "a", Field2: "a"},
	}
	expected := []StructuredData{
		{Field1: "a", Field2: "zz"},
		{Field1: "a", Field2: "z"},
		{Field1: "a", Field2: "a"},
		{Field1: "d", Field2: "h"},
		{Field1: "d", Field2: "h"},
		{Field1: "z", Field2: "g"},
		{Field1: "z", Field2: "h"},
		{Field1: "z", Field2: "i"},
		{Field1: "z", Field2: "1"},
		{Field1: "z", Field2: "o"},
		{Field1: "z", Field2: "a"},
	}
	Sort(actual, func(i interface{}) string { return i.(StructuredData).Field1 })
	assert.Equal(t, expected, actual)
}

func TestSingletonSlice(t *testing.T) {
	actual := []int{5}
	expected := []int{5}
	Sort(actual)
	assert.Equal(t, expected, actual)
}

func TestEmptySlice(t *testing.T) {
	actual := []int{}
	expected := []int{}
	Sort(actual)
	assert.Equal(t, expected, actual)
}

func TestNilSlice(t *testing.T) {
	actual := []int(nil)
	expected := []int(nil)
	Sort(actual)
	assert.Equal(t, expected, actual)
}

func TestStringerSlice(t *testing.T) {
	actual := []StructuredData{{"c", "c"}, {"b", "b"}, {"a", "a"}}
	expected := []StructuredData{{"a", "a"}, {"b", "b"}, {"c", "c"}}
	Sort(actual)
	assert.Equal(t, expected, actual)
}

func TestConversion(t *testing.T) {
	actual := []error{fmt.Errorf("B"), fmt.Errorf("A")}
	expected := []error{fmt.Errorf("A"), fmt.Errorf("B")}
	Sort(actual)
	assert.Equal(t, expected, actual)
}
