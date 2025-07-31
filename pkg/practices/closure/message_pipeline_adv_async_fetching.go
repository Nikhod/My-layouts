package closure

import (
	"log"
	"sync"
)

func FetchingPractice() {
	messages := getMessages()

	fetch := parallelFetching(messages, 2)
	resultMESSAGES := fetch()
	log.Println(len(resultMESSAGES), resultMESSAGES)

}

func parallelFetching(data []message, worker int) (fetch func() []message) {
	// какая то логика, состояние сохраняется, запоминается и может быть использована
	// возвращаемой функцией
	return func() []message {
		var wg sync.WaitGroup
		var result []message
		var pocket = make(chan []message, worker)

		// размер кусочков, по сколько будут брать каждая горутина
		chunkSize := (len(data) / worker) - 1

		for i := 0; i < len(data); i += chunkSize {
			end := i + chunkSize
			if end > len(data) {
				end = len(data)
			}
			wg.Add(1)
			go func(chunk []message) {
				defer wg.Done()
				pocket <- chunk
				log.Println("chunk in pocket")
			}(data[i:end])

		}

		// не блокирует поток main, при этом поток main читает все данные
		// после чего канал закрывается
		go func() {
			wg.Wait()
			close(pocket)
		}()

		for msg := range pocket {
			result = append(result, msg...)
		}

		return result
	}
}
