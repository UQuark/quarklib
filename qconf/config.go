package qconf

import (
	"context"
	"github.com/UQuark0/quarklib/qctx"
	"io"
	"os"
	"strings"
)

type ParseFunc = func(r io.Reader) (map[string]any, error)

func Init(ctx context.Context) (context.Context, error) {
	qCtx := qctx.From(ctx)
	return qCtx, QInit(&qCtx)
}

func QInit(ctx *qctx.Context) error {
	switch _, err := ctx.Get(Reader); err {
	default:
		an, err := ctx.Get(Filename)
		filename := an.(string)

		if err == nil {
			f, err := os.Open(filename)
			if err != nil {
				return ErrOpenFailed.Format(filename).Cause(err)
			}
			defer f.Close()

			ctx.Put(Reader, io.Reader(f))
		}
		fallthrough
	case nil:
		return readerInit(ctx)
	}
}

func readerInit(ctx *qctx.Context) error {
	an, err := ctx.Get(Parser)
	if err != nil {
		return err
	}
	parse := an.(ParseFunc)

	an, err = ctx.Get(Reader)
	if err != nil {
		return err
	}
	r := an.(io.Reader)

	prefix := ""
	an, err = ctx.Get(Prefix)
	if err == nil {
		prefix = an.(string)
	}

	entries, err := parse(r)
	if err != nil {
		return ErrParseFailed.Cause(err)
	}

	for k, v := range entries {
		if strings.HasPrefix(k, prefix) {
			ctx.Put(k, v)
		}
	}

	return nil
}
