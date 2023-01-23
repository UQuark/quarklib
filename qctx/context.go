package qctx

import (
	"context"
)

type Context struct {
	context.Context
}

func C() Context {
	return Context{context.Background()}
}

func (ctx *Context) Put(key, value any) {
	ctx.Context = context.WithValue(ctx.Context, key, value)
}

func (ctx *Context) Get(key any) (any, error) {
	i := ctx.Value(key)
	if i == nil {
		return nil, ErrKeyAbsent.Format(key)
	}
	return i, nil
}

func From(ctx context.Context) Context {
	switch ctx.(type) {
	case Context:
		return ctx.(Context)
	case *Context:
		return *ctx.(*Context)
	default:
		return Context{ctx}
	}
}
