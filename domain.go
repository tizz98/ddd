package ddd

type Domain struct {
	BankingRepository BankingRepository
}

type BankingService interface {
	GetAccount(domain *Domain, accountNumber uint) AccountAggregate
	OpenAccount(domain *Domain) (newAccount AccountAggregate)
	Deposit(domain *Domain, toAccountNumber uint, amount int) (toAccount AccountAggregate)
	Withdraw(domain *Domain, fromAccountNumber uint, amount int) (fromAccount AccountAggregate)
	Transfer(domain *Domain, fromAccountNumber uint, toAccountNumber uint, amount int) struct {
		ToAccount   AccountAggregate
		FromAccount AccountAggregate
	}
}

type BankingServiceImpl struct{}

func (b BankingServiceImpl) GetAccount(domain *Domain, accountNumber uint) AccountAggregate {
	return domain.BankingRepository.GetAccount(accountNumber)
}

func (b BankingServiceImpl) OpenAccount(domain *Domain) (newAccount AccountAggregate) {
	accountID := domain.BankingRepository.CreateAccount()
	return domain.BankingRepository.GetAccount(accountID)
}

func (b BankingServiceImpl) Deposit(domain *Domain, toAccountNumber uint, amount int) (toAccount AccountAggregate) {
	transactionID := domain.BankingRepository.RecordDepositTransaction(toAccountNumber)
	domain.BankingRepository.RecordLedgerEntry(toAccountNumber, amount, transactionID)
	return domain.BankingRepository.GetAccount(toAccountNumber)
}

func (b BankingServiceImpl) Withdraw(domain *Domain, fromAccountNumber uint, amount int) (fromAccount AccountAggregate) {
	transactionID := domain.BankingRepository.RecordWithdrawalTransaction(fromAccountNumber)
	domain.BankingRepository.RecordLedgerEntry(fromAccountNumber, -amount, transactionID)
	return domain.BankingRepository.GetAccount(fromAccountNumber)
}

func (b BankingServiceImpl) Transfer(domain *Domain, fromAccountNumber uint, toAccountNumber uint, amount int) struct {
	ToAccount   AccountAggregate
	FromAccount AccountAggregate
} {
	transactionId := domain.BankingRepository.RecordTransferTransaction(fromAccountNumber, toAccountNumber)
	domain.BankingRepository.RecordLedgerEntry(fromAccountNumber, -amount, transactionId)
	domain.BankingRepository.RecordLedgerEntry(toAccountNumber, amount, transactionId)

	return struct {
		ToAccount   AccountAggregate
		FromAccount AccountAggregate
	}{
		ToAccount:   domain.BankingRepository.GetAccount(toAccountNumber),
		FromAccount: domain.BankingRepository.GetAccount(fromAccountNumber),
	}
}

var _ BankingService = (*BankingServiceImpl)(nil)
