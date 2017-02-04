package config

type Strategy struct {
	PerPage      int        `yaml:"per_page"`
	PriorAuthors []string   `yaml:"prior_authors"`
	Categories   []Category `yaml:"categories"`
}

type Category struct {
	Name     string   `yaml:"name"`
	Size     int      `yaml:"size"`
	Issues   Issues   `yaml:"issues"`
	Comments Comments `yaml:"comments"`
}

type Issues struct {
	Author    int      `yaml:"author"`
	Activity  int      `yaml:"activity"`
	Reactions Reaction `yaml:"reaction"`
}

type Comments struct {
	Author    int      `yaml:"author"`
	Activity  int      `yaml:"activity"`
	Reactions Reaction `yaml:"reaction"`
}

type Reaction struct {
	Plus     int `yaml:"plus"`
	Heart    int `yaml:"heart"`
	Laugh    int `yaml:"laugh"`
	Hooray   int `yaml:"hooray"`
	Minus    int `yaml:"minus"`
	Confused int `yaml:"confused"`
	Total    int `yaml:"total"`
}

func (r *Reaction) HasCoeff() bool {
	return r.Total != 0 ||
		r.Plus != 0 ||
		r.Minus != 0 ||
		r.Heart != 0 ||
		r.Laugh != 0 ||
		r.Hooray != 0 ||
		r.Confused != 0
}

func (s *Strategy) IsPriorAuthor(author string) bool {
	for _, name := range s.PriorAuthors {
		if name == author {
			return true
		}
	}
	return false
}
