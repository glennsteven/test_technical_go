package transactions

import (
	"technical_test_go/technical_test_go/internal/repositories"
)

type addTransaction struct {
	consumer      repositories.ConsumerDB
	transactionCo repositories.TransactionDB
	limit         repositories.LimitAmountDB
}

func NewTransactions(
	consumer repositories.ConsumerDB,
	transactionCo repositories.TransactionDB,
	limit repositories.LimitAmountDB,
) Resolve {
	return &addTransaction{
		consumer:      consumer,
		transactionCo: transactionCo,
		limit:         limit,
	}
}
