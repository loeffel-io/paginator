package paginator

import (
	"sync"
	"testing"
)

func TestLimit(t *testing.T) {
	var tests = []struct {
		paginator *Paginator
		expected  uint64
	}{
		{
			paginator: &Paginator{
				Limit:   10,
				RWMutex: new(sync.RWMutex),
			},
			expected: 10,
		},
		{
			paginator: &Paginator{
				Limit:   20,
				RWMutex: new(sync.RWMutex),
			},
			expected: 20,
		},
	}

	for i, test := range tests {
		var limit = test.paginator.GetLimit()

		if limit != test.expected {
			t.Fatalf("test %d failed; actual: %d; expected: %d", i, limit, test.expected)
		}
	}
}

func TestOffset(t *testing.T) {
	var tests = []struct {
		paginator *Paginator
		expected  uint64
	}{
		{
			paginator: &Paginator{
				Page:    1,
				Limit:   10,
				RWMutex: new(sync.RWMutex),
			},
			expected: 0,
		},
		{
			paginator: &Paginator{
				Page:    2,
				Limit:   10,
				RWMutex: new(sync.RWMutex),
			},
			expected: 10,
		},
		{
			paginator: &Paginator{
				Page:    2,
				Limit:   50,
				RWMutex: new(sync.RWMutex),
			},
			expected: 50,
		},
	}

	for i, test := range tests {
		var offset = test.paginator.Offset()

		if offset != test.expected {
			t.Fatalf("test %d failed; actual: %d; expected: %d", i, offset, test.expected)
		}
	}
}

func TestSetTotal(t *testing.T) {
	var tests = []struct {
		paginator *Paginator
		total     uint64
		expected  uint64
	}{
		{
			paginator: &Paginator{
				RWMutex: new(sync.RWMutex),
			},
			total:    10,
			expected: 10,
		},
		{
			paginator: &Paginator{
				Total:   20,
				RWMutex: new(sync.RWMutex),
			},
			total:    20,
			expected: 20,
		},
	}

	for i, test := range tests {
		test.paginator.SetTotal(test.total)
		var total = test.paginator.GetTotal()

		if total != test.expected {
			t.Fatalf("test %d failed; actual: %d; expected: %d", i, total, test.expected)
		}
	}
}
