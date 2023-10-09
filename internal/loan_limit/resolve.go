package loan_limit

import (
	"technical_test_go/technical_test_go/internal/repositories"
)

type limitAmount struct {
	consumer repositories.ConsumerDB
	limit    repositories.LimitAmountDB
}

func NewLimitAmount(
	consumer repositories.ConsumerDB,
	limit repositories.LimitAmountDB,
) Resolve {
	return &limitAmount{
		consumer: consumer,
		limit:    limit,
	}
}
