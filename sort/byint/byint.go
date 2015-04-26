package byint

import (
	"fmt"
	"log"
	"reflect"
	"sort"
	"strconv"
)

// byInt implements sort.Interface, where Data is the slice to be sorted.
type byInt struct {
	Data       reflect.Value
	Indices    []int
	Identifier func(interface{}) int
}

// Less is a comparator. Since sort.Sort isn't stable, we use the trick of sorting [a_1, ..., a_n]
// by sorting [(a_1, 1), ..., (a_n, n)] in lexicographical order.
func (b byInt) Less(i, j int) bool {
	ithVal := b.Identifier(b.Data.Index(i).Interface())
	jthVal := b.Identifier(b.Data.Index(j).Interface())
	if ithVal == jthVal {
		return b.Indices[i] < b.Indices[j]
	}
	return ithVal < jthVal
}

// Len returns the length of the underlying data.
func (b byInt) Len() int {
	return b.Data.Len()
}

// Swap interchanges the i-th and j-th entries, also keeping track of their original indices.
func (b byInt) Swap(i, j int) {
	t := reflect.ValueOf(b.Data.Index(i).Interface())
	b.Data.Index(i).Set(b.Data.Index(j))
	b.Data.Index(j).Set(t)
	b.Indices[i], b.Indices[j] = b.Indices[j], b.Indices[i]
}

// DefaultID will use the stringer interface
func DefaultID(i interface{}) int {
	intVal, err := strconv.Atoi(fmt.Sprint(i))
	if err != nil {
		log.Fatalf(err.Error())
	}
	return intVal
}

// Sort is a stable sort that takes a slice as first argument. Will panic if data is not a slice.
func Sort(data interface{}, identifier ...func(interface{}) int) {
	val := reflect.ValueOf(data)
	identifier = append(identifier, DefaultID)
	sortable := byInt{
		Data:       val,
		Indices:    make([]int, val.Len()),
		Identifier: identifier[0],
	}
	for i := 0; i < val.Len(); i++ {
		sortable.Indices[i] = i
	}
	sort.Sort(sortable)
	return
}
