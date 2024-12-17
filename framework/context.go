package framework

import (
	"context"
	"encoding/json"
	"net/http"
	"sync"
	"time"
)

type Context struct {
	request        *http.Request
	responseWriter http.ResponseWriter

	writerMux    *sync.Mutex
	hasResponded bool

	ctx context.Context

	handlers []ControllerHandler
	index    int

	params map[string]string
}

func NewContext(responseWriter http.ResponseWriter, request *http.Request) *Context {
	ctx := &Context{
		request:        request,
		responseWriter: responseWriter,
		ctx:            nil,
		writerMux:      new(sync.Mutex),
		index:          -1,
	}

	return ctx
}

func (ctx *Context) GetRequest() *http.Request {
	return ctx.request
}

func (ctx *Context) WriterMux() *sync.Mutex {
	return ctx.writerMux
}

func (ctx *Context) HasResponded() bool {
	return ctx.hasResponded
}

// BaseContext Context Method
func (ctx *Context) BaseContext() context.Context {
	return ctx.request.Context()
}

func (ctx *Context) DeadLine() (deadLine time.Time, ok bool) {
	return ctx.BaseContext().Deadline()
}

func (ctx *Context) Err() error {
	return ctx.BaseContext().Err()
}

func (ctx *Context) Value(key interface{}) interface{} {
	return ctx.BaseContext().Value(key)
}

func (ctx *Context) JsonWithStatusCode(statusCode int, obj interface{}) error {
	ctx.WriterMux().Lock()
	defer ctx.WriterMux().Unlock()
	if ctx.HasResponded() {
		return nil
	}

	ctx.responseWriter.Header().Set("Content-Type", "application/json")
	ctx.responseWriter.WriteHeader(statusCode)
	byt, err := json.Marshal(obj)
	if err != nil {
		ctx.responseWriter.WriteHeader(500)
		return err
	}

	_, err = ctx.responseWriter.Write(byt)
	ctx.hasResponded = true
	return err
}

func (ctx *Context) SetHandlers(handlers []ControllerHandler) {
	ctx.handlers = handlers
}

func (ctx *Context) Next() error {
	ctx.index++
	if ctx.index < len(ctx.handlers) {
		if err := ctx.handlers[ctx.index](ctx); err != nil {
			return err
		}
	}

	return nil
}
