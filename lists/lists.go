package lists

import (
	"errors"
)

/*
	List
	immutable by default.
	Any operations that involve change create a new List and return a pointer to the new List
*/


// TODO - for immutability, does the pointer to `List` need to be new each time, or just the underlying `Array`?
type List[T any] struct {
	Array *[]T
}

// return underlying array
func (list *List[T]) AsArray() []T {
	return *list.Array
}

// returns pointer to the underlying array
func (list *List[T]) AsPointer() *[]T {
	return list.Array
}

func New[T any](array []T) *List[T] {
	return &List[T]{ &array }
}

func NewEmpty[T any]() *List[T] {
	array := []T{}
	return &List[T]{ &array }
}

func (list *List[T]) SetReducer(finder func(element T, index int, array *[]T) bool) (error, *T) {
	for index, value := range *list.Array {
		match := finder(value, index, list.Array)
		if match {
			return nil, &value
		}
	}
	return errors.New("Not Found"), nil
}

func (list *List[T]) Find(finder func(element T, index int, array *[]T) bool) (error, *T) {
	for index, value := range *list.Array {
		match := finder(value, index, list.Array)
		if match {
			return nil, &value
		}
	}
	return errors.New("Not Found"), nil
}

// func (list *List[T]) Reduce(mapper func(T, index int, array *[]T) T) *[]T {

// }

func (list *List[T]) IndexOf(matcher func(element T, index int, array *[]T) bool) (error, int) {
	for index, value := range *list.Array {
		match := matcher(value, index, list.Array)
		if match {
			return nil, index
		}
	}
	return errors.New("Not Found"), -1
}

/* Chainable methods */

// perform a mapping operation over each element in List.Array, return pointer to new List
func (list *List[T]) Map(mapper func(value T, index int, array *[]T) T) *List[T] {

	oldArray := *list.Array
	oldArraySize := len(oldArray)
	newArray := make([]T, oldArraySize) // create new array of same size
	
	for index := 0; index < oldArraySize; index++ {
		value := T(oldArray[index])

		newArray[index] = mapper(value, index, list.Array)
	}
	return &List[T]{ &newArray }
}

// does not return a new List pointer, merely passes each element to `operation` function
func (list *List[T]) ForEach(operation func(element T, index int, array *[]T) T) *List[T] {
	for index, value := range *list.Array {
		operation(value, index, list.Array)
	}
	return list
}

func (list *List[T]) Filter(filterFunc func(element T, index int, array *[]T) bool) *List[T] {
	oldArray := *list.Array
	newArray := make([]T, len(oldArray), len(oldArray))

	counter := 0
	for index, value := range *list.Array {
		ok := filterFunc(value, index, list.Array)
		if ok {
			newArray[counter] = value
			counter++
		}
	}
	slicedDown := newArray[:counter]
	return &List[T]{ &slicedDown }
}

func (list *List[T]) Append(addition *[]T) *List[T] {
	oldArray := list.Array
	newArray := append(*oldArray, *addition...)
	return &List[T]{ &newArray }
}

/* End List's Chainable methods */

// returns `size`, Will return error if underlying array pointer is `nil`
func (list *List[T]) Size() (error, int) {
	arrayPtr := list.Array
	if arrayPtr == nil {
		return errors.New("method called when List.Array == nil"), -1
	}
	theArray := *arrayPtr
	return nil, len(theArray)
}

// returns `true` if list has a `len() > 0`. Will return error if underlying array pointer is `nil`
func (list *List[T]) IsEmpty() (error, bool) {
	theError, size := list.Size()
	if theError != nil {
		return theError, true
	}
	return nil, size == 0
}

func firstOrLastValidation[T any](list *List[T]) (error, []T) {
	arrayPtr := list.Array
	theArray := *arrayPtr
	if arrayPtr == nil {
		var nada []T
		return errors.New(".First() called on List where `List.Array === nil`"), nada
	}
	errur, isEmpty := list.IsEmpty()
	if errur != nil {
		var nada []T
		return errur, nada
	} 
	if isEmpty {
		var nada []T
		return errors.New("Empty"), nada
	}
	return nil, theArray
}

func (list *List[T]) Last() (error, T) {
	error, array := firstOrLastValidation[T](list)
	if error != nil {
		var nada T
		return error, nada
	}
	return nil, array[len(array)- 1]
}

func (list *List[T]) First() (error, T) {
	error, array := firstOrLastValidation[T](list)
	if error != nil {
		var nada T
		return error, nada
	}
	return nil, array[0]
}

/* End of List functions & methods */

/* ReduceList */
type ReduceList[T any, R any] struct {
	Array *[]T
	Reducer func(accumulator R, value T, index int, array *[]T) R
}

func (list *ReduceList[T, R]) Reduce(inital R) R {
	reducer := list.Reducer
	array := *list.Array

	accumulator := inital
	for index, value := range array {
		accumulator = reducer(accumulator, value, index, list.Array)
	}
	return accumulator	 
}

/* End Of Reduce List */

/* List interface */
// type ListShape[T any, R any] interface {
// 	List[T] | ReduceList[T,R]
// }

// func (list *ListShape[T,R]) AsReduceList(reducer func(accumulator R, value T, index int, array *[]T) R) *ReduceList[T,R] {
// 	return &ReduceList{ list.Array, reducer }
// }
/* */

/* MutableList */

// type MutableList[T comparable] struct {
// 	Array *[]T
// }

/* Chainable */

// func (list *List[T]) NewMutable() *List[T] {

// }

// func (list *List[T]) SetSize(mapper func(T, index int, array *[]T) T) *List[T] {

// }

// func (list *List[T]) SetCapacity(mapper func(T, index int, array *[]T) T) *List[T] {

// }

// Removes all elements for which `filter` returns `true`. This method can take up to O(N)
// func (list *MutableList[T]) Filter(filter func(a T, index int, list *MutableList[T]) bool) *MutableList[T] {

// 	var targetIndex int = -1
// 	for index, current := range *list.Array {
// 		if filter(current, ) {
// 			targetIndex = index
// 		}
// 	}
// 	theArray := *list.Array
// 	newSlice := append(theArray[:targetIndex], theArray[targetIndex:]...)
// 	list.Array = &newSlice
// 	return nil
// }

/* End Chainable */

// func simpleEqualityComparator[T comparable](a T, b T) bool {
// 	return a == b
// }

// Removes an element via simple `==` equality check. This method can take up to O(N)
// func (list *MutableList[T]) Remove(value T) error  {
// 	return list.RemoveByComparator(value, simpleEqualityComparator[T])
// }

/* End MutableList */