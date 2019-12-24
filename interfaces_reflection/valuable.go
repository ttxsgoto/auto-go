package main

import "fmt"

type stockPosition struct {
	ticker     string
	sharePrice float32
	count      float32
}

func (self stockPosition) getValue() float32 {
	return self.sharePrice * self.count
}

type car struct {
	make  string
	model string
	price float32
}

func (self car) getValue() float32 {
	return self.price
	
}

type valuable interface {
	getValue() float32
}

func showValue(asset valuable) {
	fmt.Printf("Value of the asset is %f\n", asset.getValue())
	
}

func main() {
	var o valuable = stockPosition{"GOOG", 577.20, 4}
	showValue(o)
	var o1 = car{"BMW", "M3", 66500}
	showValue(o1)
}

