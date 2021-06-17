package main

import "testing"

func Test(t *testing.T) {
	cases := [][]int{
		{1905, 20},
		{1700, 17},
		{1988, 20},
		{2000, 20},
		{2001, 21},
		{200, 2},
		{374, 4},
		{45, 1},
		{8, 1},
	}

	for _, c := range cases {
		res := centuryFromYear(c[0])
		if res != c[1] {
			t.Errorf("\nInput:  %d\nExpected %d, but received %d", c[0], c[1], res)
		}
	}
}
