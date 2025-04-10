package domain

// Invoice representa uma fatura de transação
type Invoice struct {
	ID        string
	AccountID string
	Amount    float64
	Status    string
}

// AccountRepository define as operações de persistência para Account
type AccountRepository interface {
	Save(account *Account) error
	FindByAPIKey(apiKey string) (*Account, error)
	FindByID(id string) (*Account, error)
	UpdateBalance(account *Account) error
}

// TransactionRepository define as operações de persistência para Invoice
type TransactionRepository interface {
	Save(invoice *Invoice) error
	FindByID(id string) (*Invoice, error)
	FindByAccountID(accountID string) ([]*Invoice, error)
	UpdateStatus(invoice *Invoice) error
}
