package containers

import "github.com/AlexZ33/utils/algorithm"

// Container is base interface that all data structures implement.
type Container interface {
	Empty() bool
	Size() int
	Clear()
	Values() []interface{}
	String() string
}

func GetSortedValues(container Container, comparator algorithm.Comparator) []interface{} {
	values := container.Values()
	if len(values) < 2 {
		return values
	}
	algorithm.Sort(values, comparator)
	return values
}
