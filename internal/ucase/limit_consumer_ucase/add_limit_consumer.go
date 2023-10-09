package limit_consumer_ucase

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"technical_test_go/technical_test_go/internal/appctx"
	"technical_test_go/technical_test_go/internal/consts"
	"technical_test_go/technical_test_go/internal/loan_limit"
	"technical_test_go/technical_test_go/internal/presentations"
	"technical_test_go/technical_test_go/internal/ucase/contract"
)

type addLimitConsumerUcase struct {
	resolve loan_limit.Resolve
}

func NewAddLimitConsumerUcase(resolve loan_limit.Resolve) contract.UseCase {
	return &addLimitConsumerUcase{resolve: resolve}
}

func (a *addLimitConsumerUcase) Serve(dataContext *appctx.Data) appctx.Response {
	var (
		param      presentations.PayloadAddLimit
		ctx        = dataContext.Request.Context()
		params     = mux.Vars(dataContext.Request)
		consumerID = params["consumer_id"]
	)

	err := dataContext.Cast(&param)
	if err != nil {
		log.Printf("error parsing query url: %v", err)
		return *appctx.NewResponse().WithMsgKey(consts.RespValidationError).WithCode(consts.CodeUnprocessableEntity)
	}

	result := a.resolve.ResolverStoreLimit(ctx, param, consumerID)
	switch result.Code {
	case http.StatusCreated, http.StatusOK:
		log.Printf("successfully add limit consumer")
	default:
		log.Printf("got response error %v", result.Message)
	}

	return result
}
