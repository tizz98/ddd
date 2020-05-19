package inmem

import (
	"ddd"
	"sort"
	"time"
)

type BankingRepository struct {
	accountIDSeq           uint
	transactionEntityIDSeq uint

	accountsByID             map[uint]*ddd.AccountEntity
	ledgerEntriesByAccountId map[uint][]*ddd.LedgerEntry
	transactionsByID         map[uint]*ddd.TransactionEntity
}

func NewBankingRepository() ddd.BankingRepository {
	return &BankingRepository{
		accountIDSeq:           0,
		transactionEntityIDSeq: 0,

		accountsByID:             map[uint]*ddd.AccountEntity{},
		ledgerEntriesByAccountId: map[uint][]*ddd.LedgerEntry{},
		transactionsByID:         map[uint]*ddd.TransactionEntity{},
	}
}

func (b *BankingRepository) CreateAccount() (accountID uint) {
	b.accountIDSeq++

	b.accountsByID[b.accountIDSeq] = &ddd.AccountEntity{
		AccountNumber: b.accountIDSeq,
		CreationTime:  time.Now(),
		Balance:       0,
	}
	b.ledgerEntriesByAccountId[b.accountIDSeq] = []*ddd.LedgerEntry{}
	return b.accountIDSeq
}

func (b *BankingRepository) RecordLedgerEntry(accountID uint, amount int, transactionID uint) {
	entry := &ddd.LedgerEntry{
		AccountNumber: accountID,
		Time:          time.Now(),
		Amount:        amount,
		TransactionID: transactionID,
	}
	b.accountsByID[accountID].Balance += amount
	b.ledgerEntriesByAccountId[accountID] = append(b.ledgerEntriesByAccountId[accountID], entry)
}

func (b *BankingRepository) RecordDepositTransaction(toAccountNumber uint) (transactionID uint) {
	b.transactionEntityIDSeq++

	entity := &ddd.TransactionEntity{
		TransactionID:   b.transactionEntityIDSeq,
		Type:            ddd.TransactionTypeDeposit,
		ToAccountNumber: &toAccountNumber,
	}
	b.transactionsByID[b.transactionEntityIDSeq] = entity

	return b.transactionEntityIDSeq
}

func (b *BankingRepository) RecordWithdrawalTransaction(fromAccountNumber uint) (transactionID uint) {
	b.transactionEntityIDSeq++

	entity := &ddd.TransactionEntity{
		TransactionID:     b.transactionEntityIDSeq,
		Type:              ddd.TransactionTypeWithdrawal,
		FromAccountNumber: &fromAccountNumber,
	}
	b.transactionsByID[b.transactionEntityIDSeq] = entity

	return b.transactionEntityIDSeq
}

func (b *BankingRepository) RecordTransferTransaction(fromAccountNumber uint, toAccountNumber uint) (transactionID uint) {
	b.transactionEntityIDSeq++

	entity := &ddd.TransactionEntity{
		TransactionID:     b.transactionEntityIDSeq,
		Type:              ddd.TransactionTypeTransfer,
		ToAccountNumber:   &toAccountNumber,
		FromAccountNumber: &fromAccountNumber,
	}
	b.transactionsByID[b.transactionEntityIDSeq] = entity

	return b.transactionEntityIDSeq
}

func (b *BankingRepository) GetAccount(accountNumber uint) (account ddd.AccountAggregate) {
	account.AccountEntity = *b.accountsByID[accountNumber]

	for _, entry := range b.ledgerEntriesByAccountId[accountNumber] {
		account.LedgerEntries = append(account.LedgerEntries, *entry)
	}

	// Sort ascending
	sort.SliceStable(account.LedgerEntries, func(i, j int) bool {
		return account.LedgerEntries[i].Time.Before(account.LedgerEntries[j].Time)
	})

	return
}

func (b *BankingRepository) GetTransaction(transactionID uint) ddd.TransactionEntity {
	return *b.transactionsByID[transactionID]
}

var _ ddd.BankingRepository = (*BankingRepository)(nil)
