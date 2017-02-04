package github_test

import (
	"log"
	"testing"

	"github.com/chapsuk/frissgo/config"
	"github.com/chapsuk/frissgo/github"
)

func TestGetIssues(t *testing.T) {
	cfg, err := config.LoadFile("./../go.yml")
	if err != nil {
		t.Fatalf("load config error: %s", err)
	}
	log.Printf("config: %+v", cfg.Github)
	gh, err := github.New(cfg.Github)
	if err != nil {
		t.Fatalf("create client error: %s", err)
	}

	iss, err := gh.GetIssues(1, 1)
	if err != nil {
		t.Fatalf("get issues error: %s", err)
	}

	log.Printf("gotten %d issues", len(iss))
}
