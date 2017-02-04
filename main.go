package main

import (
	"flag"

	"github.com/chapsuk/frissgo/config"
	"github.com/chapsuk/frissgo/github"
	"github.com/chapsuk/frissgo/judge"
	"github.com/chapsuk/frissgo/output"
)

var (
	c = flag.String("cfg", "go.yml", "config file name")
)

func main() {
	flag.Parse()

	cfg, err := config.LoadFile(*c)
	handleError(err)

	g, err := github.New(cfg.Github)
	handleError(err)

	o, err := output.New(cfg.Output)
	handleError(err)

	t, err := judge.New(cfg.Strategy, g).GetCategoriesTop()
	handleError(err)

	err = o.Write(t)
	handleError(err)
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
