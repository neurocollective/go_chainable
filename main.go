package main

import (
	"fmt"
	"github.com/neurocollective/go_chainable/lists"
	"github.com/neurocollective/go_chainable/maps"
)

// func mapperTwo[s string](value string, key string, nativeMap *map[string]string) string {
// 	return key + "_" + value
// }

func mapperTwo(value string, key string, nativeMap *map[string]string) string {
	return key + "_" + value
}

func main() {
	fmt.Println("sup")

	arr := []int { 0, 1, 2, 3, 4 }
	var val any
	list := lists.List[int, any]{ &arr, val }

	fmt.Println(*list.Array)

	mapper := func (value int, index int, array *[]int) int {
		return value + 1
	}

	mapped := list.Map(mapper)

	fmt.Println(*mapped)

	nativeMap := map[string]string {
		"hey": "dude",
		"sup": "brah",
	}
	array := []string { "hey", "sup" }
	theMap := maps.Map[string, string, string]{
		&nativeMap,
		&lists.List[string, string]{ &array, "" },
	}

	mappedTwo := theMap.Map(mapperTwo)

	fmt.Println(*mappedTwo)

}