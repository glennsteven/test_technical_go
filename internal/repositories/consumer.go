package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"technical_test_go/technical_test_go/internal/entity"
	"technical_test_go/technical_test_go/pkg/builderx"
	"technical_test_go/technical_test_go/pkg/databasex"
	"technical_test_go/technical_test_go/pkg/tracer"
)

type consumerRepositories struct {
	db databasex.Adapter
}

func NewConsumerRepositories(db databasex.Adapter) ConsumerDB {
	return &consumerRepositories{db: db}
}

func (c *consumerRepositories) Store(ctx context.Context, param entity.Consumers) (entity.Consumers, error) {
	ctx = tracer.SpanStart(ctx, "repo.stamps")
	defer tracer.SpanFinish(ctx)

	q := `INSERT INTO consumers
    	(
            id_consumer,
            full_name,
    		nik,
    		legal_name,
    		pob,
    	 	dob,
    	 	salary,
    	 	image_identity,
    	 	image_selfie,
    	 	created_at,
    	 	updated_at
		)
		VALUES (?,?,?,?,?,?,?,?,?,?,?)`

	qValues := []interface{}{
		param.IDConsumer,
		param.FullName,
		param.NIK,
		param.LegalName,
		param.Pob,
		param.Dob,
		param.Salary,
		param.ImageIdentity,
		param.ImageSelfie,
		param.CreatedAt,
		param.UpdatedAt,
	}

	err := c.db.Transact(ctx, sql.LevelSerializable, func(db *databasex.DB) error {
		_, err := db.Exec(ctx, q, qValues...)
		if err != nil {
			log.Printf("error when insert consumer: %v", err)
			return err
		}
		return nil
	})

	if err != nil {
		log.Printf("error when insert consumer: %v", err)
		return entity.Consumers{}, err
	}

	return param, nil
}

func (c *consumerRepositories) FindOne(ctx context.Context, where entity.Consumers) (*entity.Consumers, error) {
	var (
		result entity.Consumers
		err    error
	)

	wq, err := builderx.StructToMySqlQueryWhere(where, "db")
	if err != nil {
		return nil, err
	}

	q := `SELECT
			id_consumer,
			full_name,
			nik,
			legal_name,
			pob,
			dob,
			salary,
			image_identity,
			image_selfie
		FROM consumers %s LIMIT 1`

	err = c.db.QueryRow(ctx, &result, fmt.Sprintf(q, wq.Query), wq.Values...)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &result, nil
}

func (c *consumerRepositories) Update(ctx context.Context, param entity.Consumers, wheres entity.Consumers) error {
	var (
		err error
	)

	q, vals, err := builderx.StructToQueryUpdate(param, wheres, "consumers", "db")
	if err != nil {
		return err
	}

	if len(vals) == 0 {
		return fmt.Errorf("query %s , empty value parameters %v", q, vals)
	}

	_, err = c.db.Exec(ctx, q, vals...)
	if err != nil {
		tracer.SpanError(ctx, err)
		return err
	}

	return err
}

func (c *consumerRepositories) Find(ctx context.Context) ([]entity.Consumers, error) {
	var (
		result []entity.Consumers
		err    error
	)

	q := `SELECT
			id_consumer,
			full_name,
			nik,
			legal_name,
			pob,
			dob,
			salary,
			image_identity,
			image_selfie,
			created_at,
			updated_at
		FROM consumers`

	err = c.db.Query(ctx, &result, q)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return result, nil
}
