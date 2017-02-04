package output

import (
	"encoding/json"

	"github.com/chapsuk/frissgo/judge"
)

type JSONFormatter struct{}

type Category struct {
	Name   string  `json:"category"`
	Issues []Issue `json:"issues"`
}

type Issue struct {
	URL string `json:"string"`
}

func (jf *JSONFormatter) Format(cats []*judge.Category) []byte {
	var result []Category
	for _, c := range cats {
		issues := c.GetIssues()
		cr := Category{
			Name:   c.Name,
			Issues: make([]Issue, 0, len(issues)),
		}
		for _, i := range issues {
			cr.Issues = append(cr.Issues, Issue{
				URL: *i.Iss.HTMLURL,
			})
		}
		result = append(result, cr)
	}
	res, _ := json.Marshal(result)
	return res
}
