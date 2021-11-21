package main

import (
	"fmt"
)

func main() {
	a := 2
	b := &a
	a = 50
	*b = 100
	fmt.Println(a, *b)
}
