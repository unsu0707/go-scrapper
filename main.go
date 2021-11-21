package main

import (
	"fmt"
)

func superAdd(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	fmt.Println(superAdd(1, 3, 5, 7, 9))
}
