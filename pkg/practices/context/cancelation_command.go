package context

import (
	"context"
	"log"
	"math/rand"
	"sync"
)

//todo description

//context.WithDeadline() = context.WithTimeout()

func ContextPractice() {
	var (
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
		winner = <-pocket
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
			if isValidToCondition() {
				pocket <- serviceName
				return
			}
		}

		continue
	}
}

func findingServiceWithoutEndlessLoop(ctx context.Context, serviceName string, pocket chan string) {
	select {
	case <-ctx.Done():
		log.Println("was stopped ")
		return
	default:
		if isValidToCondition() {
			pocket <- serviceName
			return
		}
	}
}

// emulation of condition
func isValidToCondition() bool {
	if rand.Float64() > 0.75 {
		return true
	} else {
		return false
	}
}
