package consumer

import (
	"context"
	"technical_test_go/technical_test_go/internal/appctx"
	"technical_test_go/technical_test_go/internal/presentations"
)

type Resolve interface {
	ResolverStore(ctx context.Context, param presentations.PayloadConsumer) appctx.Response
	ResolverFindByID(ctx context.Context, iDConsumer string) appctx.Response
	ResolverUpdate(ctx context.Context, param presentations.PayloadConsumer, id string) appctx.Response
}
