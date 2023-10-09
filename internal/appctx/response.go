// Package appctx
package appctx

import (
	"encoding/json"
	"sync"

	"technical_test_go/technical_test_go/internal/consts"
	"technical_test_go/technical_test_go/pkg/msgx"
)

var (
	rsp    *Response
	oneRsp sync.Once
)

// Response presentation contract object
type Response struct {
	Code    int    `json:"code"`
	Message any    `json:"message,omitempty"`
	Errors  any    `json:"errors,omitempty"`
	Data    any    `json:"data,omitempty"`
	lang    string `json:"-"`
	Meta    any    `json:"meta,omitempty"`
	msgKey  string
}

// MetaData represent meta data response for multi data
type MetaData struct {
	Pagination Pagination `json:"pagination"`
}

// Pagination represent meta data pagination for multi data
type Pagination struct {
	Page       int64 `json:"page"`
	Limit      int64 `json:"limit"`
	TotalPage  int64 `json:"total_page"`
	TotalCount int64 `json:"total_count"`
}

// GetMessage method to transform response name var to message detail
func (r *Response) GetMessage() string {
	return msgx.Get(r.msgKey, r.lang).Text()
}

// Generate setter message
func (r *Response) Generate() *Response {
	if r.lang == "" {
		r.lang = consts.LangDefault
	}
	msg := msgx.Get(r.msgKey, r.lang)
	if r.Message == nil {
		r.Message = msg.Text()
	}

	if r.Code == 0 {
		r.Code = msg.Status()
	}

	return r
}

// WithCode setter response var name
func (r *Response) WithCode(c int) *Response {
	r.Code = c
	return r
}

// WithData setter data response
func (r *Response) WithData(v any) *Response {
	r.Data = v
	return r
}

// WithError setter error messages
func (r *Response) WithError(v any) *Response {
	r.Errors = v
	return r
}

func (r *Response) WithMsgKey(v string) *Response {
	r.msgKey = v
	return r
}

// WithMeta setter meta data response
func (r *Response) WithMeta(v any) *Response {
	r.Meta = v
	return r
}

// WithLang setter language response
func (r *Response) WithLang(v string) *Response {
	r.lang = v
	return r
}

// WithMessage setter custom message response
func (r *Response) WithMessage(v any) *Response {
	if v != nil {
		r.Message = v
	}

	return r
}

func (r *Response) Byte() []byte {
	if r.Code == 0 || r.Message == nil {
		r.Generate()
	}

	b, _ := json.Marshal(r)
	return b
}

// NewResponse initialize response
func NewResponse() *Response {
	oneRsp.Do(func() {
		rsp = &Response{}
	})

	// clone response
	x := *rsp

	return &x
}
