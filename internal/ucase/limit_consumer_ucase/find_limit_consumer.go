package limit_consumer_ucase

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"technical_test_go/technical_test_go/internal/appctx"
	"technical_test_go/technical_test_go/internal/loan_limit"
	"technical_test_go/technical_test_go/internal/ucase/contract"
	"technical_test_go/technical_test_go/pkg/tracer"
)

type findLimitConsumerUcase struct {
	resolve loan_limit.Resolve
}

func NewFindLimitConsumerUcase(resolve loan_limit.Resolve) contract.UseCase {
	return &findLimitConsumerUcase{resolve: resolve}
}

func (f *findLimitConsumerUcase) Serve(dataContext *appctx.Data) appctx.Response {
	var (
		ctx     = tracer.SpanStart(dataContext.Request.Context(), "ucase.FindLimitConsumer")
		params  = mux.Vars(dataContext.Request)
		limitID = params["limit_id"]
	)

	result := f.resolve.ResolverGetLimit(ctx, limitID)
	switch result.Code {
	case http.StatusOK:
		log.Printf("success get data limit consumer %v", result.Message)
	default:
		log.Printf("got response %v", result.Message)
	}
	return result
}
