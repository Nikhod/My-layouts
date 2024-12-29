package practices

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
)

/*
this function has two issues:
- 1) the queue of file parts is unconstantable
- 2) last string of each file is lost during the transfer
*/
func UniteToOneFile() {
	var wg sync.WaitGroup
	resultChanWithData := make(chan []byte, 4)

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
