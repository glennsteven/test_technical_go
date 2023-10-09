// Package router
package router

import (
	"context"
	"encoding/json"
	"net/http"
	"runtime/debug"
	"technical_test_go/technical_test_go/internal/appctx"
	"technical_test_go/technical_test_go/internal/bootstrap"
	"technical_test_go/technical_test_go/internal/consts"
	"technical_test_go/technical_test_go/internal/consumer"
	"technical_test_go/technical_test_go/internal/handler"
	"technical_test_go/technical_test_go/internal/loan_limit"
	"technical_test_go/technical_test_go/internal/middleware"
	"technical_test_go/technical_test_go/internal/repositories"
	"technical_test_go/technical_test_go/internal/transactions"
	"technical_test_go/technical_test_go/internal/ucase"
	"technical_test_go/technical_test_go/internal/ucase/consumer_ucase"
	"technical_test_go/technical_test_go/internal/ucase/limit_consumer_ucase"
	"technical_test_go/technical_test_go/internal/ucase/transaction_ucase"
	"technical_test_go/technical_test_go/pkg/logger"
	"technical_test_go/technical_test_go/pkg/msgx"
	"technical_test_go/technical_test_go/pkg/routerkit"

	ucaseContract "technical_test_go/technical_test_go/internal/ucase/contract"
)

type router struct {
	config *appctx.Config
	router *routerkit.Router
}

// NewRouter initialize new router wil return Router Interface
func NewRouter(cfg *appctx.Config) Router {
	return &router{
		config: cfg,
		router: routerkit.NewRouter(routerkit.WithServiceName(cfg.App.AppName)),
	}
}

func (rtr *router) handle(hfn httpHandlerFunc, svc ucaseContract.UseCase, mdws ...middleware.MiddlewareFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lang := r.Header.Get(consts.HeaderLanguageKey)
		if !msgx.HaveLang(consts.RespOK, lang) {
			lang = rtr.config.App.DefaultLang
			r.Header.Set(consts.HeaderLanguageKey, lang)
		}

		defer func() {
			err := recover()
			if err != nil {
				w.Header().Set(consts.HeaderContentTypeKey, consts.HeaderContentTypeJSON)
				w.WriteHeader(http.StatusInternalServerError)
				res := appctx.Response{
					Code: consts.CodeInternalServerError,
				}

				res.WithLang(lang)
				logger.Error(logger.MessageFormat("error %v", string(debug.Stack())))
				json.NewEncoder(w).Encode(res.Byte())

				return
			}
		}()

		ctx := context.WithValue(r.Context(), "access", map[string]interface{}{
			"path":      r.URL.Path,
			"remote_ip": r.RemoteAddr,
			"method":    r.Method,
		})

		req := r.WithContext(ctx)

		// validate middleware
		if err := middleware.FilterFunc(w, req, rtr.config, mdws); err != nil {
			return
		}

		resp := hfn(req, svc, rtr.config)
		resp.WithLang(lang)
		rtr.response(w, resp)
	}
}

func (rtr *router) response(w http.ResponseWriter, resp appctx.Response) {
	w.Header().Set(consts.HeaderContentTypeKey, consts.HeaderContentTypeJSON)
	resp.Generate()
	w.WriteHeader(resp.Code)
	w.Write(resp.Byte())
	return
}

// Route preparing http router and will return mux router object
func (rtr *router) Route() *routerkit.Router {

	rtr.router.NotFoundHandler = http.HandlerFunc(middleware.NotFound)
	root := rtr.router.PathPrefix("/").Subrouter()
	in := root.PathPrefix("/in/").Subrouter()
	checklife := root.PathPrefix("/").Subrouter()
	inV1 := in.PathPrefix("/v1/").Subrouter()

	management := in.PathPrefix("/management/").Subrouter()

	_ = inV1

	// create session database for single database
	db := bootstrap.RegistryDatabase(rtr.config.WriteDB)
	// use case
	healthy := ucase.NewHealthCheck()
	repoConsumer := repositories.NewConsumerRepositories(db)
	repoLimitConsumer := repositories.NewConsumerLimitRepositories(db)
	repoTransaction := repositories.NewTransactionRepositories(db)

	entityConsumer := consumer.NewConsumerResolve(repoConsumer, rtr.config)
	entityLimitConsumer := loan_limit.NewLimitAmount(repoConsumer, repoLimitConsumer)
	entityTransaction := transactions.NewTransactions(repoConsumer, repoTransaction, repoLimitConsumer)

	ucaseConsumer := consumer_ucase.NewConsumerUCase(entityConsumer)
	ucaseFindConsumer := consumer_ucase.NewFindConsumer(entityConsumer)
	ucaseUpdateConsumer := consumer_ucase.NewUpdateConsumer(entityConsumer)

	ucaseLimitConsumer := limit_consumer_ucase.NewAddLimitConsumerUcase(entityLimitConsumer)
	ucaseUpdateLimitConsumer := limit_consumer_ucase.NewUpdateLimitConsumerUcase(entityLimitConsumer)
	ucaseFindLimitConsumer := limit_consumer_ucase.NewFindLimitConsumerUcase(entityLimitConsumer)

	ucaseAddTransaction := transaction_ucase.NewTransactionUcase(entityTransaction)
	ucaseUpdateTransaction := transaction_ucase.NewUpdateTransactionUcase(entityTransaction)
	ucaseFindTransaction := transaction_ucase.NewFindTransactionUcase(entityTransaction)
	ucaseManagementTransaction := transaction_ucase.NewManagementTransactionUcase(entityTransaction)

	// healthy
	checklife.HandleFunc("/checklife", rtr.handle(
		handler.HttpRequest,
		healthy,
	)).Methods(http.MethodGet)

	in.HandleFunc("/consumer", rtr.handle(
		handler.HttpRequest,
		ucaseConsumer,
	)).Methods(http.MethodPost)

	in.HandleFunc("/consumer/{consumer_id}", rtr.handle(
		handler.HttpRequest,
		ucaseFindConsumer,
	)).Methods(http.MethodGet)

	in.HandleFunc("/consumer/{consumer_id}", rtr.handle(
		handler.HttpRequest,
		ucaseUpdateConsumer,
	)).Methods(http.MethodPut)

	in.HandleFunc("/limit/{consumer_id}", rtr.handle(
		handler.HttpRequest,
		ucaseLimitConsumer,
	)).Methods(http.MethodPost)

	in.HandleFunc("/limit/{limit_id}", rtr.handle(
		handler.HttpRequest,
		ucaseUpdateLimitConsumer,
	)).Methods(http.MethodPut)

	in.HandleFunc("/limit/{limit_id}", rtr.handle(
		handler.HttpRequest,
		ucaseFindLimitConsumer,
	)).Methods(http.MethodGet)

	in.HandleFunc("/transaction/{consumer_id}", rtr.handle(
		handler.HttpRequest,
		ucaseAddTransaction,
	)).Methods(http.MethodPost)

	in.HandleFunc("/transaction/{transaction_id}", rtr.handle(
		handler.HttpRequest,
		ucaseUpdateTransaction,
	)).Methods(http.MethodPut)

	in.HandleFunc("/transaction/{transaction_id}", rtr.handle(
		handler.HttpRequest,
		ucaseFindTransaction,
	)).Methods(http.MethodGet)

	management.HandleFunc("/transaction/{consumer_id}", rtr.handle(
		handler.HttpRequest,
		ucaseManagementTransaction,
	)).Methods(http.MethodGet)

	return rtr.router
}
