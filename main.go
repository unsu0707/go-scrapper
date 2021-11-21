package main

import (
	"fmt"
)

func main() {
	user := map[string]string{"name": "unsu", "age": "31"}
	for key, value := range user {
		fmt.Println(key, value)
	}
}
