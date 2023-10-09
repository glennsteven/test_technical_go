package loan_limit

import (
	"context"
	"log"
	"technical_test_go/technical_test_go/internal/appctx"
	"technical_test_go/technical_test_go/internal/consts"
	"technical_test_go/technical_test_go/internal/entity"
)

func (l *limitAmount) ResolverGetLimit(ctx context.Context, iDLimit string) appctx.Response {
	if iDLimit == "all" {
		find, err := l.limit.Find(ctx)
		if err != nil {
			log.Printf("find all data limit consumer got error %v", err.Error())
			return *appctx.NewResponse().WithCode(consts.CodeInternalServerError)
		}

		return *appctx.NewResponse().
			WithCode(consts.CodeSuccess).
			WithData(find).
			WithMessage("Success get information limit consumers")
	}

	findLimitConsumer, err := l.limit.FindOne(ctx, entity.ConsumerLimits{IDLimit: iDLimit})
	if err != nil {
		log.Printf("find limit consumer got error %v", err.Error())
		return *appctx.NewResponse().WithCode(consts.CodeInternalServerError)
	}

	if findLimitConsumer == nil {
		return *appctx.NewResponse().
			WithCode(consts.CodeNotFound).
			WithMessage("Data limit consumer not found")
	}

	return *appctx.NewResponse().
		WithCode(consts.CodeSuccess).
		WithMessage("Success get information limit consumer").
		WithData(findLimitConsumer)
}
