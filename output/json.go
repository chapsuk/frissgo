package output

import "github.com/chapsuk/gofriss/judge"

type JSONFormatter struct{}

func (jf *JSONFormatter) Format(t *judge.Top) []byte {
	return nil
}
