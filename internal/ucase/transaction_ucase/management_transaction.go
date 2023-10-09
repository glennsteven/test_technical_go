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

type managementTransactionUcase struct {
	resolve transactions.Resolve
}

func NewManagementTransactionUcase(resolve transactions.Resolve) contract.UseCase {
	return &managementTransactionUcase{resolve: resolve}
}

func (m *managementTransactionUcase) Serve(dataContext *appctx.Data) appctx.Response {
	var (
		ctx        = tracer.SpanStart(dataContext.Request.Context(), "ucase.ManagementTransactnion")
		params     = mux.Vars(dataContext.Request)
		consumerID = params["consumer_id"]
	)

	result := m.resolve.ResolverManagementTransaction(ctx, consumerID)
	switch result.Code {
	case http.StatusOK:
		log.Printf("success get data management transaction consumer %v", result.Message)
	default:
		log.Printf("got responsesssssss %v", result.Message)
	}
	return result
}
