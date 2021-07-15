package main

import "fmt"

func areSimilar(a []int, b []int) bool {
    swap := false
    for i:=0; i<len(a); i++ {
        if a[i] == b[i] {
            continue
        }
        if swap {
            return false
        }
        for j:=i+1; j<len(a); j++ {
            if b[j] == a[i] && a[j] == b[i] {
                b[i], b[j] = b[j], b[i]
                swap = true
                break
            }
        }
        if !swap {
            return false
        }
    }
    return true
}

// Challenge link
// https://app.codesignal.com/arcade/intro/level-4/xYXfzQmnhBvEKJwXP

func main() {
	a := []int{1, 2, 3}
	b := []int{2, 1, 3}
	fmt.Println(areSimilar(a, b)) // Should return true
}