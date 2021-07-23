package main

import (
	"fmt"
	"regexp"
)

func variableName(name string) bool {
	match := regexp.MustCompile("(?i)^[a-z_][a-z0-9_]*$").MatchString(name)
	if match {
		return true
	}
	return false
}

// Chanllenge:
// https://app.codesignal.com/arcade/intro/level-6/6Wv4WsrsMJ8Y2Fwno
func main() {
	str := "var_1__Int"
	fmt.Println(variableName(str)) //should print true
}
