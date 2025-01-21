package set

type Set struct {
	integerMap map[int]bool
}

func New() *Set {
	set := &Set{
		integerMap: make(map[int]bool),
	}
	return set
}

func (set *Set) ContainsElement(element int) bool {
	_, exists := set.integerMap[element]
	return exists
}

func (set *Set) AddElement(element int) {
	if !set.ContainsElement(element) {
		set.integerMap[element] = true
	}
}

func (set *Set) DeleteElement(element int) {
	delete(set.integerMap, element)
}

func (set *Set) Intersect(anotherSet *Set) *Set {
	intersectSet := New()
	for value := range set.integerMap {
		if anotherSet.ContainsElement(value) {
			intersectSet.AddElement(value)
		}
	}
	return intersectSet
}

func (set *Set) Union(anotherSet *Set) *Set {
	unionSet := New()
	for value := range set.integerMap {
		unionSet.AddElement(value)
	}
	for value := range anotherSet.integerMap {
		unionSet.AddElement(value)
	}
	return unionSet
}
