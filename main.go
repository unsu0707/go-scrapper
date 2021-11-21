package main

import (
	"fmt"

	"github.com/unsu0707/learn-golang/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	word := "hello"
	definition := "Greeting"
	err := dictionary.Add(word, definition)
	if err != nil {
		fmt.Println(err)
	}
	result, _ := dictionary.Search(word)
	fmt.Println("found", word, "definition:", result)
	err2 := dictionary.Add(word, definition)
	if err2 != nil {
		fmt.Println(err2)
	}
}
