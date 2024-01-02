package tw

import (
	"testing"
)

func TestTwMerge(t *testing.T) {
	name := "Test Tailwind Merge"

	testString := TwString{Input: "text-red-500"}
	testObject := TwObject{Input: map[string]bool{"text-red-500": true, "text-blue-500": true, "text-green-500": false}}
	testSlice := TwSlice{Input: []string{"text-red-500", "text-blue-500", "text-purple-500"}}

	result := Merge(&testString, &testObject, &testSlice)

	if result != "text-red-500 text-blue-500 text-purple-500" {
		t.Errorf("%s: %s", name, "Result is not correct")
	}
}
