package main

import (
	"fmt"

	"github.com/unsu0707/learn-golang/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	word := "hello"
	dictionary.Add(word, "First")
	dictionary.Update(word, "Second")
	result, err := dictionary.Search(word)
	fmt.Println(result, err)
	dictionary.Delete(word)
	result2, err := dictionary.Search(word)
	fmt.Println(result2, err)
}
