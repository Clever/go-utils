package stringset

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClone(t *testing.T) {
	initialSet := FromList([]string{"test"})
	clonedSet := initialSet.Clone()

	assert.Equal(t, initialSet, clonedSet)

	delete(clonedSet, "test")
	assert.Equal(t, 0, len(clonedSet))
	assert.Equal(t, 1, len(initialSet))
}

func TestIntersect(t *testing.T) {
	setA := FromList([]string{"a", "b"})
	setB := FromList([]string{"b", "c"})

	intersection := setA.Intersect(setB)
	assert.Equal(t, intersection["b"], struct{}{})
	assert.Equal(t, 1, len(intersection))
}

func TestMinus(t *testing.T) {
	setA := FromList([]string{"a", "b"})
	setB := FromList([]string{"b", "c"})

	minusSet := setA.Minus(setB)
	assert.Equal(t, minusSet["a"], struct{}{})
	assert.Equal(t, 1, len(minusSet))
}

func TestAddingToSet(t *testing.T) {
	setBase := FromList([]string{"a", "b"})

	setBase.AddAll([]string{"a", "c"})
	setBase.AddSet(FromList([]string{"b", "d"}))

	assert.Equal(t, 4, len(setBase))
}

func TestEqual(t *testing.T) {
	empty1 := FromList([]string{})
	empty2 := FromList([]string{})
	assert.True(t, empty1.Equals(empty2))

	oneVariable := FromList([]string{"a"})
	sameVariable := FromList([]string{"a"})
	assert.True(t, oneVariable.Equals(sameVariable))
	assert.False(t, empty1.Equals(oneVariable))

	otherVariable := FromList([]string{"b"})
	assert.False(t, oneVariable.Equals(otherVariable))

	twoVariables := FromList([]string{"a", "b"})
	assert.False(t, oneVariable.Equals(twoVariables))
}

func TestPartition(t *testing.T) {
	a, b, c := Partition([]string{"a", "b"}, []string{"b", "c"})
	assert.Equal(t, a, []string{"a"})
	assert.Equal(t, b, []string{"b"})
	assert.Equal(t, c, []string{"c"})
}

func TestAddRemoveContains(t *testing.T) {
	set := New("a", "b")	
	assert.True(t, set.Contains("a"))

	set.Remove("a")
	assert.False(t, set.Contains("a"))

	set.Add("a")
	assert.True(t, set.Contains("a"))
}
