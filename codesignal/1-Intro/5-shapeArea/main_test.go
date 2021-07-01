package main

import "testing"

func Test(t *testing.T) {
	cases := [][2]int{
		{1, 1},
		{2, 5},
		{3, 13},
		{4, 25},
	}
	for _, testCase := range cases {
		parameter, expected := testCase[0], testCase[1]
		res := shapeArea(parameter)
		if res != expected {
			t.Errorf("\nInput: %v\nExpected %d, but received %d", parameter, expected, res)
		}
	}
}
