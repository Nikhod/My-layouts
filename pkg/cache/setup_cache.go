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

func SetupCache() *CachedLimits {

	cached := CachedLimits{
		Limits: make(map[int]models.AmountLimits)}

	return &cached
}

func (cl *CachedLimits) SetLimit(limit *models.AmountLimits) error {
	cl.Mtx.Lock()
	defer cl.Mtx.Unlock()
	id := cl.SortedLimitsById[len(cl.SortedLimitsById)-1].Id + 1
	limit.Id = id

	// todo попробуй использовать метод create gorm, возможно он сразу создаст запись и вернет id записи в саму структуру.
	//  todo а потом уже сохраняй в кеш. Пусть методы ДБ будут в связке с методами кеша!

}
