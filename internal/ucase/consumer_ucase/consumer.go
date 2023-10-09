package consumer_ucase

import (
	"fmt"
	"log"
	"net/http"
	"technical_test_go/technical_test_go/internal/appctx"
	"technical_test_go/technical_test_go/internal/consts"
	"technical_test_go/technical_test_go/internal/consumer"
	"technical_test_go/technical_test_go/internal/presentations"
	"technical_test_go/technical_test_go/internal/ucase/contract"
	"technical_test_go/technical_test_go/pkg/logger"
	"technical_test_go/technical_test_go/pkg/tracer"
)

type consumerUcase struct {
	resolve consumer.Resolve
}

func NewConsumerUCase(resolve consumer.Resolve) contract.UseCase {
	return &consumerUcase{resolve: resolve}
}

func (c *consumerUcase) Serve(data *appctx.Data) appctx.Response {
	var (
		param presentations.PayloadConsumer
		ctx   = data.Request.Context()
	)

	err := data.Cast(&param)
	if err != nil {
		log.Printf("error parsing query url: %v", err)
		return *appctx.NewResponse().WithMsgKey(consts.RespValidationError).WithCode(consts.CodeUnprocessableEntity)
	}

	params, err := parseForm(data.Request)
	if err != nil {
		log.Printf("error when parse form: %v", err)
		return *appctx.NewResponse().WithCode(http.StatusBadRequest).WithError(err)
	}

	err = validateParams(params)
	if err != nil {
		log.Printf("error when validate params: %v", err.Error())
		return *appctx.NewResponse().WithCode(http.StatusUnprocessableEntity).WithMessage("Validation(s) Error").WithError(err)
	}

	result := c.resolve.ResolverStore(ctx, params)
	switch result.Code {
	case http.StatusCreated, http.StatusOK:
		log.Printf("success create consumer")
	default:
		log.Printf("got response error %v", result.Message)
		tracer.SpanError(ctx, err)
		logger.WarnWithContext(ctx, fmt.Sprintf("got response error %v", result.Message))
	}

	return result
}
