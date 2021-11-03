package paginator

import (
	"sync"
)

type Paginator struct {
	Page  int `gorm:"-" json:"page"`
	Limit int `gorm:"-" json:"limit"`
	Total int `gorm:"total" json:"total"`
	*sync.RWMutex
}

func (paginator Paginator) GetLimit() int {
	paginator.RLock()
	defer paginator.RUnlock()

	return paginator.Limit
}

func (paginator *Paginator) Offset() int {
	paginator.RLock()
	defer paginator.RUnlock()

	if paginator.Page == 1 {
		return 0
	}

	return (paginator.Page - 1) * paginator.Limit
}

func (paginator *Paginator) SetTotal(total *int64) {
	paginator.Lock()
	defer paginator.Unlock()

	paginator.Total = int(*total)
}
