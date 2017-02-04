package main

import (
	"flag"

	"github.com/chapsuk/gofriss/config"
	"github.com/chapsuk/gofriss/github"
	"github.com/chapsuk/gofriss/judge"
	"github.com/chapsuk/gofriss/output"
)

var (
	c = flag.String("cfg", "config.yml", "config file name")
)

func main() {
	flag.Parse()

	cfg, err := config.Load(*c)
	handleError(err)

	g, err := github.New(cfg.Github)
	handleError(err)

	o, err := output.New(cfg.Output)
	handleError(err)

	j, err := judge.New(cfg.Strategy, g)
	handleError(err)

	t, err := j.GetTop()
	handleError(err)

	_, err = o.Flush(t)
	handleError(err)
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
