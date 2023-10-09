package transaction_ucase

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"technical_test_go/technical_test_go/internal/appctx"
	"technical_test_go/technical_test_go/internal/transactions"
	"technical_test_go/technical_test_go/internal/ucase/contract"
	"technical_test_go/technical_test_go/pkg/tracer"
)

type findTransactionUcase struct {
	resolve transactions.Resolve
}

func NewFindTransactionUcase(resolve transactions.Resolve) contract.UseCase {
	return &findTransactionUcase{resolve: resolve}
}

func (f *findTransactionUcase) Serve(dataContext *appctx.Data) appctx.Response {
	var (
		ctx           = tracer.SpanStart(dataContext.Request.Context(), "ucase.FindLimitConsumer")
		params        = mux.Vars(dataContext.Request)
		transactionID = params["transaction_id"]
	)

	result := f.resolve.ResolverGetTransactions(ctx, transactionID)
	switch result.Code {
	case http.StatusOK:
		log.Printf("success get data transaction consumer %v", result.Message)
	default:
		log.Printf("got response %v", result.Message)
	}
	return result
}
