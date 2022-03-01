package main

import (
	"fmt"
	"github.com/neurocollective/go_chainable/lists"
)

func main() {

	theMap := maps.NewEmpty[string, string, string]()
	fmt.Println(theMap)
	// theMap.Add("hey", "dude")
}