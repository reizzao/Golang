package main

import "fmt"

type Foo struct {
	C1 int
	C2 int
}
func bar() (int, int) {
  return 5555, 77777
}



func main() {
	a, _ := bar()

	newfoo := Foo{C1: a, C2: 10}

	fmt.Println(newfoo)
}
