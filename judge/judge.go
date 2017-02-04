package judge

import (
	"github.com/chapsuk/gofriss/config"
	"github.com/chapsuk/gofriss/github"
)

type Category struct {
	Name string
	Top  *Top
}

type Top struct {
}

type Judge struct {
}

func New(cfg *config.Strategy, git *github.Client) (*Judge, error) {
	return nil, nil
}

func (j *Judge) GetCategoriesTop() ([]*Category, error) {
	return []*Category{
		&Category{},
	}, nil
}
