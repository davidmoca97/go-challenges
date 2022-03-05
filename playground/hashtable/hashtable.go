package main

import "fmt"

type HashTable struct {
	array []string
}

func (ht *HashTable) GetItem(key string) string {
	idx := hashStringToInt(key)
	return ht.array[idx]
}

func (ht *HashTable) SetItem(key string, value string) {
	if ht.array == nil {
		ht.array = make([]string, 10)
	}
	idx := hashStringToInt(key)
	ht.array[idx] = value
}

func (ht HashTable) Sing() {
}

func hashStringToInt(str string) int {
	return 0
}

type Hasher interface {
	GetItem(key string) string
	SetItem(key string, value string)
}

type Lopere interface {
	Sing()
}

func main() {
	var myHashTable *HashTable
	myHashTable.SetItem("David", "Congratulations")
	fmt.Println(myHashTable.GetItem("David"))
}

func Hey(h Hasher) {
	fmt.Println(h)
}
