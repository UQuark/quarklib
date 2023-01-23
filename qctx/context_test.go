package qctx_test

import (
	"github.com/UQuark0/quarklib/qctx"
	"testing"
)

func TestValues(t *testing.T) {
	const (
		cKey        = "qctx.test"
		cInvalidKey = "qctx.invalid"
		cVal        = "test"
	)

	ctx := qctx.C()

	ctx.Put(cKey, cVal)

	v, err := ctx.Get(cKey)
	if err != nil {
		t.Fatal(err)
	}

	if v.(string) != cVal {
		t.Fatal("invalid value")
	}

	_, err = ctx.Get(cInvalidKey)
	if err == nil {
		t.Fatal("error expected")
	}
}

func TestPass(t *testing.T) {
	const (
		cOutKey = "qctx.test.out"
		cInKey  = "qctx.test.in"
		cVal    = "test"
	)

	// copy

	ctx := qctx.C()
	ctx.Put(cOutKey, cVal)

	func(ctx qctx.Context) {
		if _, err := ctx.Get(cOutKey); err != nil {
			t.Fatal(err)
		}
		ctx.Put(cInKey, cVal)
	}(ctx)

	if _, err := ctx.Get(cInKey); err == nil {
		t.Fatal("error expected")
	}

	// ref

	ctx = qctx.C()
	ctx.Put(cOutKey, cVal)

	func(ctx *qctx.Context) {
		if _, err := ctx.Get(cOutKey); err != nil {
			t.Fatal(err)
		}
		ctx.Put(cInKey, cVal)
	}(&ctx)

	if _, err := ctx.Get(cInKey); err != nil {
		t.Fatal(err)
	}
}
