package output

import "github.com/chapsuk/gofriss/judge"

type TextFormatter struct {
}

func (tf *TextFormatter) Format(t *judge.Top) []byte {
	return nil
}
