package lists

type List[T any] struct {
	Array *[]T
}

func (list *List[T]) AsArray() []T {
	return *list.Array
}

func (list *List[T]) AsArrayPointer() *[]T {
	return list.Array
}

// perform a mapping operation over each element in List.Array, return new raw array
func (list *List[T]) Map(mapper func(value T, index int, array *[]T) T) *[]T {

	oldArray := *list.Array
	oldArraySize := len(oldArray)
	newArray := make([]T, oldArraySize) // create new array of same size
	
	for index := 0; index < oldArraySize; index++ {
		value := T(oldArray[index])

		newArray[index] = mapper(value, index, list.Array)
	}
	return &newArray
}

// func BuildList[T any](array []T) *List[T] {
// 	return &List{ &array }
// }

// func (arrayList *List[T]) Find(mapper func(T, index int, array *[]T) T) *[]T {

// }

// func (arrayList *List[T]) Reduce(mapper func(T, index int, array *[]T) T) *[]T {

// }

// func (arrayList *List[T]) ForEach(mapper func(T, index int, array *[]T) T) *[]T {

// }

// func (arrayList *List[T]) Filter(mapper func(T, index int, array *[]T) T) *[]T {

// }

// func (arrayList *List[T]) IndexOf(mapper func(T, index int, array *[]T) T) *[]T {

// }
