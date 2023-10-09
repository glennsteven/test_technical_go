package transaction_ucase

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"technical_test_go/technical_test_go/internal/appctx"
	"technical_test_go/technical_test_go/internal/consts"
	"technical_test_go/technical_test_go/internal/presentations"
	"technical_test_go/technical_test_go/internal/transactions"
	"technical_test_go/technical_test_go/internal/ucase/contract"
	"technical_test_go/technical_test_go/pkg/tracer"
)

type addTransactionUcase struct {
	resolve transactions.Resolve
}

func NewTransactionUcase(resolve transactions.Resolve) contract.UseCase {
	return &addTransactionUcase{resolve: resolve}
}

func (a *addTransactionUcase) Serve(dataContext *appctx.Data) appctx.Response {
	var (
		param      presentations.PayloadAddTransaction
		params     = mux.Vars(dataContext.Request)
		consumerID = params["consumer_id"]
		ctx        = tracer.SpanStart(dataContext.Request.Context(), "ucase.AddTransaction")
	)

	err := dataContext.Cast(&param)
	if err != nil {
		log.Printf("error parsing query url %v", err)
		return *appctx.NewResponse().WithMsgKey(consts.RespValidationError).WithCode(consts.CodeUnprocessableEntity)
	}

	param.IDConsumer = consumerID

	result := a.resolve.ResolverStoreTransaction(ctx, param)
	switch result.Code {
	case http.StatusCreated, http.StatusOK:
		log.Printf("successfully created transaction %v", result.Data)
	default:
		log.Printf("got response error %v", result.Message)
	}

	return result
}
