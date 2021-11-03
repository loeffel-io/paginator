package paginator

import (
	"sync"
)

type Paginator struct {
	Page  uint64 `gorm:"-" json:"page"`
	Limit uint64 `gorm:"-" json:"limit"`
	Total uint64 `gorm:"total" json:"total"`
	*sync.RWMutex
}

func (paginator Paginator) GetLimit() uint64 {
	paginator.RLock()
	defer paginator.RUnlock()

	return paginator.Limit
}

func (paginator *Paginator) Offset() uint64 {
	paginator.RLock()
	defer paginator.RUnlock()

	if paginator.Page == 1 {
		return 0
	}

	return (paginator.Page - 1) * paginator.Limit
}

func (paginator *Paginator) SetTotal(total uint64) {
	paginator.Lock()
	defer paginator.Unlock()

	paginator.Total = total
}

func (paginator *Paginator) GetTotal() uint64 {
	paginator.RLock()
	defer paginator.RUnlock()

	return paginator.Total
}
