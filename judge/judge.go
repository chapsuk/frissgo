package judge

import (
	"sync"

	"github.com/chapsuk/frissgo/config"
	"github.com/chapsuk/frissgo/github"
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

// GetBestOfTheBest recalculate weight for top items from categories tops
// seeing categories coefficients (config.Strategy.Category.Coeff)
func (j *Judge) GetBestOfTheBest() (*Category, error) {
	cats, err := j.GetCategoriesTop()
	if err != nil {
		return nil, err
	}

	cbotb := &Category{
		Name: "BestOfTheBest",
		top:  newTopChart(10),
	}

	f := func(iss *github.Issue, cats []*Category) *topItem {
		item := &topItem{issue: iss}
		for _, cat := range cats {
			i := cat.estimator(iss)
			item.weight += cat.Coeff * i.weight
		}
		// log.Printf("isss: %s weight: %d", item.issue.Iss.GetURL(), item.weight)
		return item
	}

	t := make(map[int]struct{}, 1024)
	for _, cat := range cats {
		for _, i := range cat.top.items {
			if _, ok := t[i.issue.Iss.GetID()]; ok {
				continue
			}
			t[i.issue.Iss.GetID()] = struct{}{}

			j.wg.Add(1)
			go func(iss *github.Issue) {
				cbotb.top.add(f(iss, cats))
				j.wg.Done()
			}(i.issue)
		}
	}
	j.wg.Wait()

	return cbotb, nil
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
				// t := cat.estimator(iss)
				// log.Printf("Try add to category %s issue: %s weight: %d", cat.Name, *iss.Iss.HTMLURL, t.weight)
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
