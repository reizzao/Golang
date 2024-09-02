package main

import "fmt"

func use(i int) int {
	tom := i
	var ch int

	// definindo valor d euma var conforme a entrada de dado

	switch {
	case tom == 1:
		ch = 11
	case tom == 2:
		ch = 22
	default:
		ch = 00
	}
	return ch

}

func main() {
	fmt.Println(use(1))
}
