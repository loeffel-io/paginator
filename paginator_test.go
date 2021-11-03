package paginator

import (
	"sync"
	"testing"
)

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

		if test.paginator.Offset() != test.expected {
			t.Fatalf("test %d failed; actual: %d; expected: %d", i, offset, test.expected)
		}
	}
}
