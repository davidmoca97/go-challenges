package main

import "testing"

func tests(t *testing.T) {
	cases := [][]int{
		{1905, 20},
		{1700, 70},
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
			t.Errorf("Expected %d, but received %d", c[1], res)
		}
	}
}
