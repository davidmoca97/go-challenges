package main

import "testing"

func Test(t *testing.T) {
	cases := [][]int{
		{1, 1, 2},
		{-1, -1, -2},
		{0, -1, -1},
		{999999, 1, 1000000},
		{-1000, 1000, 0},
	}

	for _, c := range cases {
		res := add(c[0], c[1])
		if res != c[2] {
			t.Errorf("\nInput: %d, %d\nExpected %d, but received %d", c[0], c[1], c[2], res)
		}
	}
}
