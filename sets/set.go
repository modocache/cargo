package sets

type setMap map[interface{}]interface{}
type Set struct {
	elements setMap
}

func NewSet() *Set {
	return &Set{elements: make(setMap)}
}

func (set *Set) Add(element interface{}) {
	set.elements[element] = element
}

func (set *Set) Remove(element interface{}) {
	delete(set.elements, element)
}

func (set *Set) Contains(element interface{}) bool {
	_, contains := set.elements[element]
	return contains
}

func (set *Set) Cardinality() int {
	return len(set.elements)
}

func (set *Set) Elements() []interface{} {
	elements := make([]interface{}, 0, set.Cardinality())
	for _, element := range set.elements {
		elements = append(elements, element)
	}
	return elements
}
