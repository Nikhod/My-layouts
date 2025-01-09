package goroutine_practice

import (
	"fmt"
	"sync"
)

func MainWorkWithSelect() {
	var wg sync.WaitGroup

	dataChan := make(chan int)
	exitChan := make(chan string)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			if i == 9 {
				exitChan <- "stop"
			}
			dataChan <- i
		}
		close(dataChan)
		close(exitChan)

	}()

	ChooseOneOption(dataChan, exitChan)

	wg.Wait()
	fmt.Printf("work is completed")
}

func ChooseOneOption(dataChan chan int, exit chan string) {
	for {
		select {
		case <-dataChan:
			fmt.Printf("number from data channel: %d \n", <-dataChan)
		case <-exit:
			fmt.Printf("execution was interrupted")
			return
		}
	}
}
