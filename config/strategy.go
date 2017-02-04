package config

type Strategy struct {
	PriorAuthors []string   `yaml:"prior_authors"`
	Categories   []Category `yaml:"categories"`
}

type Category struct {
	Name     string   `yaml:"name"`
	Issues   Issues   `yaml:"issues"`
	Comments Comments `yaml:"comments"`
}

type Issues struct {
	Author    int `yaml:"author"`
	Activity  int `yaml:"activity"`
	Reactions int `yaml:"reaction"`
}

type Comments struct {
	Author    int `yaml:"author"`
	Activity  int `yaml:"activity"`
	Reactions int `yaml:"reaction"`
}
