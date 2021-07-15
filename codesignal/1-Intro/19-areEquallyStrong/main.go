package main

import "fmt"

func areEquallyStrong(yourLeft int, yourRight int, friendsLeft int, friendsRight int) bool {
    return (yourLeft == friendsLeft && yourRight == friendsRight) || (yourLeft == friendsRight && yourRight == friendsLeft)
}

// Challenge link
// https://app.codesignal.com/arcade/intro/level-5/g6dc9KJyxmFjB98dL
func main() {
	fmt.Println(areEquallyStrong(10, 15, 15, 10)) //Should print true
}
