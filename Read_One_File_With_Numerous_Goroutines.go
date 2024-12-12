package test

import (
	"fmt"
	"log"
	"os"
	"sync"
)

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
