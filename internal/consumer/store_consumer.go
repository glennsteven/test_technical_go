package consumer

import (
	"context"
	"fmt"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/google/uuid"
	"golang.org/x/sync/errgroup"
	"log"
	"technical_test_go/technical_test_go/internal/appctx"
	"technical_test_go/technical_test_go/internal/consts"
	"technical_test_go/technical_test_go/internal/entity"
	"technical_test_go/technical_test_go/internal/presentations"
	"technical_test_go/technical_test_go/pkg/tracer"
	"time"
)

func (c *consumer) ResolverStore(ctx context.Context, param presentations.PayloadConsumer) appctx.Response {
	ctx = tracer.SpanStart(ctx, "Resolve.NewConsumerResolve")
	defer tracer.SpanFinish(ctx)
	var (
		errs                    []presentations.MessageErrorValidation
		uploadKtp, uploadSelfie *uploader.UploadResult
		cloudName               = c.cfg.Cloudinary.CloudName
		apiKey                  = c.cfg.Cloudinary.ApiKey
		secretKey               = c.cfg.Cloudinary.ApiSecret
	)

	findNik, err := c.co.FindOne(ctx, entity.Consumers{
		NIK: param.NIK,
	})

	if err != nil {
		log.Printf("find data nik got error %v", err.Error())
		return *appctx.NewResponse().WithCode(consts.CodeInternalServerError)
	}

	if findNik != nil {
		return *appctx.NewResponse().WithCode(consts.CodeDuplicateEntry).WithMessage("nik has been use by another user")
	}

	group, _ := errgroup.WithContext(ctx)

	group.Go(func() error {
		publicIDKTP := fmt.Sprintf("image/identity/%s", uuid.New().String())
		uploadKtp, err = UploadImage(ctx, cloudName, apiKey, secretKey, publicIDKTP, param.ImageIdentity.Buffer)
		if err != nil {
			log.Printf("upload image identity got error %v", err.Error())
			return err
		}
		return err
	})

	err = group.Wait()
	if err != nil {
		log.Printf("failed upload image identity")
		errs = append(errs, presentations.MessageErrorValidation{
			Field:    "image_identity",
			Messages: []string{"file image identity failed upload"},
		})
	}

	group.Go(func() error {
		publicIDSelfie := fmt.Sprintf("image/selfie/%s", uuid.New().String())
		uploadSelfie, err = UploadImage(ctx, cloudName, apiKey, secretKey, publicIDSelfie, param.ImageSelfie.Buffer)
		if err != nil {
			log.Printf("upload image selfie got error %v", err.Error())
			return err
		}
		return err
	})

	err = group.Wait()
	if err != nil {
		log.Printf("failed upload image selfie")
		errs = append(errs, presentations.MessageErrorValidation{
			Field:    "image_selfie",
			Messages: []string{"file image selfie failed upload"},
		})
	}

	paramDOB, _ := time.Parse(consts.LayoutDateFormat, param.Dob)

	saveConsumer := entity.Consumers{
		IDConsumer:    uuid.New().String(),
		FullName:      param.FullName,
		NIK:           param.NIK,
		LegalName:     param.LegalName,
		Pob:           param.Pob,
		Dob:           paramDOB,
		Salary:        param.Salary,
		ImageIdentity: uploadKtp.URL,
		ImageSelfie:   uploadSelfie.URL,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	result, err := c.co.Store(ctx, saveConsumer)
	if err != nil {
		log.Printf("error when store consumer: %v", err)
		return *appctx.NewResponse().WithCode(consts.CodeInternalServerError)
	}

	response := presentations.ResponseConsumer{
		IDConsumer:     result.IDConsumer,
		IdentityNumber: result.NIK,
		FullName:       result.FullName,
		LegalName:      result.LegalName,
		BirthPlace:     result.Pob,
		BirthDate:      result.Dob.Format(consts.LayoutDateFormat),
		Salary:         result.Salary,
		URL: presentations.URL{
			ImageIdentity: result.ImageIdentity,
			ImageSelfie:   result.ImageSelfie,
		},
	}

	return *appctx.NewResponse().WithCode(consts.CodeCreated).WithData(response).WithMessage("successfully create new consumer")
}
