package qmw_test

import (
	"context"
	"github.com/UQuark0/quarklib/qctx"
	"github.com/UQuark0/quarklib/qerr"
	"github.com/UQuark0/quarklib/qmw"
	"io"
	"net/http"
	"testing"
)

var ErrTest = qerr.NewError("testcode", "pattern")

func handle(ctx context.Context, r *http.Request) (qmw.Response, error) {
	resp := struct {
		Test string
	}{
		Test: "test",
	}
	switch r.Method {
	case http.MethodGet:
		return qmw.R().Header("Test", "test").Json(resp), nil
	case http.MethodPost:
		return nil, nil
	case http.MethodDelete:
		return nil, io.ErrClosedPipe
	case http.MethodPatch:
		return nil, ErrTest.New()
	case http.MethodPut:
		return nil, qmw.E(http.StatusBadRequest, ErrTest.New())
	}
	return nil, nil
}

func TestMw(t *testing.T) {
	ctx := qctx.C()

	http.HandleFunc("/", qmw.H(ctx, handle))
	http.ListenAndServe(":8080", nil)
}
