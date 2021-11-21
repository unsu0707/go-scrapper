package main

import (
	"fmt"
	"log"

	"github.com/unsu0707/learn-golang/banking"
)

func main() {
	account := banking.NewAccount("unsu")
	account.Deposit(1000)
	fmt.Println(account)
	err := account.Withdraw(5000)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(account)
}
