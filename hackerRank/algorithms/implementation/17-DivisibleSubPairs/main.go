package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func divisibleSumPairs(n int32, k int32, ar []int32) int32 {
	var count int32
	for i := 0; i < len(ar)-1; i++ {
		for j := i + 1; j < len(ar); j++ {
			if (ar[i]+ar[j])%k == 0 {
				count++
			}
		}
	}
	return count
}

// Challenge:
// https://www.hackerrank.com/challenges/divisible-sum-pairs/problem
func main() {
	// Should print 3
	fmt.Println(divisibleSumPairs(6, 5, []int32{1, 2, 3, 4, 5, 6}))

	f, err := os.OpenFile("./hackerRank/algorithms/implementation/17-DivisibleSubPairs/file.json", os.O_RDWR|os.O_APPEND, 0775)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	server := mux.NewRouter()
	server.HandleFunc("/{key}", func(w http.ResponseWriter, r *http.Request) {
		search := r.URL.Query().Get("search")
		vars := mux.Vars(r)
		w.Write([]byte("{\"name\":\"" + search + vars["key"] + "\"}"))
	})
	http.ListenAndServe(":9090", server)
}
