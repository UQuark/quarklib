package qconf_test

import (
	"errors"
	"github.com/UQuark0/quarklib/qconf"
	"github.com/UQuark0/quarklib/qctx"
	"io"
	"strings"
	"testing"
)

func TestQInit(t *testing.T) {
	ctx := qctx.C()
	ctx.Put(qconf.Filename, "qconf.test.kev")
	ctx.Put(qconf.Prefix, "qconf.test.")
	ctx.Put(qconf.Parser, func(r io.Reader) (map[string]any, error) {
		bytes, err := io.ReadAll(r)
		if err != nil {
			return nil, err
		}
		res := make(map[string]any)
		for _, row := range strings.Split(string(bytes), "\n") {
			parts := strings.Split(row, "=")
			if len(parts) != 2 {
				return nil, errors.New("invalid input")
			}
			res[parts[0]] = parts[1]
		}
		return res, nil
	})
	if err := qconf.QInit(&ctx); err != nil {
		t.Fatal(err)
	}

	if _, err := ctx.Get("qconf.garbage"); err == nil {
		t.Fatal("error expected")
	}

	v, err := ctx.Get("qconf.test.key")
	if err != nil {
		t.Fatal(err)
	}
	if v.(string) != "value1" {
		t.Fatal("invalid value")
	}

	v, err = ctx.Get("qconf.test.prefix.a")
	if err != nil {
		t.Fatal(err)
	}
	if v.(string) != "value2" {
		t.Fatal("invalid value")
	}

	v, err = ctx.Get("qconf.test.prefix.b")
	if err != nil {
		t.Fatal(err)
	}
	if v.(string) != "value3" {
		t.Fatal("invalid value")
	}
}
