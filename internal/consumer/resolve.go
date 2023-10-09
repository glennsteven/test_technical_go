package consumer

import (
	"technical_test_go/technical_test_go/internal/appctx"
	"technical_test_go/technical_test_go/internal/repositories"
)

type consumer struct {
	co  repositories.ConsumerDB
	cfg *appctx.Config
}

func NewConsumerResolve(
	co repositories.ConsumerDB,
	cfg *appctx.Config,
) Resolve {
	return &consumer{
		co:  co,
		cfg: cfg,
	}
}
