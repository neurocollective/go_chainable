package maps

import (
	"testing"
)

func TestMapDotMap(t *testing.T) {

	theMap := NewEmpty[string, string, string]()

	theMap.Add("hey", "dude")
	theMap.Add("sup", "brah")

	mapped := theMap.Map(func(value string, key string) string {
		return key + "_" + value
	})

	if false {
		t.Error("ruh roh")
	}

	if error, value := mapped.Get(0); error != nil || value != "hey_dude" {
		t.Error("oh noes")
	}
}