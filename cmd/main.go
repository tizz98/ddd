package main

import (
	"ddd"
	"ddd/inmem"
	"fmt"
)

func main() {
	domain := &ddd.Domain{BankingRepository: inmem.NewBankingRepository()}

	var banking ddd.BankingService
	banking = ddd.BankingServiceImpl{}

	account1Number := banking.OpenAccount(domain).AccountNumber
	account2Number := banking.OpenAccount(domain).AccountNumber

	banking.Deposit(domain, account1Number, 100)
	banking.Deposit(domain, account2Number, 50)

	banking.Transfer(domain, account1Number, account2Number, 5)

	account1 := banking.GetAccount(domain, account1Number)
	account2 := banking.GetAccount(domain, account2Number)

	fmt.Println("=== BALANCES ===")
	fmt.Printf("Account 1: %d\n", account1.Balance)
	fmt.Printf("Account 2: %d\n", account2.Balance)
}
