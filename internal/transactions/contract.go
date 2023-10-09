package transactions

import (
	"context"
	"technical_test_go/technical_test_go/internal/appctx"
	"technical_test_go/technical_test_go/internal/presentations"
)

type Resolve interface {
	ResolverStoreTransaction(ctx context.Context, payload presentations.PayloadAddTransaction) appctx.Response
	ResolverUpdateTransaction(ctx context.Context, payload presentations.PayloadAddTransaction, iDTransaction string) appctx.Response
	ResolverGetTransactions(ctx context.Context, iDTransaction string) appctx.Response
	ResolverManagementTransaction(ctx context.Context, iDConsumer string) appctx.Response
}
