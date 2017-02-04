package github

import (
	hub "github.com/google/go-github/github"
)

type Issue struct {
	Iss      *hub.Issue
	Comments []*hub.IssueComment
}
