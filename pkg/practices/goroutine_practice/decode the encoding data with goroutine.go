package goroutine_practice

import (
	"Nikcase/pkg/models"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"sync"
)

func DecodeTheEncodingDataWithGoroutine() {
	var user models.Users
	var wg sync.WaitGroup
	pocket := make(chan []byte)

	wg.Add(2)
	go EncodingPractice(pocket, &wg)
	go DecodingPractice(pocket, &wg, &user)
	wg.Wait()

	fmt.Println(user)
}

func EncodingPractice(pocket chan []byte, wg *sync.WaitGroup) {
	defer wg.Done()
	var user = models.Users{
		Name:     "Maduro",
		Login:    "Nikolas",
		Password: "Johnson",
		Active:   true,
	}
	firstBuffer := bytes.Buffer{}

	// one of the way of buffer initialization
	//var buf []byte
	//secondBuffer := bytes.NewBuffer(buf)

	err := json.NewEncoder(&firstBuffer).Encode(user)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Printf("len of bytes = %d\ncapacity of buffer = %d\n",
		firstBuffer.Len(), firstBuffer.Cap())

	lineInBytes, err := firstBuffer.ReadBytes('\n')
	if err != nil {
		log.Println(err)
		return
	}

	pocket <- lineInBytes
}

func DecodingPractice(pocket chan []byte, wg *sync.WaitGroup, user *models.Users) {
	//var user models.Users
	defer wg.Done()
	bytesFromPocket, isOpen := <-pocket
	if !isOpen {
		log.Println("the channel is closed")
		return
	}

	reader := bytes.NewReader(bytesFromPocket)
	err := json.NewDecoder(reader).Decode(&user)
	if err != nil {
		log.Println(err)
		return
	}
}
