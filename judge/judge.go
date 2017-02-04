package judge

import (
	"github.com/chapsuk/gofriss/config"
	"github.com/chapsuk/gofriss/github"
)

type Top struct {
}

type Judge struct {
}

func New(cfg config.Strategy, git *github.Client) (*Judge, error) {
	return nil, nil
}

func (j *Judge) GetTop() (*Top, error) {
	return &Top{}, nil
}
