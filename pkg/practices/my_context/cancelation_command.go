package my_context

import (
	"context"
	"log"
	"math/rand"
	"sync"
)

// у нас есть сервисы такси, мы делаем поиск самого быстрого (это условие будет симитировано ниже заглушкой)
// как только мы нашли самый быстрый, мы печатаем название его в консоль

//my_context.WithDeadline() = my_context.WithTimeout()

func ContextPractice() {
	var (
		mtx      sync.Mutex
		wg       sync.WaitGroup
		winner   string
		services = []string{"olucha", "raksh", "eco", "asian express", "maxim", "yandex"}
		pocket   = make(chan string)
	)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for _, serviceName := range services {
		wg.Add(1)
		go func(serviceName string) {
			findingService(ctx, serviceName, pocket)
			wg.Done()
		}(serviceName)
	}

	go func() {
		mtx.Lock()
		winner = <-pocket
		mtx.Unlock()
		cancel()
	}()
	wg.Wait()
	log.Println("service_winner:", winner)
}

func findingService(ctx context.Context, serviceName string, pocket chan string) {
	for {
		select {
		case <-ctx.Done():
			log.Println("was stopped ")
			return
		default:
			if isFastest() {
				pocket <- serviceName
				return
			}
		}

	}
}

// emulation of condition
func isFastest() bool {
	return rand.Float64() > 0.5
}

func findingServiceWithoutEndlessLoop(ctx context.Context, serviceName string, pocket chan string) {
	select {
	case <-ctx.Done():
		log.Println("was stopped ")
		return
	default:
		if isFastest() {
			pocket <- serviceName
			return
		}
	}
}
