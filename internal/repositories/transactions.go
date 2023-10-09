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

type transactionRepositories struct {
	db databasex.Adapter
}

func NewTransactionRepositories(db databasex.Adapter) TransactionDB {
	return &transactionRepositories{db: db}
}

func (t *transactionRepositories) Store(ctx context.Context, param entity.Transactions) (entity.Transactions, error) {
	ctx = tracer.SpanStart(ctx, "repo.stamps")
	defer tracer.SpanFinish(ctx)

	q := `INSERT INTO transactions
    	(
    	 	id_transaction,
    	 	id_consumer,
    	 	contract_number,
    	 	otr,
    	 	fee_admin,
    	 	installment_amount,
    	 	total_interest,
    	 	asset_name,
    	 	transaction_date,
    	 	created_at,
    	 	updated_at
    	)
    	VALUES (?,?,?,?,?,?,?,?,?,?,?)`

	qValues := []interface{}{
		param.IDTransaction,
		param.IDConsumer,
		param.ContractNumber,
		param.Otr,
		param.FeeAdmin,
		param.InstallmentAmount,
		param.TotalInterest,
		param.AssetName,
		param.TransactionDate,
		param.CreatedAt,
		param.UpdatedAt,
	}

	err := t.db.Transact(ctx, sql.LevelSerializable, func(db *databasex.DB) error {
		_, err := db.Exec(ctx, q, qValues...)
		if err != nil {
			log.Printf("error when insert data transaction consumer %v", err.Error())
			return err
		}
		return nil
	})

	if err != nil {
		log.Printf("error transact transaction consumer %v", err.Error())
		return entity.Transactions{}, err
	}

	return param, nil
}

func (t *transactionRepositories) Update(ctx context.Context, param entity.Transactions, where entity.Transactions) error {
	var (
		err error
	)

	q, vals, err := builderx.StructToQueryUpdate(param, where, "transactions", "db")
	if err != nil {
		return err
	}

	if len(vals) == 0 {
		return fmt.Errorf("query %s , empty value parameters %v", q, vals)
	}

	_, err = t.db.Exec(ctx, q, vals...)
	if err != nil {
		tracer.SpanError(ctx, err)
		return err
	}

	return err
}

func (t *transactionRepositories) FindOne(ctx context.Context, where entity.Transactions) (*entity.Transactions, error) {
	var (
		result entity.Transactions
		err    error
	)
	wq, err := builderx.StructToMySqlQueryWhere(where, "db")
	if err != nil {
		return nil, err
	}

	q := `SELECT 
			id_transaction,
    	 	id_consumer,
    	 	contract_number,
    	 	otr,
    	 	fee_admin,
    	 	installment_amount,
    	 	total_interest,
    	 	asset_name,
    	 	transaction_date,
    	 	created_at,
    	 	updated_at
		FROM transactions %s LIMIT 1`

	err = t.db.QueryRow(ctx, &result, fmt.Sprintf(q, wq.Query), wq.Values...)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &result, nil
}

func (t *transactionRepositories) Find(ctx context.Context) ([]entity.Transactions, error) {
	var (
		result []entity.Transactions
	)

	q := `SELECT 
			id_transaction,
    	 	id_consumer,
    	 	contract_number,
    	 	otr,
    	 	fee_admin,
    	 	installment_amount,
    	 	total_interest,
    	 	asset_name,
    	 	transaction_date,
    	 	created_at,
    	 	updated_at
		FROM transactions `

	err := t.db.Query(ctx, &result, q)
	if err == sql.ErrNoRows {
		return nil, err
	} else if err != nil {
		return nil, err
	}

	return result, nil
}
