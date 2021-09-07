package main

import "fmt"

func fileNaming(names []string) []string {
	occurrences := make(map[string]int)
	result := []string{}
	for _, fileName := range names {
		if occurrences[fileName] > 0 {
			i := 1
			for occurrences[fmt.Sprintf("%s(%d)", fileName, i)] > 0 {
				i++
			}
			fileName = fmt.Sprintf("%s(%d)", fileName, i)
		}
		result = append(result, fileName)
		occurrences[fileName]++
	}
	return result
}

// Challenge:
// https://app.codesignal.com/arcade/intro/level-12/sqZ9qDTFHXBNrQeLC
func main() {
	res := fileNaming([]string{"doc", "doc", "image", "doc(1)", "doc"})
	fmt.Println(res) // Should print [doc, doc(1), image, doc(1)(1), doc(2)]
}
