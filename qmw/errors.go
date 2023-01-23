package qmw

import "github.com/UQuark0/quarklib/qerr"

var (
	ErrWriteResponse = qerr.NewError("qmw.write_response", "failed to write response")
)
