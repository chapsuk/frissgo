package judge

import (
	"sync"

	"github.com/chapsuk/gofriss/config"
	"github.com/chapsuk/gofriss/github"
)

type Judge struct {
	cfg *config.Strategy
	hub *github.Client
	wg  sync.WaitGroup
}

func New(cfg *config.Strategy, gh *github.Client) *Judge {
	return &Judge{
		cfg: cfg,
		hub: gh,
	}
}

func (j *Judge) GetCategoriesTop() ([]*Category, error) {
	issues, err := j.loadIssues()
	if err != nil {
		return nil, err
	}
	categories := j.createCategories()

	j.wg.Add(len(issues))
	for _, iss := range issues {
		go func(iss *github.Issue) {
			for _, cat := range categories {
				cat.top.add(cat.estimator(iss))
			}
			j.wg.Done()
		}(iss)
	}
	j.wg.Wait()

	return categories, nil
}

func (j *Judge) loadIssues() ([]*github.Issue, error) {
	var result []*github.Issue
	var page int
	for {
		page++
		iss, err := j.hub.GetIssues(j.cfg.PerPage, page)
		if err != nil {
			return nil, err
		}
		result = append(result, iss...)
		if len(iss) < j.cfg.PerPage {
			break
		}
	}
	return result, nil
}

func (j *Judge) createCategories() []*Category {
	var result []*Category
	for _, catcfg := range j.cfg.Categories {
		result = append(result, newCategory(j.cfg, catcfg))
	}
	return result
}
