package solutions

import (
	"Nikcase/pkg/models"
	"bufio"
	"encoding/json"
	"fmt"
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

		bytesInFile, err := os.ReadFile("example.json")
		if err != nil {
			log.Println(err)
			close(pocket)
			return
		}

		err = json.Unmarshal(bytesInFile, &user)
		if err != nil {
			log.Println(err)
			close(pocket)
			return
		}

		pocket <- user
		close(pocket)
	}()

	if user, isOpen := <-pocket; isOpen {
		fmt.Println(user)
	} else {
		fmt.Println("no data were received")
	}
	wg.Wait()

	fmt.Println("code after processing data in goroutine and after reading that from MAIN goroutine")
}

func ReadOneFileWithNumerousGoroutines() {
	// one can add the param "n" and "fileName" to signature of function to control
	//amount of goroutines, which can read from specific (which have been pointed in signature) file
	const amountOfGoroutines = 5
	chanWithDataFromFile := make(chan []byte, amountOfGoroutines)
	var wg sync.WaitGroup

	file, err := os.OpenFile("my_file.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		log.Println("error in reading file: ", err)
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		log.Println(err)
		return
	}

	sizeFile := stat.Size()
	amountOfSegments := amountOfGoroutines / sizeFile // how many bytes have to process each goroutine

	for queueOfGo := 1; queueOfGo <= amountOfGoroutines; queueOfGo++ {
		wg.Add(1)

		start := int64(queueOfGo) * amountOfSegments // offset is counted with bytes
		end := start + amountOfSegments

		if queueOfGo == (amountOfGoroutines) {
			end = sizeFile
		}

		go readSegment(&wg, file, start, end, chanWithDataFromFile)
	}

	go func() {
		wg.Wait()
		close(chanWithDataFromFile)
	}()

	// reading from channels
	// code
	for data := range chanWithDataFromFile {
		fmt.Println(string(data))
	}
}

func readSegment(wg *sync.WaitGroup, file *os.File, start, end int64, chanWithDataFromFile chan []byte) {
	var mtx sync.Mutex
	defer wg.Done()
	segment := make([]byte, end-start) // segment = buffer
	// "end" and "start" - valuable with amount of bytes [offset=смещение]
	mtx.Lock()
	_, err := file.ReadAt(segment, start)
	mtx.Unlock()
	if err != nil {
		log.Println(err)
		return
	}

	chanWithDataFromFile <- segment
}

func UniteToOneFile() {
	var wg sync.WaitGroup
	resultChanWithData := make(chan []byte)

	var fileNames = []string{
		"first_file.txt",
		"second_file.txt",
		"third_file.txt",
		"fourth_file.txt"}

	for _, fileName := range fileNames {
		wg.Add(1)
		go readFile(&wg, fileName, resultChanWithData)
	}

	go func() {
		wg.Wait() // in order to stop main goroutine
		close(resultChanWithData)
	}()

	var combinedData []byte
	for dataFromChan := range resultChanWithData {
		combinedData = append(combinedData, dataFromChan...)
	}

	fmt.Println(string(combinedData))
}

func readFile(wg *sync.WaitGroup, fileName string, resultChan chan []byte) {
	defer wg.Done()
	file, err := os.Open(fileName)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()

	var resultOfBytes []byte
	reader := bufio.NewReader(file)

	for {
		bytesFromFile, err := reader.ReadBytes('\n')
		if err != nil {
			break
		}
		resultOfBytes = append(resultOfBytes, bytesFromFile...)
	}

	// there is two line break in the very ending of each file
	resultOfBytes = append(resultOfBytes, []byte("\n")...)

	resultChan <- resultOfBytes
}
