package judge

import (
	"sort"

	"github.com/chapsuk/gofriss/config"
	"github.com/chapsuk/gofriss/github"
)

type byWeight []*topItem

func (w byWeight) Len() int           { return len(w) }
func (w byWeight) Swap(i, j int)      { w[i], w[j] = w[j], w[i] }
func (w byWeight) Less(i, j int) bool { return w[i].weight > w[j].weight }

type estimatorFunc func(iss *github.Issue) *topItem

type Category struct {
	Name      string
	estimator estimatorFunc
	top       *top
}

func newCategory(strategy *config.Strategy, cfg config.Category) *Category {
	return &Category{
		Name:      cfg.Name,
		top:       newTopChart(cfg.Size),
		estimator: createEstimatorFunc(strategy, cfg),
	}
}

func (c *Category) GetIssues() []*github.Issue {
	result := make([]*github.Issue, 0, len(c.top.items))
	sort.Sort(byWeight(c.top.items))
	for _, t := range c.top.items {
		result = append(result, t.issue)
	}
	return result
}

func createEstimatorFunc(strategy *config.Strategy, cfg config.Category) estimatorFunc {
	return func(iss *github.Issue) *topItem {
		item := &topItem{
			issue: iss,
		}

		if cfg.Issues.Activity != 0 {
			item.weight += cfg.Issues.Activity * *iss.Iss.Comments
		}

		if cfg.Issues.Author != 0 {
			if strategy.IsPriorAuthor(*iss.Iss.User.Login) {
				item.weight += cfg.Issues.Author
			}
		}

		if cfg.Issues.Reactions.HasCoeff() {
			item.weight += cfg.Issues.Reactions.Total * *iss.Iss.Reactions.TotalCount
			item.weight += cfg.Issues.Reactions.Plus * *iss.Iss.Reactions.PlusOne
			item.weight += cfg.Issues.Reactions.Laugh * *iss.Iss.Reactions.Laugh
			item.weight += cfg.Issues.Reactions.Hooray * *iss.Iss.Reactions.Hooray
			item.weight += cfg.Issues.Reactions.Heart * *iss.Iss.Reactions.Heart
			item.weight += cfg.Issues.Reactions.Minus * *iss.Iss.Reactions.MinusOne
			item.weight += cfg.Issues.Reactions.Confused * *iss.Iss.Reactions.Confused
		}

		if cfg.Comments.Activity != 0 {
			item.weight += cfg.Comments.Activity * len(iss.Comments)
		}

		if cfg.Comments.Author != 0 || cfg.Comments.Reactions.HasCoeff() {
			for _, comm := range iss.Comments {
				if cfg.Comments.Author != 0 {
					if strategy.IsPriorAuthor(*comm.User.Login) {
						item.weight += cfg.Issues.Author
					}
				}

				if cfg.Comments.Reactions.HasCoeff() {
					item.weight += cfg.Comments.Reactions.Total * *comm.Reactions.TotalCount
					item.weight += cfg.Comments.Reactions.Plus * *comm.Reactions.PlusOne
					item.weight += cfg.Comments.Reactions.Laugh * *comm.Reactions.Laugh
					item.weight += cfg.Comments.Reactions.Hooray * *comm.Reactions.Hooray
					item.weight += cfg.Comments.Reactions.Heart * *comm.Reactions.Heart
					item.weight += cfg.Comments.Reactions.Minus * *comm.Reactions.MinusOne
					item.weight += cfg.Comments.Reactions.Confused * *comm.Reactions.Confused
				}
			}
		}

		return item
	}
}
