package transactions

import (
	"context"
	"log"
	"technical_test_go/technical_test_go/internal/appctx"
	"technical_test_go/technical_test_go/internal/consts"
	"technical_test_go/technical_test_go/internal/entity"
)

func (a *addTransaction) ResolverGetTransactions(ctx context.Context, iDTransaction string) appctx.Response {
	if iDTransaction == "all" {
		findTransactions, err := a.transactionCo.Find(ctx)
		if err != nil {
			log.Printf("find transactions data got error %v", err)
			return *appctx.NewResponse().WithCode(consts.CodeInternalServerError)
		}

		return *appctx.NewResponse().
			WithCode(consts.CodeSuccess).
			WithData(findTransactions).
			WithMessage("Success get information transactions consumer")
	}

	findTransaction, err := a.transactionCo.FindOne(ctx, entity.Transactions{IDTransaction: iDTransaction})
	if err != nil {
		log.Printf("find transaction data got error %v", err)
		return *appctx.NewResponse().WithCode(consts.CodeSuccess)
	}

	if findTransaction == nil {
		return *appctx.NewResponse().WithCode(consts.CodeNotFound).WithMessage("Information Transaction consumer not found")
	}

	return *appctx.NewResponse().
		WithCode(consts.CodeSuccess).
		WithData(findTransaction).
		WithMessage("Success get information transaction consumer")
}
