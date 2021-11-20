package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func main() {
	totalLength, upperName := lenAndUpper("unsu0707")
	fmt.Println(totalLength, upperName)
}
