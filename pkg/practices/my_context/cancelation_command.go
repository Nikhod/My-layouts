package my_context

import (
	"context"
	"log"
	"math/rand"
	"sync"
	"time"
)

// у нас есть сервисы такси, мы делаем поиск самого быстрого (это условие будет симитировано ниже заглушкой)
// как только мы нашли самый быстрый, мы печатаем название его в консоль

func ContextPractice() {
	var (
		wg       sync.WaitGroup
		winner   string
		services = []string{"olucha", "raksh", "eco", "asian express", "maxim", "yandex"}
		//oneCall  sync.Once
		pocket = make(chan string)
	)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for _, serviceName := range services {
		wg.Add(1)
		go func(nameOfService string) {
			findingServiceWithoutEndlessLoop(ctx, nameOfService, pocket)
			//findServiceViaOneCall(ctx, nameOfService, pocket, &oneCall)
			//findingService(ctx, nameOfService, pocket) // - second mean of implementation that same logic as findService
			wg.Done()
		}(serviceName)
	}

	winner = <-pocket
	cancel()

	wg.Wait()
	log.Println("service_winner:", winner)

}

// Функция findingService устроена так, что она многократно проверяет условие isFastest().
// isFastest() возвращает true только с вероятностью ~50%.
// Если бы мы вызвали её один раз как в findingServiceWithoutEndlessLoop и получили false, горутина сразу бы
// завершилась, и этот сервис вообще
// не участвовал бы в гонке.
//
//	Цикл нужен, чтобы сервис мог несколько раз попытаться "победить" до тех пор, пока:
//	_+_+_он не стал победителем, или
//
// _+_+_не пришла отмена (ctx.Done()).
func findingService(ctx context.Context, serviceName string, pocket chan string) {
	for {
		select {
		case <-ctx.Done():
			log.Println(serviceName, ">>> was stopped, ")
			return
		default:
			if isFastest() {
				select {
				case pocket <- serviceName:
					log.Println("winner is known")
				default:
					log.Println("cannot send, the channel is closed")
				}
			}
		}
	}
}

// another way of finding service without race condition via sync.Once.
func findServiceViaOneCall(ctx context.Context, serviceName string, pocket chan string, oneCall *sync.Once) {
	send := func() {
		pocket <- serviceName
	}
	for {
		select {
		case <-ctx.Done():
			log.Println("was stopped ")
			return
		default:
			if isFastest() {
				oneCall.Do(send)
				return
			}
		}
	}
}

// emulation of condition
func isFastest() bool {
	time.Sleep(time.Duration(100+rand.Intn(100)) * time.Millisecond)
	return rand.Float64() > 0.5
}

// смотреть выше, там подробнее
func findingServiceWithoutEndlessLoop(ctx context.Context, serviceName string, pocket chan string) {
	select {
	case <-ctx.Done():
		log.Println("was stopped ")
		return
	default:
		if isFastest() {
			select {
			case pocket <- serviceName:
			default:
				log.Println("channel is closed")
			}
			return
		}
	}
}
