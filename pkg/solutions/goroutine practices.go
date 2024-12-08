package solutions

import (
	"Nikcase/pkg/models"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"runtime"
	"sync"
	"time"
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
			//totalSumOfFirstChan <- firstTotal
			//totalSumOFSecondChan <- secondTotal

			fmt.Printf("first Total: %d\nsecond total: %d\n", firstTotal, secondTotal)
			return firstTotal, secondTotal
			//return totalSumOfFirstChan, totalSumOFSecondChan
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

	TotalSumOfTwoChan(firstChan, secondChan)

	wg.Wait()
	log.Printf("total was decided,\nfrom first chan: %d\nfrom second chan: %d\n")
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

func PracticeWithScheduleOfGoroutines() {
	runtime.GOMAXPROCS(1) // max cores that can be used during the execution of app

	fmt.Printf("FROM MAIN:::		max CPU cores: %d\n", runtime.NumCPU())
	go CountUntiln(50)

	runtime.Gosched()

	fmt.Println("EXIT")
}

func CountUntiln(n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("FROM GOROUTINE:::	number: %d\n", i)
	}
	fmt.Println("THE COUNT IS OVER")
}

func CorrectUsageOfMutex() {
	var wg sync.WaitGroup

	var mutex sync.Mutex
	rawOfNumbers := []int{3, 4, 5, 6}
	results := make([]int, len(rawOfNumbers))

	for index, aNumber := range rawOfNumbers {
		go func(index, number int) {
			mutex.Lock()
			results[index] = FindSquare(number)
			time.Sleep(time.Second) // imitation of work
			fmt.Println(results)
			mutex.Unlock()
		}(index, aNumber)
	}

	wg.Wait()
	log.Println(results)
}

func ReadFileWithGoroutine() {
	var wg sync.WaitGroup
	pocket := make(chan models.Users)
	wg.Add(1)

	go func() {
		defer wg.Done()
		var user models.Users
		var mtx sync.Mutex

		mtx.Lock()
		file, err := os.OpenFile("./example.json", os.O_WRONLY|os.O_RDWR, 0666)
		if err != nil {
			log.Println(err)
			return
		}

		bytes, err := io.ReadAll(file)
		if err != nil {
			log.Println(err)
			return
		}

		err = json.Unmarshal(bytes, &user)
		if err != nil {
			log.Println(err)
			return
		}
		mtx.Unlock()

		pocket <- user
	}()
	wg.Wait()

	fmt.Println(<-pocket)
}
