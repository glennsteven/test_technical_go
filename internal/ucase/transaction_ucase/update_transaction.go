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

type updateTransactionUcase struct {
	resolve transactions.Resolve
}

func NewUpdateTransactionUcase(resolve transactions.Resolve) contract.UseCase {
	return &updateTransactionUcase{resolve: resolve}
}

func (u *updateTransactionUcase) Serve(dataContext *appctx.Data) appctx.Response {
	var (
		param         presentations.PayloadAddTransaction
		ctx           = tracer.SpanStart(dataContext.Request.Context(), "ucase.UpdateTransactionID")
		params        = mux.Vars(dataContext.Request)
		transcationID = params["transaction_id"]
	)

	defer tracer.SpanFinish(ctx)

	err := dataContext.Cast(&param)
	if err != nil {
		log.Printf("error parsing query url: %v", err)
		return *appctx.NewResponse().WithMsgKey(consts.RespValidationError)
	}

	result := u.resolve.ResolverUpdateTransaction(ctx, param, transcationID)
	switch result.Code {
	case http.StatusOK:
		log.Println("success update data transaction consumer")
	default:
		log.Printf("got response %v", result.Message)
	}

	return result
}
