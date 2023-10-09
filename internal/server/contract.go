// Package router
package server

import (
	"context"
	"net/http"

	"technical_test_go/technical_test_go/internal/appctx"
	ucase "technical_test_go/technical_test_go/internal/ucase/contract"
)

// httpHandlerFunc abstraction for http handler
type httpHandlerFunc func(request *http.Request, svc ucase.UseCase, conf *appctx.Config) appctx.Response

// Server contract
type Server interface {
	Run(context.Context) error
	Done()
	Config() *appctx.Config
}
