package config

import "errors"

const (
	GithubFilterMilestone = "milestone"
	GithubFilterState     = "state"
	GithubFilterAssignee  = "assignee"
	GithubFilterCreator   = "creator"
	GithubFilterMentioned = "mentioned"
	GithubFilterLabels    = "lables"
	GithubFilterSort      = "sort"
	GithubFilterDirection = "direction"
	GithubFilterPeriod    = "period"
)

type Filters map[string]string

type Github struct {
	AccessToken string  `yaml:"access_token"`
	Owner       string  `yaml:"owner"`
	Repository  string  `yaml:"repo"`
	Filters     Filters `yaml:"filters"`
}

func (f *Filters) Check() error {
	for k := range *f {
		switch k {
		case GithubFilterMilestone:
		case GithubFilterState:
		case GithubFilterAssignee:
		case GithubFilterCreator:
		case GithubFilterMentioned:
		case GithubFilterLabels:
		case GithubFilterSort:
		case GithubFilterDirection:
		case GithubFilterPeriod:
		default:
			return errors.New("unsupported filter key " + k)
		}
	}
	return nil
}
