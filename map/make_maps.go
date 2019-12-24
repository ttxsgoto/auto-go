package main

import "fmt"

func main() {
	var mapLit map[string]int
	var mapAssigned map[string]int
	
	mapLit = map[string]int{"one":1, "two":2}
	mapCreate := make(map[string]float64)
	mapAssigned = mapLit
	
	mapCreate["key1"] = 4.5
	mapCreate["key2"] = 3.14
	mapAssigned["two"] = 3
	
	fmt.Printf("Map literal at \"one\" is: %d\n", mapLit["one"])
	fmt.Printf("Map created at \"key2\" is: %f\n", mapCreate["key2"])
	fmt.Printf("Map assigned at \"two\" is: %d\n", mapLit["two"])
	fmt.Printf("Map literal at \"ten\" is: %d\n", mapLit["ten"])
	
	mf := map[int]func() int{
		1: func() int { return 10 },
		2: func() int { return 20 },
		5: func() int { return 50 },
	}
	fmt.Println(mf)
	
}
