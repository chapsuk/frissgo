package main

import (
	"errors"
	"flag"

	"github.com/chapsuk/frissgo/config"
	"github.com/chapsuk/frissgo/github"
	"github.com/chapsuk/frissgo/judge"
	"github.com/chapsuk/frissgo/output"
)

const (
	MODE_TOP      = "top"
	MODE_CATEGORY = "category"
)

var (
	c = flag.String("cfg", "go.yml", "config file name")
	m = flag.String("mode", "top", "output mode: top or category")
)

func main() {
	flag.Parse()

	cfg, err := config.LoadFile(*c)
	handleError(err)

	g, err := github.New(cfg.Github)
	handleError(err)

	o, err := output.New(cfg.Output)
	handleError(err)

	j := judge.New(cfg.Strategy, g)

	var t []*judge.Category
	switch *m {
	case MODE_CATEGORY:
		t, err = j.GetCategoriesTop()
		handleError(err)
	case MODE_TOP:
		c, err := j.GetBestOfTheBest()
		handleError(err)
		t = append(t, c)
	default:
		handleError(errors.New("unsupported mode"))
	}

	err = o.WriteCategories(t)
	handleError(err)
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
