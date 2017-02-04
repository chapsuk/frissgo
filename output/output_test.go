package output_test

import (
	"testing"

	"github.com/chapsuk/frissgo/config"
	"github.com/chapsuk/frissgo/judge"
	"github.com/chapsuk/frissgo/output"
)

func TestOutputWrite(t *testing.T) {

	cfg := &config.Output{
		Format: config.OutputFormatText,
		Target: config.OutputTargetStdout,
	}

	o, err := output.New(cfg)
	if err != nil {
		t.Fatal(err)
	}

	top := []*judge.Category{}

	err = o.Write(top)
	if err != nil {
		t.Fatal(err)
	}
}
