package loan_limit

import (
	"context"
	"github.com/google/uuid"
	"log"
	"technical_test_go/technical_test_go/internal/appctx"
	"technical_test_go/technical_test_go/internal/consts"
	"technical_test_go/technical_test_go/internal/entity"
	"technical_test_go/technical_test_go/internal/presentations"
	"time"
)

func (l *limitAmount) ResolverStoreLimit(ctx context.Context, payload presentations.PayloadAddLimit, iDConsumer string) appctx.Response {
	findConsumer, err := l.consumer.FindOne(ctx, entity.Consumers{
		IDConsumer: iDConsumer,
	})
	if err != nil {
		log.Printf("find consumer by id got error %v", err.Error())
		return *appctx.NewResponse().WithCode(consts.CodeInternalServerError)
	}

	if findConsumer == nil {
		log.Printf("data consumer not found")
		return *appctx.NewResponse().WithCode(consts.CodeNotFound).WithMessage("data consumer not found")
	}

	storeLimit := entity.ConsumerLimits{
		IDLimit:     uuid.New().String(),
		IDConsumer:  findConsumer.IDConsumer,
		Tenor:       payload.Tenor,
		LimitAmount: payload.Amount,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	saveLimitConsumer, err := l.limit.Store(ctx, storeLimit)
	if err != nil {
		log.Printf("store data limit consumer got error %v", err.Error())
		return *appctx.NewResponse().WithCode(consts.CodeInternalServerError)
	}

	return *appctx.NewResponse().
		WithCode(consts.CodeCreated).
		WithData(saveLimitConsumer).
		WithMessage("Loan limit has been successfully added.")
}
