package goroutine_practice

import (
	"Nikcase/pkg/solutions"
	"fmt"
	"log"
	"sync"
	"time"
)

func CorrectUsageOfMutex() {
	var wg sync.WaitGroup

	var mutex sync.Mutex
	rawOfNumbers := []int{3, 4, 5, 6}
	results := make([]int, len(rawOfNumbers))

	for index, aNumber := range rawOfNumbers {
		go func(index, number int) {
			mutex.Lock()
			results[index] = solutions.FindSquare(number)
			time.Sleep(time.Second) // imitation of work
			fmt.Println(results)
			mutex.Unlock()
		}(index, aNumber)
	}

	wg.Wait()
	log.Println(results)
}
