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

func (c *consumer) ResolverFindByID(ctx context.Context, iDConsumer string) appctx.Response {
	if iDConsumer == "all" {
		find, err := c.co.Find(ctx)
		if err != nil {
			log.Printf("find all data consumer got error %v", err.Error())
			return *appctx.NewResponse().WithCode(consts.CodeInternalServerError)
		}

		return *appctx.NewResponse().WithCode(consts.CodeSuccess).WithData(find).WithMessage("successfully get all data consumer")
	}

	findOne, err := c.co.FindOne(ctx, entity.Consumers{
		IDConsumer: iDConsumer,
	})

	if err != nil {
		log.Printf("find one consumer got error %v", err.Error())
		return *appctx.NewResponse().WithCode(http.StatusInternalServerError)
	}

	if findOne == nil {
		return *appctx.NewResponse().WithCode(http.StatusNotFound).WithMessage("data consumer not found")
	}

	response := presentations.ResponseConsumer{
		IDConsumer:     findOne.IDConsumer,
		IdentityNumber: findOne.NIK,
		FullName:       findOne.FullName,
		LegalName:      findOne.LegalName,
		BirthPlace:     findOne.Pob,
		BirthDate:      findOne.Dob.Format(consts.LayoutDateFormat),
		Salary:         findOne.Salary,
		URL: presentations.URL{
			ImageIdentity: findOne.ImageIdentity,
			ImageSelfie:   findOne.ImageSelfie,
		},
	}

	return *appctx.NewResponse().
		WithCode(http.StatusOK).
		WithData(response).
		WithMessage("successfully fetch data consumer")
}
