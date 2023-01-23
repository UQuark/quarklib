package qctx

import "github.com/UQuark0/quarklib/qerr"

var ErrKeyAbsent = qerr.NewError("qctx.key_absent", "key '%s' absent in the context")
