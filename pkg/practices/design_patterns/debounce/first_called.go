package debounce

import (
	"context"
	"errors"
	"log"
	"sync"
	"time"
)

type Circuit func(ctx context.Context) (resp any, err error)

// d - это интервал времени до истечения которого повторного вызова функции не будет
func DebounceFirst(circuit Circuit, d time.Duration) Circuit {
	var (
		mtx       sync.Mutex
		threshold time.Time
	)

	return func(ctx context.Context) (resp any, err error) {
		mtx.Lock()
		defer func() {
			threshold = time.Now().Add(d)
			mtx.Unlock()
		}()

		if time.Now().Before(threshold) {
			return nil, errors.New("function is unreachable")
		}
		resp, err = circuit(ctx)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		return resp, nil
	}
}
