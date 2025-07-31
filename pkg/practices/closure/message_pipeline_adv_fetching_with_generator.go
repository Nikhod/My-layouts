package closure

import (
	"fmt"
	"log"
)

func GeneratorFetchPractice() {
	data := getMessages()
	getNextChunk := generateMessages(data, 4)

	chunk := getNextChunk()
	log.Println(chunk)

	chunk = getNextChunk()
	log.Println(chunk)

	chunk = getNextChunk()
	log.Println(chunk)

	chunk = getNextChunk()
	log.Println(chunk)

	next := generateNextMessage(data)

	fmt.Println("\n\n<<<<<<<<<<<ITERATOR>>>>>>>>>")
	for i := 0; i < len(data); i++ {
		log.Println(next()) // <---*
	}

}

// reads and returns data by chunk. when we reach to the last chunk of elements, then after calling 'generate'
// we always get the last chunk even if we call the 'generate' endlessly
func generateMessages(data []message, chunkSize int) (generate func() []message) {
	var start int
	return func() []message {
		endIndex := start + chunkSize // остаточный кусок
		if len(data) < endIndex {
			endIndex = len(data)
			return data[start:endIndex]
		}

		currentChunk := data[start:endIndex]

		start += chunkSize
		return currentChunk
	}
}

func generateNextMessage(data []message) (next func() (message, bool)) {
	var index int // состояние сохраняется
	return func() (message, bool) {
		if len(data) < index {
			return message{}, false
		}

		current := data[index]

		index++
		return current, true
	}
}
