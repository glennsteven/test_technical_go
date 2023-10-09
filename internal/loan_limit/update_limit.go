package loan_limit

import (
	"context"
	"log"
	"technical_test_go/technical_test_go/internal/appctx"
	"technical_test_go/technical_test_go/internal/consts"
	"technical_test_go/technical_test_go/internal/entity"
	"technical_test_go/technical_test_go/internal/presentations"
)

func (l *limitAmount) ResolverUpdateLimit(ctx context.Context, payload presentations.PayloadAddLimit, iDLimit string) appctx.Response {
	findLimitConsumer, err := l.limit.FindOne(ctx, entity.ConsumerLimits{IDLimit: iDLimit})
	if err != nil {
		log.Printf("find data consumer limit got error %v", err)
		return *appctx.NewResponse().WithCode(consts.CodeInternalServerError)
	}

	if findLimitConsumer == nil {
		return *appctx.NewResponse().
			WithCode(consts.CodeNotFound).
			WithData("Data Limit information not found")
	}

	err = l.limit.Update(ctx, entity.ConsumerLimits{
		Tenor:       payload.Tenor,
		LimitAmount: payload.Amount,
	}, entity.ConsumerLimits{
		IDLimit: iDLimit,
	})
	if err != nil {
		log.Printf("update limit consumer got error %v", err.Error())
		return *appctx.NewResponse().WithCode(consts.CodeInternalServerError)
	}

	return *appctx.NewResponse().
		WithCode(consts.CodeSuccess).
		WithMessage("Loan limit information has been successfully updated.")
}
