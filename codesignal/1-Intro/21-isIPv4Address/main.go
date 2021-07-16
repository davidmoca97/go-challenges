package main

import (
	"fmt"
    "net"
    "regexp"
)

func isIPv4Address(inputString string) bool {
    isValid := regexp.MustCompile(`^([^0]\d{0,2}|0)(\.([^0]\d{0,2}|0)){3}$`)
    if !isValid.MatchString(inputString) {
        return false
    }
    return net.ParseIP(inputString) != nil
}

func main() {
	fmt.Println(isIPv4Address("192.168.1.254")) //Should print true
}
