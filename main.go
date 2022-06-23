package main

import (
	"fmt"
	"log"

	"github.com/nagneo/go-jobs/banking"
)

func main() {
	account := banking.NewAccount("nico")
	account.Deposit(1000)
	err := account.Withdraw(20)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(account)
}
