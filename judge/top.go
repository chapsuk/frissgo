package judge

import (
	"sync"

	"github.com/chapsuk/frissgo/github"
)

type top struct {
	minWeight int
	mu        sync.Mutex
	items     []*topItem
}

type topItem struct {
	weight int
	issue  *github.Issue
}

func newTopChart(size int) *top {
	return &top{
		items: make([]*topItem, 0, size),
	}
}

func (t *top) add(i *topItem) bool {
	if i.weight <= 0 {
		return false
	}

	t.mu.Lock()
	defer t.mu.Unlock()

	if len(t.items) != cap(t.items) {
		t.items = append(t.items, i)
		if i.weight < t.minWeight {
			t.minWeight = i.weight
		}
		return true
	}

	if i.weight <= t.minWeight {
		return false
	}

	n := make([]*topItem, 0, cap(t.items))
	min := i.weight
	for _, item := range t.items {
		if item.weight > i.weight {
			n = append(n, item)
			if item.weight < min {
				min = item.weight
			}
		}
	}
	t.items = append(n, i)
	t.minWeight = min

	return true
}
