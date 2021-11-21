package main

import (
	"fmt"
)

type user struct {
	name         string
	age          int
	favoriteFood []string
}

func main() {
	user1 := user{"unsu", 20, []string{"yakiniku", "ramen"}}
	fmt.Println(user1)
}
