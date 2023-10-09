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

type consumerLimitRepositories struct {
	db databasex.Adapter
}

func NewConsumerLimitRepositories(db databasex.Adapter) LimitAmountDB {
	return &consumerLimitRepositories{db: db}
}

func (c *consumerLimitRepositories) Store(ctx context.Context, param entity.ConsumerLimits) (entity.ConsumerLimits, error) {
	ctx = tracer.SpanStart(ctx, "repo.stamps")
	defer tracer.SpanFinish(ctx)

	q := `INSERT INTO consumer_limits
    	(
    	 	id_limit,
    	 	id_consumer,
    	 	tenor,
    	 	limit_amount,
    	 	created_at,
    	 	updated_at
    	)
    	VALUES (?,?,?,?,?,?)`

	qValues := []interface{}{
		param.IDLimit,
		param.IDConsumer,
		param.Tenor,
		param.LimitAmount,
		param.CreatedAt,
		param.UpdatedAt,
	}

	err := c.db.Transact(ctx, sql.LevelSerializable, func(db *databasex.DB) error {
		_, err := db.Exec(ctx, q, qValues...)
		if err != nil {
			log.Printf("error when insert consumer limit %v", err)
			return err
		}
		return nil
	})

	if err != nil {
		log.Printf("error transact consumer limit %v", err)
		return entity.ConsumerLimits{}, err
	}

	return param, nil
}

func (c *consumerLimitRepositories) Update(ctx context.Context, param entity.ConsumerLimits, where entity.ConsumerLimits) error {
	var (
		err error
	)

	q, vals, err := builderx.StructToQueryUpdate(param, where, "consumer_limits", "db")
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

func (c *consumerLimitRepositories) FindOne(ctx context.Context, where entity.ConsumerLimits) (*entity.ConsumerLimits, error) {
	var (
		result entity.ConsumerLimits
		err    error
	)

	wq, err := builderx.StructToMySqlQueryWhere(where, "db")
	if err != nil {
		return nil, err
	}

	q := `SELECT
			id_limit,
			id_consumer,
			tenor,
			limit_amount
		FROM consumer_limits %s LIMIT 1`

	err = c.db.QueryRow(ctx, &result, fmt.Sprintf(q, wq.Query), wq.Values...)
	if err == sql.ErrNoRows {
		return nil, err
	}

	return &result, nil
}

func (c *consumerLimitRepositories) Find(ctx context.Context) ([]entity.ConsumerLimits, error) {
	var (
		result []entity.ConsumerLimits
		err    error
	)

	q := `SELECT
			id_limit,
			id_consumer,
			tenor,
			limit_amount,
			created_at,
			updated_at
		FROM consumer_limits`

	err = c.db.Query(ctx, &result, q)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return result, nil
}
