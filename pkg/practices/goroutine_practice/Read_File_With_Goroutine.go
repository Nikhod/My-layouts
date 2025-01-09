package goroutine_practice

import (
	"Nikcase/pkg/models"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
)

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
