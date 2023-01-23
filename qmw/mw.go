package qmw

import (
	"context"
	"net/http"
)

type HandlerFunc func(ctx context.Context, req *http.Request) (Response, error)

func (h HandlerFunc) ServeHTTP(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	resp, err := h(ctx, r)
	if err == nil {
		if resp == nil {
			resp = R()
		}
	} else {
		resp = toResponse(err)
	}
	return resp.Write(w)
}

func H(ctx context.Context, h HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_ = h.ServeHTTP(ctx, w, r)
	}
}
