package loan_limit

import (
	"context"
	"technical_test_go/technical_test_go/internal/appctx"
	"technical_test_go/technical_test_go/internal/presentations"
)

type Resolve interface {
	ResolverStoreLimit(ctx context.Context, payload presentations.PayloadAddLimit, iDConsumer string) appctx.Response
	ResolverUpdateLimit(ctx context.Context, payload presentations.PayloadAddLimit, iDLimit string) appctx.Response
	ResolverGetLimit(ctx context.Context, iDLimit string) appctx.Response
}
