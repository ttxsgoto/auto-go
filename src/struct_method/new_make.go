package main

// slice, map channel 使用 make方法初始化
// struct 使用new方法初始化


type Foo map[string]string
type Bar struct {
	firstname string
	lastname  string
}

func main() {
	y := new(Bar)
	(*y).firstname = "ttxs"
	(*y).lastname = "goto"
	
	z := make(Foo)
	z["firstname"] = "abc"
	z["lastname"] = "bcd"
	
}


