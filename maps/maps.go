package maps

import (
	// "reflect"
	"github.com/neurocollective/go_chainable/lists"
)

// ordered HashMap type, wrapper around native map
// because of key array, removes are slow
// K
type Map[K comparable, V comparable, R any] struct {
	NativeMap *map[K]V
	KeysList *lists.List[K, R]
}

func (h *Map[K, V, R]) Add(key K, value V) *Map[K, V, R] {

	// check if map already has key. if so, do nothing.

	(*h.NativeMap)[key] = value
	// add to keys
	return h
}

func (h *Map[K, V, R]) Remove(key K) *Map[K, V, R] {
	// remove from h.NativeMap,
	// use h.KeysList.Remove(key)
	return h
}

// returns keys in order of being added
// removed keys are gone and no longer part of the order
func (h *Map[K, V, R]) Keys() *lists.List[K, R] {
	return h.KeysList
}

func (h *Map[K, V, R]) Values() *lists.List[V, R] {
	size := len(*h.KeysList.Array)
	values := make([]V, size)

	for index, key := range *h.KeysList.Array {
		values[index] = (*h.NativeMap)[key]
	}
	var cypher R
	return &lists.List[V, R]{ &values, cypher }
}

func New[K comparable, V comparable, R any]() *Map[K, V, R] {
	
	nativeMap := make(map[K]V)
	array := make([]K, 0, 50)
	var val R
	newMap := Map[K, V, R]{ &nativeMap, &lists.List[K, R]{ &array, val } }
	return &newMap
}

func (theMap *Map[K, V, R]) Map(mapper func(value V, key K, index int, nativeMap *map[K]V) R) *[]R {

	nativeMap := theMap.NativeMap
	keysArray := *theMap.KeysList.Array
	keysCount := len(keysArray)
	
	newArray := make([]R, keysCount)

	for index, key := range keysArray {
		value := (*nativeMap)[key]
		newArray[index] = mapper(value, key, index, nativeMap)
	}
	return &newArray
}

func (theMap *Map[K, V, R]) Reduce(
	reducer func(accumulator R, value V, key K, index int, nativeMap *map[K]V) R,
	initial R,
) R {

	nativeMap := theMap.NativeMap
	keysArray := *theMap.KeysList.Array

	accumulator := initial
	for index, key := range keysArray {
		value := (*nativeMap)[key]
		accumulator = reducer(accumulator, value, key, index, nativeMap)
	}
	return accumulator
}
