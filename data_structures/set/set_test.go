package set

import "testing"

func TestSet(t *testing.T) {
	set := New()
	set.AddElement(1)
	set.AddElement(2)
	t.Logf("initial set: %+v\n", set)
	t.Log(set.ContainsElement(1))
	anotherSet := New()
	anotherSet.AddElement(2)
	anotherSet.AddElement(4)
	anotherSet.AddElement(5)
	t.Logf("another set: %+v\n", anotherSet)
	t.Logf("intersect set: %+v\n", set.Intersect(anotherSet))
	t.Logf("union set: %+v\n", set.Union(anotherSet))
}
