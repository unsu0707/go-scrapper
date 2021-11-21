package main

import (
	"fmt"

	"github.com/unsu0707/learn-golang/banking"
)

func main() {
	account := banking.NewAccount("unsu")
	account.Deposit(1000)
	fmt.Println(account)
	account.Withdraw(5000)
	fmt.Println(account)
}
