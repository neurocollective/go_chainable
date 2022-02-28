# go_chainable

`NOTE: Requires go1.18`

Go-Chainable is a library using generics to mimic the functional `.map(x -> y).reduce(x -> y, z)` patterns of javascript.

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
array := []string { 1, 2 }
list := lists.New[int, any](array)
doubled := list.Map(func(val string, index int, list *lists.List[int, any]) int {
	return val * 2
})
fmt.Println(doubled) // -> [2 4]
```

### Reduce A List

```
array := []string { 1, 2 }
list := lists.New[int, int](array)
added := list.Reduce(func(accumulator int, val string, index int, list *lists.List[int, int]) int {
	return accumulator + val
}, 0)
fmt.Println(added) // -> 3
```

## Map

(more documentation coming soon)