package main

import "testing"

func Test(t *testing.T) {
	slice := []int{6, 2, 3, 8}
	expectedReult := 3
	res := makeArrayConsecutive2(slice)
	if res != expectedReult {
		t.Errorf("\nInput: %v\nExpected %d, but received %d", slice, expectedReult, res)
	}
}
