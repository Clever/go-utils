package byint

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

type StructuredData struct {
	Field1 int
	Field2 string
}

func (s StructuredData) String() string {
	return strconv.Itoa(s.Field1)
}

func TestHappyPath(t *testing.T) {
	actual := []int{3, 2, 1, 4}
	expected := []int{1, 2, 3, 4}
	Sort(actual, func(i interface{}) int { return i.(int) })
	assert.Equal(t, expected, actual)
}

func TestStableSort(t *testing.T) {
	actual := []StructuredData{
		{Field1: 1, Field2: "zz"},
		{Field1: 3, Field2: "g"},
		{Field1: 3, Field2: "h"},
		{Field1: 2, Field2: "h"},
		{Field1: 3, Field2: "i"},
		{Field1: 3, Field2: "1"},
		{Field1: 3, Field2: "o"},
		{Field1: 2, Field2: "h"},
		{Field1: 1, Field2: "3"},
		{Field1: 3, Field2: "1"},
		{Field1: 1, Field2: "1"},
	}
	expected := []StructuredData{
		{Field1: 1, Field2: "zz"},
		{Field1: 1, Field2: "3"},
		{Field1: 1, Field2: "1"},
		{Field1: 2, Field2: "h"},
		{Field1: 2, Field2: "h"},
		{Field1: 3, Field2: "g"},
		{Field1: 3, Field2: "h"},
		{Field1: 3, Field2: "i"},
		{Field1: 3, Field2: "1"},
		{Field1: 3, Field2: "o"},
		{Field1: 3, Field2: "1"},
	}
	Sort(actual, func(i interface{}) int { return i.(StructuredData).Field1 })
	assert.Equal(t, expected, actual)
}

func TestSingletonSlice(t *testing.T) {
	actual := []string{"5"}
	expected := []string{"5"}
	Sort(actual)
	assert.Equal(t, expected, actual)
}

func TestEmptySlice(t *testing.T) {
	actual := []string{}
	expected := []string{}
	Sort(actual)
	assert.Equal(t, expected, actual)
}

func TestNilSlice(t *testing.T) {
	actual := []string(nil)
	expected := []string(nil)
	Sort(actual)
	assert.Equal(t, expected, actual)
}

func TestStringerSlice(t *testing.T) {
	actual := []StructuredData{{22, "c"}, {3, "b"}, {1, "a"}}
	expected := []StructuredData{{1, "a"}, {3, "b"}, {22, "c"}}
	Sort(actual)
	assert.Equal(t, expected, actual)
}

func TestConversion(t *testing.T) {
	actual := []string{"11", "2"}
	expected := []string{"2", "11"}
	Sort(actual)
	assert.Equal(t, expected, actual)
}
