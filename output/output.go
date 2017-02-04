package output

import (
	"errors"
	"os"

	"github.com/chapsuk/gofriss/config"
	"github.com/chapsuk/gofriss/judge"
)

var (
	ErrInvalidTarget = errors.New("incorrect output target value")
)

type Formatter interface {
	Format(t *judge.Top) []byte
}

type Output struct {
	formatter Formatter
	cfg       *config.Output
}

func New(cfg *config.Output) (*Output, error) {
	var f Formatter
	switch cfg.Format {
	case config.OutputFormatText:
		f = new(TextFormatter)
	case config.OutputFormatJSON:
		f = new(JSONFormatter)
	default:
		return nil, errors.New("incorrect output format value")
	}

	switch cfg.Target {
	case config.OutputTargetStdout:
	case config.OutputTargetFile:
		if _, err := os.Stat(cfg.FileName); err != nil {
			return nil, err
		}
	default:
		return nil, ErrInvalidTarget
	}

	return &Output{
		cfg:       cfg,
		formatter: f,
	}, nil
}

func (o *Output) Flush(t *judge.Top) (int, error) {
	if o.cfg.Target == config.OutputTargetStdout {
		return os.Stdout.Write(o.formatter.Format(t))
	}

	if o.cfg.Target == config.OutputTargetFile {
		f, err := os.Open(o.cfg.FileName)
		if err != nil {
			return 0, err
		}
		defer f.Close()
		return f.Write(o.formatter.Format(t))
	}

	return 0, ErrInvalidTarget
}
