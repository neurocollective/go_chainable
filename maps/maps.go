package maps

import (
	// "reflect"
	"github.com/neurocollective/godash/lists"
)

// ordered HashMap type, wrapper around native map
// because of key array, removes are slow
type Map[K comparable, V comparable] struct {
	NativeMap *map[K]V
	KeysList *lists.List[K]
}

func (h *Map[K, V]) Add(key K, value V) *Map[K, V] {

	// check if map already has key. if so, do nothing.

	(*h.NativeMap)[key] = value
	// add to keys
	return h
}

func (h *Map[K, V]) Remove(key K) *Map[K, V] {
	// remove from h.NativeMap,
	// use h.KeysList.Remove(key)
	return h
}

// returns keys in order of being added
// removed keys are gone and no longer part of the order
func (h *Map[K, V]) Keys() *[]K {
	return h.KeysList.Array
}

func (h *Map[K, V]) Values() *[]V {
	size := len(*h.KeysList.Array)
	values := make([]V, size)

	// for i = 0; i < size; i++ {
	// 	key := h.KeysList.Array[i]
	// 	values[i] = h.Map
	// }

	for index, key := range *h.KeysList.Array {
		values[index] = (*h.NativeMap)[key]
	}
	return &values
}

func New[K comparable, V comparable]() *Map[K, V] {
	
	nativeMap := make(map[K]V)
	array := make([]K, 0, 50)
	newMap := Map[K, V]{ &nativeMap, &lists.List[K]{ &array } }
	return &newMap
}

func FunctionalMapping[K comparable, V comparable, T any](
	theMap *Map[K, V],
	mapper func(value V, key K, nativeMap *map[K]V) T,
) *[]T {

	nativeMap := theMap.NativeMap
	keysArray := *theMap.KeysList.Array
	keysCount := len(keysArray)
	
	newArray := make([]T, keysCount)

	for index, key := range keysArray {
		value := (*nativeMap)[key]
		newArray[index] = mapper(value, key, nativeMap)
	}

	return &newArray
}

// func Reduce[T any](mapper func(value V, key K, nativeMap *map[K]V) T) *T {

// }

// func (theMap *Map[K, V]) MapByKeys(mapper func(value V, key K, nativeMap *map[K]V) K) *[]K {

// 	nativeMap := theMap.NativeMap

// 	var keysArray []K
// 	keysArray = reflect.ValueOf(nativeMap).MapKeys()
// 	keysCount := len(keysArray)
	
// 	for index := 0; index < keysCount; index++ {
// 		key := keysArray[index]
// 		value := V(nativeMap[key])

// 		newArray[index] = mapper(value, key, nativeMap)
// 	}
// 	return &newArray
// }

// func (theMap *Map[K, V]) MapByValues(mapper func(value V, key K, nativeMap *map[K]V) V) *[]V {

// 	nativeMap := theMap.NativeMap

// 	var keysArray []K
// 	keysArray = reflect.ValueOf(nativeMap).MapKeys()
// 	keysCount := len(keysArray)
	
// 	for index := 0; index < keysCount; index++ {
// 		key := keysArray[index]
// 		value := V(nativeMap[key])

// 		newArray[index] = mapper(value, key, nativeMap)
// 	}
// 	return &newArray
// }

// func (list *List[T]) Reduce(mapper func(value V, key K, map *[]T) T) *[]T {

// 	oldArray := *list.Array
// 	oldArraySize := len(oldArray)
// 	newArray := make([]T, oldArraySize) // create new array of same size
	
// 	for index := 0; index < oldArraySize; index++ {
// 		value := T(oldArray[index])

// 		newArray[index] = mapper(value, index, list.Array)
// 	}
// 	return &newArray
// }