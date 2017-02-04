package output

import (
	"bytes"
	"fmt"

	"github.com/chapsuk/frissgo/judge"
)

type TextFormatter struct {
}

func (tf *TextFormatter) Format(cats []*judge.Category) []byte {
	var buf bytes.Buffer
	for _, c := range cats {
		buf.WriteRune('\n')
		buf.WriteString(fmt.Sprintf("[ %s ]", c.Name))
		buf.WriteRune('\n')
		buf.WriteString("=====================")
		buf.WriteRune('\n')
		for _, i := range c.GetIssues() {
			buf.WriteString(*i.Iss.HTMLURL)
			buf.WriteRune('\n')
		}
		buf.WriteString("=====================")
		buf.WriteRune('\n')
	}
	return buf.Bytes()
}
