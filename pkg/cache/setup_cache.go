package cache

import (
	"Nikcase/pkg/models"
	"sync"
)

type CachedLimits struct {
	Mtx              sync.RWMutex
	Limits           map[int]models.AmountLimits
	SortedLimitsById []models.AmountLimits
}

func (cl *CachedLimits) SetLimit(limit *models.AmountLimits) error {
	cl.Mtx.Lock()
	cl.Mtx.Unlock()
}
