package debounce

import (
	"context"
	"sync"
	"time"
)

func DebounceLast(circuit Circuit, d time.Duration) Circuit {
	var (
		ticker    *time.Ticker
		threshold time.Time
		mtx       sync.Mutex
		once      sync.Once
	)
	return func(ctx context.Context) (resp any, err error) {
		mtx.Lock()
		defer mtx.Unlock()

		threshold = time.Now().Add(d)
		once.Do(func() {
			ticker = time.NewTicker(time.Millisecond * 100)

			go func() {
				defer func() {
					once = sync.Once{}
					ticker.Stop()
					mtx.Unlock()
				}() // до завершения горутины должна быть гигиена

				for {
					select {
					case <-ticker.C:
						mtx.Lock()
						if time.Now().After(threshold) {
							resp, err = circuit(ctx)
						}
						return
					case <-ctx.Done():
						mtx.Lock()
						resp, err = nil, ctx.Err()
						return
					}
				}
			}()
		})
		return
	}
}
