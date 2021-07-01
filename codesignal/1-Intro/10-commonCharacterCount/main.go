package main

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func commonCharacterCount(s1 string, s2 string) int {
	s1Map := make(map[rune]int)
	s2Map := make(map[rune]int)
	commonChars := 0

	for _, char := range s1 {
		s1Map[char]++
	}
	for _, char := range s2 {
		s2Map[char]++
	}

	for char, count := range s1Map {
		if s2Map[char] > 0 {
			commonChars += min(count, s2Map[char])
		}
	}

	return commonChars
}

func main() {
	s1 := "aabcc"
	s2 := "adcaa"
	// Should return 3
	commonCharacterCount(s1, s2)
}
