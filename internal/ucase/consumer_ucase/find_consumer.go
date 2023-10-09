package consumer_ucase

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"technical_test_go/technical_test_go/internal/appctx"
	"technical_test_go/technical_test_go/internal/consumer"
	"technical_test_go/technical_test_go/internal/ucase/contract"
	"technical_test_go/technical_test_go/pkg/tracer"
)

type findConsumer struct {
	resolve consumer.Resolve
}

func NewFindConsumer(resolve consumer.Resolve) contract.UseCase {
	return &findConsumer{resolve: resolve}
}

func (f *findConsumer) Serve(dataContext *appctx.Data) appctx.Response {
	var (
		ctx        = tracer.SpanStart(dataContext.Request.Context(), "ucase.FindConsumer")
		params     = mux.Vars(dataContext.Request)
		consumerID = params["consumer_id"]
	)

	result := f.resolve.ResolverFindByID(ctx, consumerID)
	switch result.Code {
	case http.StatusOK:
		log.Println("success get data consumer")
	default:
		log.Printf("got response %v", result.Message)
	}

	return result
}
