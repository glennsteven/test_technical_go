package consumer_ucase

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"technical_test_go/technical_test_go/internal/appctx"
	"technical_test_go/technical_test_go/internal/consts"
	"technical_test_go/technical_test_go/internal/consumer"
	"technical_test_go/technical_test_go/internal/presentations"
	"technical_test_go/technical_test_go/internal/ucase/contract"
	"technical_test_go/technical_test_go/pkg/tracer"
)

type updateConsumer struct {
	resolve consumer.Resolve
}

func NewUpdateConsumer(resolve consumer.Resolve) contract.UseCase {
	return &updateConsumer{resolve: resolve}
}

func (f *updateConsumer) Serve(dataContext *appctx.Data) appctx.Response {
	var (
		param      presentations.PayloadConsumer
		ctx        = tracer.SpanStart(dataContext.Request.Context(), "ucase.UpdateConsumer")
		params     = mux.Vars(dataContext.Request)
		consumerID = params["consumer_id"]
	)

	defer tracer.SpanFinish(ctx)

	err := dataContext.Cast(&param)
	if err != nil {
		log.Printf("error parsing query url: %v", err)
		return *appctx.NewResponse().WithMsgKey(consts.RespValidationError)
	}

	result := f.resolve.ResolverUpdate(ctx, param, consumerID)
	switch result.Code {
	case http.StatusOK:
		log.Println("success update data consumer")
	default:
		log.Printf("got response %v", result.Message)
	}

	return result
}
