package main

import "testing"

func Test(t *testing.T) {
	slice := []int{3, 6, -2, -5, 7, 3}
	expectedReult := 21
	res := adjacentElementsProduct(slice)
	if res != expectedReult {
		t.Errorf("\nInput: %v\nExpected %d, but received %d", slice, expectedReult, res)
	}
}
