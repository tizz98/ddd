package ddd

import "time"

type TransactionType int

const (
	TransactionTypeDeposit TransactionType = iota
	TransactionTypeWithdrawal
	TransactionTypeTransfer
)

type LedgerEntry struct {
	AccountNumber uint
	Time          time.Time
	Amount        int
	TransactionID uint
}

type AccountEntity struct {
	AccountNumber uint
	CreationTime  time.Time
	Balance       int
}

type AccountAggregate struct {
	AccountEntity
	LedgerEntries []LedgerEntry
}

type TransactionEntity struct {
	TransactionID uint
	Type          TransactionType

	FromAccountNumber *uint
	ToAccountNumber   *uint
}

type BankingRepository interface {
	// CreateAccount returns the number of the newly created account
	CreateAccount() (accountID uint)
	// RecordLedgerEntry creates a new LedgerEntry based on the given parameters
	RecordLedgerEntry(accountID uint, amount int, transactionID uint)
	// RecordDepositTransaction creates a new transaction of type "Deposit" and returns the transaction id
	RecordDepositTransaction(toAccountNumber uint) (transactionID uint)
	// RecordWithdrawalTransaction creates a new transaction of type "Withdrawal" and returns the transaction id
	RecordWithdrawalTransaction(fromAccountNumber uint) (transactionID uint)
	// RecordTransferTransaction creates a new transaction of type "Transfer" and returns the transaction id
	RecordTransferTransaction(fromAccountNumber uint, toAccountNumber uint) (transactionID uint)

	// GetAccount retrieves an AccountAggregate by the given accountNumber (aka Account ID)
	GetAccount(accountNumber uint) AccountAggregate
	// GetTransaction retrieves a TransactionEntity by the given transactionID
	GetTransaction(transactionID uint) TransactionEntity
}
