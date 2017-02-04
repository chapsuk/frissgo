package output

import "github.com/chapsuk/gofriss/judge"

type JSONFormatter struct{}

func (jf *JSONFormatter) Format(cats []*judge.Category) []byte {
	return []byte("json")
}
