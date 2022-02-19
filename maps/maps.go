package maps

// import (
// 	"reflect"
// )

type Map[K comparable, V comparable] struct {
	NativeMap *map[K]V
}

// iterate over keys, 
func MappingByFunction[T any](mapper func(value V, key K, nativeMap *map[K]V) T) *[]T {

	nativeMap := theMap.NativeMap

	var keysArray []K
	keysArray = reflect.ValueOf(nativeMap).MapKeys()
	keysCount := len(keysArray)
	
	for index := 0; index < keysCount; index++ {
		key := keysArray[index]
		value := V(nativeMap[key])

		newArray[index] = mapper(value, key, nativeMap)
	}
	return &newArray
}

func Reduce[T any](mapper func(value V, key K, nativeMap *map[K]V) T) *T {

}

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