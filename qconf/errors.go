package qconf

import "github.com/UQuark0/quarklib/qerr"

var (
	ErrOpenFailed  = qerr.NewError("qconf.open_failed", "failed to open config file '%s'")
	ErrParseFailed = qerr.NewError("qconf.parse_failed", "failed to parse config")
)
