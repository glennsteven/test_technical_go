package transactions

import (
	"context"
	"fmt"
	"log"
	"technical_test_go/technical_test_go/internal/appctx"
	"technical_test_go/technical_test_go/internal/consts"
	"technical_test_go/technical_test_go/internal/entity"
	"technical_test_go/technical_test_go/internal/presentations"
)

func (a *addTransaction) ResolverManagementTransaction(ctx context.Context, iDConsumer string) appctx.Response {
	findConsumer, err := a.consumer.FindOne(ctx, entity.Consumers{IDConsumer: iDConsumer})
	if err != nil {
		log.Printf("find data consumer got error %v", err)
		return *appctx.NewResponse().WithCode(consts.CodeInternalServerError)
	}

	if findConsumer == nil {
		return *appctx.NewResponse().WithCode(consts.CodeNotFound).WithData("Data customer not found")
	}

	findTransactionConsumer, err := a.transactionCo.FindOne(ctx, entity.Transactions{
		IDConsumer: findConsumer.IDConsumer,
	})

	if err != nil {
		log.Printf("find data transaction by %v got error %v", findConsumer.IDConsumer, err)
		return *appctx.NewResponse().WithCode(consts.CodeInternalServerError)
	}

	if findTransactionConsumer == nil {
		return *appctx.NewResponse().WithCode(consts.CodeNotFound).WithMessage("Customer does not have any transactions.")
	}

	response := presentations.ResponseManagementTransaction{
		Name:           findConsumer.FullName,
		ContractNumber: findTransactionConsumer.ContractNumber,
		Link: presentations.Link{
			ImageIdentity: findConsumer.ImageIdentity,
			ImageSelfie:   findConsumer.ImageIdentity,
		},
		Transaction: presentations.TransactionsData{
			OTR:               findTransactionConsumer.Otr,
			InstallmentAmount: int(findTransactionConsumer.InstallmentAmount),
			TotalInterest:     int(findTransactionConsumer.TotalInterest),
			AssetName:         findTransactionConsumer.AssetName,
			TransactionDate:   findTransactionConsumer.TransactionDate.Format(consts.LayoutDateTimeFormat),
		},
	}

	log.Printf("success get data transaction from user %s", response.Name)

	msg := fmt.Sprintf("The following is a list of %s transactions", response.Name)
	return *appctx.NewResponse().WithCode(consts.CodeSuccess).WithData(response).WithMessage(msg)
}
