package main

import "testing"

type Case struct {
	word         string
	isPalindrome bool
}

func Test(t *testing.T) {
	cases := []Case{
		{"anitalavalatina", true},
		{"lol", true},
		{"tattarrattat", true},
		{"nosoypalindroma", false},
		{"", false},
	}

	for _, c := range cases {
		res := checkPalindromeWord(c.word)
		if res != c.isPalindrome {
			t.Errorf("\nInput: %s\nExpected %t, but received %t", c.word, c.isPalindrome, res)
		}
	}
}
