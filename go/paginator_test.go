package paginator

import (
	"testing"
)

func TestLimit(t *testing.T) {
	var tests = []struct {
		paginator *Paginator
		expected  uint64
	}{
		{
			paginator: &Paginator{
				Limit: 10,
			},
			expected: 10,
		},
		{
			paginator: &Paginator{
				Limit: 20,
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
				Page:  1,
				Limit: 10,
			},
			expected: 0,
		},
		{
			paginator: &Paginator{
				Page:  2,
				Limit: 10,
			},
			expected: 10,
		},
		{
			paginator: &Paginator{
				Page:  2,
				Limit: 50,
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

func TestTotal(t *testing.T) {
	var tests = []struct {
		paginator *Paginator
		total     uint64
		expected  uint64
	}{
		{
			paginator: &Paginator{},
			total:     10,
			expected:  10,
		},
		{
			paginator: &Paginator{
				Total: 20,
			},
			total:    20,
			expected: 20,
		},
	}

	for i, test := range tests {
		test.paginator.Total = test.total
		var total = test.paginator.GetTotal()

		if total != test.expected {
			t.Fatalf("test %d failed; actual: %d; expected: %d", i, total, test.expected)
		}
	}
}
