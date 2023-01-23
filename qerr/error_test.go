package qerr_test

import (
	"github.com/UQuark0/quarklib/qerr"
	"io"
	"testing"
)

var ErrTest = qerr.NewError("qerr.test", "test %s")

func TestError(t *testing.T) {
	err := ErrTest.Cause(io.ErrClosedPipe).Format("abc")
	if err.GetCode() != "qerr.test" {
		t.Fatal("invalid code")
	}
	if err.Error() != "test abc\ncaused by: io: read/write on closed pipe" {
		t.Fatal("invalid message")
	}
}
