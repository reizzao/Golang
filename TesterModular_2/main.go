package main

import "fmt"

type Foo1 struct {
	Bar1 Bar1
}
type Bar1 struct {
	C1 int
	C2 int
}
type Bar2 struct {
	B1 int
	B2 int
}

func main() {
	res := Foo1{
		Bar1: Bar1{C1: 10, C2: 20},
	}
	fmt.Println(res)
}
