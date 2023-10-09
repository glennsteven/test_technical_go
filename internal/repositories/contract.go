package repositories

import (
	"context"
	"technical_test_go/technical_test_go/internal/entity"
)

type ConsumerDB interface {
	Store(ctx context.Context, param entity.Consumers) (entity.Consumers, error)
	FindOne(ctx context.Context, where entity.Consumers) (*entity.Consumers, error)
	Update(ctx context.Context, param entity.Consumers, where entity.Consumers) error
	Find(ctx context.Context) ([]entity.Consumers, error)
}

type LimitAmountDB interface {
	Store(ctx context.Context, param entity.ConsumerLimits) (entity.ConsumerLimits, error)
	Update(ctx context.Context, param entity.ConsumerLimits, where entity.ConsumerLimits) error
	FindOne(ctx context.Context, where entity.ConsumerLimits) (*entity.ConsumerLimits, error)
	Find(ctx context.Context) ([]entity.ConsumerLimits, error)
}

type TransactionDB interface {
	Store(ctx context.Context, param entity.Transactions) (entity.Transactions, error)
	Update(ctx context.Context, param entity.Transactions, where entity.Transactions) error
	FindOne(ctx context.Context, where entity.Transactions) (*entity.Transactions, error)
	Find(ctx context.Context) ([]entity.Transactions, error)
}
