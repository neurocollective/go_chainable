# go_chainable

`NOTE: Requires go1.18`

Go-Chainable is a library using generics to mimic the functional `.map(x -> y).reduce(x -> y, z)` patterns of javascript. `List` is also inspired by Java's `ArrayList`.

## List

### Construct a List

```
array := []string { "hello", "world" }
list := lists.New[string, any](array)
```

or get an empty list

```
list := lists.NewEmpty[string, any]()
```

#### First type parameter

First type parameter is the type of the List's values.

#### "Result" type parameter

The second type parameter for a list is needed only for when calling `list.Reduce()`. If you won't need to call `.Reduce()` then set the second type argument as `any` and forget it.

### Map Over A List

```
import "github.com/neurocollective/go_chainable/lists"

array := []int{1, 2}
list := lists.New[int, any](array)
doubled := list.Map(func(val int, index int) int {
	return val * 2
	})
fmt.Println(doubled.String()) // -> [2 4]
```

### Filter A List

```
import "github.com/neurocollective/go_chainable/lists"
array := []int { 1, 2 }
list := lists.New[int, any](array)
newList := list.Filter(func(val int, index int) bool {
	return val < 2
})
error, size := newList.Size()
fmt.Println(size) // -> 1
```

### Reduce A List

```
import "github.com/neurocollective/go_chainable/lists"

array := []int { 1, 2 }
list := lists.New[int, int](array) // second type passed to New is the "result" type used by reducer as return type
added := list.Reduce(func(accumulator int, val int, index int) int {
	return accumulator + val
}, 0)
fmt.Println(added) // -> 3
```

### Chain Operations Over A List

```
array := []int { 1, 2, 3 }
list := lists.New[int, int](array)
added := list.Map(func(val int, index int) int {
	return val + 1
}).Filter(func(val int, index int) bool {
	return val < 4
}).Reduce(func(accumulator int, val int, index int) int {
	return accumulator + val
}, 0)
fmt.Println(added) // -> 5
```

## Map

### `.Map` over key/value pairs

```
import "github.com/neurocollective/go_chainable/maps"

theMap := maps.NewEmpty[string, string, string]()

theMap.Set("hey", "dude")
theMap.Set("sup", "brah")

mappedList := theMap.Map(func(value string, key string, i int) string {
	return key + "_" + value
})

fmt.Println(mappedList.String()) // -> [hey_dude sup_brah]
```

## `.Reduce` over key/value pairs

```
import "github.com/neurocollective/go_chainable/maps"
theMap := maps.NewEmpty[string, string, string]()

theMap.Set("hey", "dude")
theMap.Set("sup", "brah")

initial := "When I meet someone new, I always say: "
message := theMap.Reduce(func(accumulator string, value string, key string, i int) string {
	return accumulator + key + " " + value + " "
}, initial)
fmt.Println(message) // -> "When I meet someone new, I always say: hey dude sup brah"
```

## Current Test Coverage

```
ok  	github.com/neurocollective/go_chainable/lists	0.001s	coverage: 80.4% of statements
ok  	github.com/neurocollective/go_chainable/maps	0.001s	coverage: 97.4% of statements
```

## Methods

### List Methods

`.Raw()`

Returns a raw underlying `[]T`

`.RawPointer()`

Returns the raw Pointer to the underlying `[]T`

`.Map(mapper func(value T, index int) T) *List[T, R]`

Calls the `mapper` function for each element, passing in `value` and `int`. The underlying slice is changed to a new slice, with each element being the returned value from `mapper`.

`.MapFull(mapper func(value T, index int, array *[]T) T) *List[T, R]`

Same as `.Map` but the `mapper` function takes a pointer to the underlying slice as an additional argument.

`.Reduce(reducer func(accumulator R, value T, index int) R, initial R) R)`

Calls the `reducer` function for each element, passing in `accumulator` `value` and `int`. On the first call to `reducer` the `accumulator` value is the `initial` value. But each subsequent call receives an `accumulator` that is the returned `R` value from the previous call to `reducer`.

`.ReduceFull(reducer func(accumulator R, value T, index int, array *[]T) R, initial R) R`

Same as `.Reduce` but the `reducer` function takes a pointer to the underlying slice as an additional argument.

`.ForEach(operation func(element T, index int) T) *List[T, R]`

Calls the `operation` on each element in the slice, returning nothing from each invocation. The unmodified `*List[T, R]` is returned.

`.ForEachFull(operation func(element T, index int, array *[]T)) *List[T, R]`

Same as `.ForEach` but the `operation` function takes a pointer to the underlying slice as an additional argument.

`Filter(filterFunc func(element T, index int) bool) *List[T, R]`

Calls `filterFunc` for each element in the slice. This populates a new slice, which only receives the `T` at that index if `filterFunc` return `true`. The `List[T, R]` then points to the new, filtered array.

`.FilterFull(filterFunc func(element T, index int, array *[]T) bool) *List[T, R]`

Same as `.ForEach` but the `filterFunc` function takes a pointer to the underlying slice as an additional argument.

`.Append(addition *[]T) *List[T, R]`

Adds the values from `[]T` to the `List[T, R]`.

`.Add(value T) *List[T, R]`

Adds the value `T` to the `List[T, R]`.

`.Get(index int) (error, T)`

Returns the value `T` of the element at `index`. Returns an error if the underlying slice is a `nil` pointer or requested `index` would be out of bounds.

`.SetCap(capacity int) *List[T, R]`

Set the capacity of the underlying slice.

`.IncrementCap(capacity int) *List[T, R]`

Increase the capacity of the underlying slice by the `int` value `capacity`.

`.Find(finder func(element T, index int) bool) (error, *T)`

Calls `finder` on every element in the slice. As soon as `finder` returns `true`, the `T` value at that index is retured. If `finder`never returns `true` after traversing all values, an `error` is returned.

`.String() string`

Returns a `string` representation of the underlying slice.

`.IndexOf(matcher func(element T) bool) (error, int)`

Calls `matcher` on each element in the array. For the first element from which `matcher` returns `true`, the `int` index of that element will be returned.

`.Size() (error, int)`

Returns the `int` length of the underlying slice.

`.Cap() (error, int)`

Returns the `int` capacity of the underlying slice.

`.IsEmpty() (error, bool)`

Returns `true` if the underlying slice has a length of 0. Returns an error if it is a `nil` pointer.

`.Last() (error, T)`

Returns the value `T` of the element at the last index. Returns an error if the underlying slice is a `nil` pointer or list is empty.

`.First() (error, T)`

Returns the value `T` of the element at index 0. Returns an error if the underlying slice is a `nil` pointer or list is empty.

### List Functions

`lists.New[T any, R any](array []T) *List[T, R]`

Build a `*List[T, R]` from an array or slice.

`lists.NewEmpty[T any, R any]() *List[T, R]`

Build an empty `*List[T, R]` from an array or slice.

`ResultTypeSwap[T any, OldR any, NewR any] (list *List[T, OldR]) *List[T, NewR]`

Get a new `List[T, NewR]`, in the case where result type was incorrect or a new result type if needed.

### Map Methods

`.Map(mapper func(value V, key K, index int) R) *lists.List[R, R]`

Call `mapper` for each key value pair, returning an `R` value which is appended into the returned `*lists.List[R, R]`.

`.Reduce(reducer func(accumulator R, value V, key K, index int) R, initial R) R`

Calls the `reducer` function for each key-value pair, passing in `accumulator` `value` `key` and `index` (`index` corresonds to the order added to the `Map`). On the first call to `reducer` the `accumulator` value is the `initial` value. But each subsequent call receives an `accumulator` that is the returned `R` value from the previous call to `reducer`.

`.Set(key K, value V) *Map[K, V, R]`

Add a key-value pair to the `Map`. Key will be stored by order added.

`.Get(key K) (V, bool)`

Returns the `V` at `key`. Second value is `false` if the key was not found, and a zero-value was returned.

`.Keys() *lists.List[K, R]`

Returns a `*lists.List[K, R]` of all the keys `K` added to the `Map`, in order of addition.

`.Values() *lists.List[V, R]`

Returns a `*lists.List[K, R]` of all the values `V` added to the `Map`, in order of addition.

`.String() string`

Returns a string representation of the map - order not yet guaranteed.

### Map Functions

`maps.New[K comparable, V comparable, R any](nativeMap map[K]V) *Map[K, V, R]`

Get a new `*Map[K, V, R]` from a `map[K]V`

`maps.NewEmpty[K comparable, V comparable, R any]() *Map[K, V, R]`

Get a new empty `*Map[K, V, R]`
