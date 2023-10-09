package transactions

import (
	"context"
	"log"
	"technical_test_go/technical_test_go/internal/appctx"
	"technical_test_go/technical_test_go/internal/consts"
	"technical_test_go/technical_test_go/internal/entity"
	"technical_test_go/technical_test_go/internal/presentations"
	"time"
)

func (a *addTransaction) ResolverUpdateTransaction(ctx context.Context, payload presentations.PayloadAddTransaction, iDTransaction string) appctx.Response {
	err := a.transactionCo.Update(ctx, entity.Transactions{
		Otr:               payload.OTR,
		FeeAdmin:          payload.FeeAdmin,
		InstallmentAmount: int64(payload.InstallmentAmount),
		TotalInterest:     float64(payload.TotalInterest),
		AssetName:         payload.AssetName,
		UpdatedAt:         time.Now(),
	}, entity.Transactions{
		IDTransaction: iDTransaction,
	})
	if err != nil {
		log.Printf("update data transaction consumer got error %v", err.Error())
		return *appctx.NewResponse().WithCode(consts.CodeInternalServerError)
	}

	return *appctx.NewResponse().
		WithCode(consts.CodeSuccess).
		WithMessage("Transaction information updated successfully.")
}
