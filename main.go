package main

import (
	"fmt"
)

func superAdd(numbers ...int) int {
	for index, number := range numbers {
		fmt.Println(index, number)
	}
	return 0
}

func main() {
	superAdd(1, 3, 5, 7, 9)
}
