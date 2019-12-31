package main

import (
	"fmt"
	"sort"
)

func main() {
	var (
		barVal = map[string]int{"alpha": 34, "bravo": 56, "charlie": 23,
			"delta": 87, "echo": 56, "foxtrot": 12,
			"golf": 34, "hotel": 16, "indio": 87,
			"juliet": 65, "kili": 43, "lima": 98}
	)
	
	for k, v := range barVal {
		fmt.Printf("Key: %v, Value: %v \n ", k, v)
	}
	
	keys := make([]string, len(barVal))
	
	i := 0
	for k, _ := range barVal {
		keys[i] = k
		i++
	}
	
	sort.Strings(keys)
	fmt.Println(keys)
	fmt.Println("Sorted")
	for _, k := range keys {
		fmt.Printf("Key: %v, Value: %v \n ", k, barVal[k])
	}
}
