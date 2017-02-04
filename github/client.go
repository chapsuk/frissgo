package github

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/chapsuk/gofriss/config"
	hub "github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type Client struct {
	cfg    *config.Github
	github *hub.Client
}

func New(cfg *config.Github) (*Client, error) {
	var tc *http.Client
	if cfg.AccessToken != "" {
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: cfg.AccessToken},
		)
		tc = oauth2.NewClient(oauth2.NoContext, ts)
	}
	return &Client{
		github: hub.NewClient(tc),
		cfg:    cfg,
	}, nil
}

func (c *Client) GetIssues(count, page int) ([]*Issue, error) {
	opt, err := c.buildOpt()
	if err != nil {
		return nil, err
	}
	opt.PerPage = count
	opt.Page = page

	iss, res, err := c.github.Issues.ListByRepo(c.cfg.Owner, c.cfg.Repository, opt)
	log.Printf("github limits response: %+v", res)
	if err != nil {
		return nil, err
	}

	result := make([]*Issue, 0, len(iss))
	for _, i := range iss {

		var cmmnts []*hub.IssueComment
		if *i.Comments > 0 {
			opts := new(hub.IssueListCommentsOptions)
			opts.Since = opt.Since
			opts.PerPage = *i.Comments

			cmmnts, res, err := c.github.Issues.ListComments(c.cfg.Owner, c.cfg.Repository, *i.Number, opts)
			if err != nil {
				return nil, err
			}
			log.Printf("github limits response: %+v", res)
			log.Printf("gotten %d comments for issue %d", len(cmmnts), i.Number)
		}

		result = append(result, &Issue{
			Iss:      i,
			Comments: cmmnts,
		})
	}
	return result, nil
}

func (c *Client) buildOpt() (*hub.IssueListByRepoOptions, error) {
	opt := new(hub.IssueListByRepoOptions)
	for k, v := range c.cfg.Filters {
		switch k {
		case config.GithubFilterMilestone:
			opt.Milestone = v
		case config.GithubFilterState:
			opt.State = v
		case config.GithubFilterAssignee:
			opt.Assignee = v
		case config.GithubFilterCreator:
			opt.Creator = v
		case config.GithubFilterMentioned:
			opt.Mentioned = v
		case config.GithubFilterLabels:
			opt.Labels = strings.Split(v, ",")
		case config.GithubFilterSort:
			opt.Sort = v
		case config.GithubFilterDirection:
			opt.Direction = v
		case config.GithubFilterPeriod:
			t, err := time.ParseDuration(v)
			if err != nil {
				return nil, err
			}
			opt.Since = time.Now().Add(-1 * t)
		}
	}
	return opt, nil
}
