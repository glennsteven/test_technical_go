package middleware

import (
	"net/http"
	"technical_test_go/technical_test_go/internal/appctx"
	"technical_test_go/technical_test_go/internal/consts"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	rsp := appctx.NewResponse().
		WithMsgKey(consts.RespNoRouteFound).
		Generate()
	w.Header().Set("Content-Type", consts.HeaderContentTypeJSON)
	w.WriteHeader(rsp.Code)
	w.Write(rsp.Byte())
	return
}
