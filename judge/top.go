package judge

import (
	"sort"
	"sync"

	"github.com/chapsuk/frissgo/github"
)

type top struct {
	mu        sync.Mutex
	items     topItems
	size      int
	minWeight int
	sorted    bool
}

type topItem struct {
	weight int
	issue  *github.Issue
}

func newTopChart(size int) *top {
	return &top{
		items: make([]*topItem, 0, size),
		size:  size,
	}
}

type topItems []*topItem

func (t topItems) Len() int           { return len(t) }
func (t topItems) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }
func (t topItems) Less(i, j int) bool { return t[i].weight > t[j].weight }

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

	if !t.sorted {
		sort.Sort(t.items)
		t.sorted = true
	}

	t.items[t.size-1] = i
	t.minWeight = i.weight

	return true
}
