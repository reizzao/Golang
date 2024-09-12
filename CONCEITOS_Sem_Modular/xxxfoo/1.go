package main

import "fmt"

func foo() {
	fmt.Println("ola_1")
}

var bar = foo

func main() {
	bar()
}
