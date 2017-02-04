package config_test

import (
	"testing"

	"github.com/chapsuk/frissgo/config"
)

func TestLoad(t *testing.T) {
	bcfg := []byte(`
github:
strategy:
output:
  format: text
  target: stdout
  file_name:
`)

	cfg, err := config.Load(bcfg)
	if err != nil {
		t.Fatalf("parse config error: %s", err)
	}

	if cfg.Output.Format != "text" {
		t.Fatalf("incorrect parse config, gotten: %+v", cfg)
	}

	if cfg.Output.Target != "stdout" {
		t.Fatalf("incorrect parse config, gotten: %+v", cfg)
	}

	if cfg.Output.FileName != "" {
		t.Fatalf("incorrect parse config, gotten: %+v", cfg)
	}
}
