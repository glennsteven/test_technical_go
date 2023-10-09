package transactions

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"log"
	"sync"
	"technical_test_go/technical_test_go/internal/appctx"
	"technical_test_go/technical_test_go/internal/consts"
	"technical_test_go/technical_test_go/internal/entity"
	"technical_test_go/technical_test_go/internal/presentations"
	"time"
)

func (a *addTransaction) ResolverStoreTransaction(ctx context.Context, payload presentations.PayloadAddTransaction) appctx.Response {
	var (
		limitInstallment int64
		limitOTR         float64
		mtx              sync.Mutex
	)

	findLimit, err := a.limit.FindOne(ctx, entity.ConsumerLimits{
		IDConsumer: payload.IDConsumer,
	})

	if err != nil {
		log.Printf("find data limit consumer got error %v", err)
		return *appctx.NewResponse().WithCode(consts.CodeInternalServerError)
	}

	limitInstallment = findLimit.Tenor
	limitOTR = findLimit.LimitAmount

	findConsumer, err := a.consumer.FindOne(ctx, entity.Consumers{IDConsumer: payload.IDConsumer})
	if err != nil {
		log.Printf("find data consumer got error %v", err.Error())
		return *appctx.NewResponse().WithCode(consts.CodeInternalServerError)
	}

	mtx.Lock()
	defer mtx.Unlock()
	isLimitTransaction := limitInstallment == int64(payload.InstallmentAmount)
	isLimitOTR := limitOTR >= payload.OTR

	if !isLimitTransaction {
		msg := fmt.Sprintf("the installment limit is %v", limitInstallment)
		return *appctx.NewResponse().WithCode(consts.CodeUnprocessableEntity).WithMessage(msg)
	}

	if !isLimitOTR {
		msg := fmt.Sprintf("the otr limit is %.2f", limitOTR)
		return *appctx.NewResponse().WithCode(consts.CodeUnprocessableEntity).WithMessage(msg)
	}

	dataTransaction := entity.Transactions{
		IDTransaction:     uuid.New().String(),
		IDConsumer:        findConsumer.IDConsumer,
		ContractNumber:    payload.ContractNumber,
		Otr:               payload.OTR,
		FeeAdmin:          payload.FeeAdmin,
		InstallmentAmount: int64(payload.InstallmentAmount),
		TotalInterest:     float64(payload.TotalInterest),
		AssetName:         payload.AssetName,
		TransactionDate:   time.Now(),
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}

	saveTransaction, err := a.transactionCo.Store(ctx, dataTransaction)
	if err != nil {
		log.Printf("save data transaction consumer got error %v", err.Error())
		return *appctx.NewResponse().WithCode(consts.CodeInternalServerError)
	}

	updateLimit := limitOTR - payload.OTR
	if updateLimit == 0 {
		updateLimit = 0.01
	}
	err = a.limit.Update(ctx, entity.ConsumerLimits{
		LimitAmount: updateLimit,
	}, entity.ConsumerLimits{
		IDConsumer: findConsumer.IDConsumer,
	})

	if err != nil {
		log.Printf("update limit consumer got error %v", err)
		return *appctx.NewResponse().WithCode(consts.CodeInternalServerError)
	}

	return *appctx.NewResponse().
		WithCode(consts.CodeCreated).
		WithData(saveTransaction).
		WithMessage("Transaction created successfully.")
}
