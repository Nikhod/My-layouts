package goroutine_practice

import (
	"fmt"
	"sync"
)

func OrdinaryExampleWorkingWithChannel() {
	var wg sync.WaitGroup
	firstChan := make(chan int)

	wg.Add(1)
	go func() {
		defer wg.Done()
		firstChan <- 777
	}()

	valueWasGivenWithDataFromChan := <-firstChan
	fmt.Println(valueWasGivenWithDataFromChan)

	wg.Wait()
}
