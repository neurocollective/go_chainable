package lists

import (
	"fmt"
	"errors"
)

/*
	List
	immutable by default.
	Any operations that involve change create a new List and return a pointer to the new List
*/


// TODO - for immutability, does the pointer to `List` need to be new each time, or just the underlying `Array`?
type List[T any, R any] struct {
	Array *[]T
	Value R
}

// return underlying array
func (list *List[T, R]) Raw() []T {
	return *list.Array
}

// returns pointer to the underlying array
func (list *List[T, R]) RawPointer() *[]T {
	return list.Array
}

func New[T any, R any](array []T) *List[T, R] {
	var val R
	return &List[T,R]{ &array, val }
}

func NewEmpty[T any, R any]() *List[T, R] {
	var val R
	var array []T
	return &List[T,R]{ &array, val }
}

/* Chainable methods */
// perform a mapping operation over each element in List.Array, return pointer to new List
func (list *List[T, R]) Map(mapper func(value T, index int) T) *List[T, R] {

	oldArray := *list.Array
	oldArraySize := len(oldArray)
	newArray := make([]T, oldArraySize) // create new array of same size
	
	for index := 0; index < oldArraySize; index++ {
		value := T(oldArray[index])

		newArray[index] = mapper(value, index)
	}
	return &List[T, R]{ &newArray, list.Value }
}

// same as Map, but function gets access to array itself
func (list *List[T, R]) MapFull(mapper func(value T, index int, array *[]T) T) *List[T, R] {

	oldArray := *list.Array
	oldArraySize := len(oldArray)
	newArray := make([]T, oldArraySize) // create new array of same size
	
	for index := 0; index < oldArraySize; index++ {
		value := T(oldArray[index])

		newArray[index] = mapper(value, index, list.Array)
	}
	return &List[T, R]{ &newArray, list.Value }
}

func (list *List[T, R]) Reduce(
	reducer func(accumulator R, value T, index int) R,
	initial R,
) R {
	array := *list.Array

	accumulator := initial
	for index, value := range array {
		accumulator = reducer(accumulator, value, index)
	}
	return accumulator
}

func (list *List[T, R]) ReduceFull(
	reducer func(accumulator R, value T, index int, array *[]T) R,
	initial R,
) R {
	array := *list.Array

	accumulator := initial
	for index, value := range array {
		accumulator = reducer(accumulator, value, index, list.Array)
	}
	return accumulator
}

// does not return a new List pointer, merely passes each element to `operation` function
func (list *List[T, R]) ForEach(operation func(element T, index int) T) *List[T, R] {
	for index, value := range *list.Array {
		operation(value, index)
	}
	return list
}

func (list *List[T, R]) ForEachFull(operation func(element T, index int, array *[]T) T) *List[T, R] {
	for index, value := range *list.Array {
		operation(value, index, list.Array)
	}
	return list
}

func (list *List[T, R]) Filter(filterFunc func(element T, index int) bool) *List[T, R] {
	oldArray := *list.Array
	newArray := make([]T, len(oldArray), len(oldArray))

	counter := 0
	for index, value := range *list.Array {
		ok := filterFunc(value, index)
		if ok {
			newArray[counter] = value
			counter++
		}
	}
	slicedDown := newArray[:counter]
	return &List[T, R]{ &slicedDown, list.Value }
}

func (list *List[T, R]) FilterFull(filterFunc func(element T, index int, array *[]T) bool) *List[T, R] {
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
	return &List[T, R]{ &slicedDown, list.Value }
}

// TODO - implement sort method
// func (list *List[T, R]) Sort(compare func(element T, index int) *List[T, R]*List[T, R] {
// }

// func (list *List[T, R]) SortFull(compare func(element T, index int, array *[]T) bool) *List[T, R] {
// }

func (list *List[T, R]) Append(addition *[]T) *List[T, R] {
	oldArray := list.Array
	newArray := append(*oldArray, *addition...)
	return &List[T, R]{ &newArray, list.Value }
}

func (list *List[T, R]) Add(value T) *List[T, R] {
	oldArray := list.Array
	newArray := append(*oldArray, value)
	list.Array = &newArray
	return list
}

// may `panic` if capacity is less than current length of underlying array
func (list *List[T, R]) SetCap(capacity int) *List[T, R] {
	theArray := *list.Array
	if (list.Array == nil) {
		newSlice := make([]T, 0, capacity)
		list.Array = &newSlice 
		return list
	}
	size := len(theArray)
	newSlice := append(make([]T, size, capacity), theArray...)
	list.Array = &newSlice
	return list
}

func (list *List[T, R]) IncrementCap(capacity int) *List[T, R] {
	theArray := *list.Array
	if (list.Array == nil) {
		return list.SetCap(capacity)
	}
	size := len(theArray)
	oldCap := cap(theArray)
	newSlice := append(make([]T, size, oldCap + capacity), theArray...)
	list.Array = &newSlice
	return list
}

/* End List's Chainable methods */

func (list *List[T, R]) Find(finder func(element T, index int) bool) (error, *T) {
	for index, value := range *list.Array {
		match := finder(value, index)
		if match {
			return nil, &value
		}
	}
	return errors.New("Not Found"), nil
}

func (list *List[T, R]) String() string {
	return fmt.Sprint((*list.Array))
}

// returns first match from matcher function
func (list *List[T, R]) IndexOf(matcher func(element T) bool) (error, int) {
	for index, value := range *list.Array {
		match := matcher(value)
		if match {
			return nil, index
		}
	}
	return errors.New("Not Found"), -1
}

// returns `size`, Will return error if underlying array pointer is `nil`
func (list *List[T, R]) Size() (error, int) {
	arrayPtr := list.Array
	if arrayPtr == nil {
		return errors.New("method called when List.Array == nil"), -1
	}
	theArray := *arrayPtr
	return nil, len(theArray)
}

func (list *List[T, R]) Cap() (error, int) {
	arrayPtr := list.Array
	if arrayPtr == nil {
		return errors.New("method called when List.Array == nil"), -1
	}
	theArray := *arrayPtr
	return nil, cap(theArray)
}

// returns `true` if list has a `len() > 0`. Will return error if underlying array pointer is `nil`
func (list *List[T, R]) IsEmpty() (error, bool) {
	theError, size := list.Size()
	if theError != nil {
		return theError, true
	}
	return nil, size == 0
}

func listValidation[T any, R any](list *List[T, R]) (error, []T) {
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

func (list *List[T, R]) Get(index int) (error, T) {
	error, array := listValidation[T, R](list)
	if error != nil {
		var nada T
		return error, nada
	}
	size := len(array)
	if index < 0 || index >= size {
		var nada T
		return errors.New(".Get() seeking an index out of bound"), nada
	}
	return nil, array[index]
}

func (list *List[T, R]) Last() (error, T) {
	error, array := listValidation[T, R](list)
	if error != nil {
		var nada T
		return error, nada
	}
	return nil, array[len(array)- 1]
}

func (list *List[T, R]) First() (error, T) {
	return list.Get(0)
}

/* Standalone Functions */

func ResultTypeSwap[T any, OldR any, NewR any] (list *List[T, OldR]) *List[T, NewR] {
	var value NewR
	newList := List[T, NewR]{ list.Array, value }
	return &newList
}