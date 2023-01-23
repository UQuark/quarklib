package qmw

import (
	"bytes"
	"encoding/json"
	"github.com/UQuark0/quarklib/qerr"
	"io"
	"net/http"
)

type Response interface {
	Status(status int) Response
	Json(v any) Response
	Bytes(b []byte) Response
	Plain(s string) Response
	Reader(rd io.Reader) Response
	Raw(b []byte) Response
	Header(key, value string) Response
	Write(w http.ResponseWriter) error
}

type response struct {
	status  int
	body    io.Reader
	headers http.Header
}

func R() Response {
	return response{status: http.StatusOK, headers: http.Header{}}
}

func (r response) Status(status int) Response {
	r.status = status
	return r
}

func (r response) Json(v any) Response {
	buf, _ := json.Marshal(v)
	return r.Header("Content-Type", "application/json").Raw(buf)
}

func (r response) Bytes(b []byte) Response {
	return r.Header("Content-Type", "application/octet-stream").Raw(b)
}

func (r response) Plain(s string) Response {
	return r.Header("Content-Type", "text/plain").Raw([]byte(s))
}

func (r response) Reader(rd io.Reader) Response {
	r.body = rd
	return r
}

func (r response) Raw(b []byte) Response {
	r.body = bytes.NewReader(b)
	return r
}

func (r response) Header(key, value string) Response {
	r.headers.Set(key, value)
	return r
}

func (r response) Write(w http.ResponseWriter) error {
	for k, v := range r.headers {
		w.Header()[k] = v
	}
	w.WriteHeader(r.status)
	_, err := io.Copy(w, r.body)
	if err != nil {
		return ErrWriteResponse.Cause(err)
	}
	return nil
}

func toResponse(e error) Response {
	switch e.(type) {
	case statusError:
		return toResponse(e.(statusError).e).Status(e.(statusError).status)
	case qerr.Error:
		return R().Status(http.StatusInternalServerError).Json(ErrBody{
			Code:    e.(qerr.Error).GetCode(),
			Message: e.Error(),
		})
	default:
		return R().Status(http.StatusInternalServerError).Json(ErrBody{
			Code:    "unknown",
			Message: e.Error(),
		})
	}
}
