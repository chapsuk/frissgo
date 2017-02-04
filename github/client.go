package github

import "github.com/chapsuk/gofriss/config"

type Client struct {
}

func New(cfg config.Github) (*Client, error) {
	return &Client{}, nil
}
