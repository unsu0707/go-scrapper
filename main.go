package main

import (
	"fmt"
)

func main() {
	names := []string{"aaa", "bbb", "ccc"}
	names = append(names, "ddd")
	names[3] = "ddd"
	fmt.Println(names)
}
