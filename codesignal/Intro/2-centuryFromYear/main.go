package main

func centuryFromYear(year int) int {
	if year%100 == 0 {
		return year / 100
	}
	return year/100 + 1
}