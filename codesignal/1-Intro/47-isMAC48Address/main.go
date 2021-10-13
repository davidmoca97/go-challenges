package main

import (
	"fmt"
	"net"
)

func isMAC48Address(inputString string) bool {
	if _, err := net.ParseMAC(inputString); err != nil {
		return false
	}
	return true
}

// Challenge:
// https://app.codesignal.com/arcade/intro/level-10/HJ2thsvjL25iCvvdm
func main() {
	fmt.Println(isMAC48Address("00-1B-63-84-45-E6"))    // Should print true
	fmt.Println(isMAC48Address("Z1-1B-63-84-45-E6"))    // Should print false
	fmt.Println(isMAC48Address("not a MAC-48 address")) // Should print false
}
