package practices

import (
	"log"
	"math/rand"
	"sync"
)

func TotalSumOfTwoChan(firstChan chan int, secondChan chan int) (int, int) {
	var firstTotal int
	var secondTotal int

	for {
		select {
		case firstMsg, isOpen := <-firstChan:
			if isOpen {
				firstTotal += firstMsg // firstMsg = number int
			} else {
				firstChan = nil // чтобы в дальнейшем был проигнорирован, поскольку, если мы используем return, то мы
				// выйдем из функции не обработав secondChan
			}

		case secondMsg, isOpen := <-secondChan:
			if isOpen {
				secondTotal += secondMsg // secondMsg = number int
			} else {
				secondChan = nil // закрыт ли этот канал или нет
			}
		}

		// в цикле таким образом мы проверим закрыты ли оба канала или  нет
		if firstChan == nil && secondChan == nil {
			return firstTotal, secondTotal
		}
	}
}

func TotalSumOfNumberInTwoChannels() {
	var wg sync.WaitGroup
	firstChan, secondChan := make(chan int), make(chan int)

	wg.Add(1)
	go func() {
		var randomNum int
		defer wg.Done()
		for i := 0; i < 10; i++ {
			randomNum = rand.Intn(50)
			firstChan <- randomNum

			randomNum = rand.Intn(50)
			secondChan <- randomNum
		}
		close(firstChan)
		close(secondChan)
	}()

	firstTotal, secondTotal := TotalSumOfTwoChan(firstChan, secondChan)

	wg.Wait()
	log.Printf("total was decided,\nfrom first chan: %d\nfrom second chan: %d\n", firstTotal, secondTotal)
}
