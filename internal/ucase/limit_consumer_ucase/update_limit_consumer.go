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
	"technical_test_go/technical_test_go/pkg/tracer"
)

type updateLimitConsumerUcase struct {
	resolve loan_limit.Resolve
}

func NewUpdateLimitConsumerUcase(resolve loan_limit.Resolve) contract.UseCase {
	return &updateLimitConsumerUcase{resolve: resolve}
}

func (u *updateLimitConsumerUcase) Serve(dataContext *appctx.Data) appctx.Response {
	var (
		param   presentations.PayloadAddLimit
		ctx     = tracer.SpanStart(dataContext.Request.Context(), "ucase.UpdateConsumer")
		params  = mux.Vars(dataContext.Request)
		limitID = params["limit_id"]
	)

	defer tracer.SpanFinish(ctx)

	err := dataContext.Cast(&param)
	if err != nil {
		log.Printf("error parsing query url: %v", err)
		return *appctx.NewResponse().WithMsgKey(consts.RespValidationError)
	}

	result := u.resolve.ResolverUpdateLimit(ctx, param, limitID)
	switch result.Code {
	case http.StatusOK:
		log.Println("success update data limit consumer")
	default:
		log.Printf("got response %v", result.Message)
	}

	return result
}
