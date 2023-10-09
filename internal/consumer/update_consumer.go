package consumer

import (
	"context"
	"log"
	"net/http"
	"technical_test_go/technical_test_go/internal/appctx"
	"technical_test_go/technical_test_go/internal/consts"
	"technical_test_go/technical_test_go/internal/entity"
	"technical_test_go/technical_test_go/internal/presentations"
)

func (c *consumer) ResolverUpdate(ctx context.Context, param presentations.PayloadConsumer, id string) appctx.Response {

	err := c.co.Update(ctx, entity.Consumers{
		FullName: param.FullName,
		Salary:   param.Salary,
	}, entity.Consumers{IDConsumer: id})
	if err != nil {
		log.Printf("update data consumer got error %v", err.Error())
		return *appctx.NewResponse().WithCode(http.StatusInternalServerError)
	}

	return *appctx.NewResponse().WithCode(consts.CodeSuccess).WithMessage("successfully update data consumer")
}
