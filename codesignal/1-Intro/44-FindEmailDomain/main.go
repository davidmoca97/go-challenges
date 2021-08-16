package main

import (
	"fmt"
	"regexp"
)

func findEmailDomain(address string) string {
	matches := regexp.MustCompile(`^.+@(.+)$`).FindAllStringSubmatch(address, -1)
	if len(matches) < 0 {
		return ""
	}
	return matches[0][1]
}

// Challenge:
// https://app.codesignal.com/arcade/intro/level-10/TXFLopNcCNbJLQivP
func main() {
	fmt.Println(findEmailDomain("prettyandsimple@example.com")) //Should print @example.com
}
