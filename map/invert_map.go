package main

import "fmt"

func main() {
	var (
		barVal = map[string]int{"alpha": 34, "bravo": 56, "charlie": 23,
			"delta": 87, "echo": 56, "foxtrot": 12,
			"golf": 34, "hotel": 16, "indio": 87,
			"juliet": 65, "kili": 43, "lima": 98}
	)
	
	invMap := make(map[int]string, len(barVal))
	
	for k, v := range barVal {
		invMap[v] = k
	}
	fmt.Println(invMap)
	for k, v := range invMap {
		fmt.Printf("Key: %v, Value: %v \n ", k, v)
	}
}