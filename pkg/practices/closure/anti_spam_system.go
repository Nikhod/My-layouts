package closure

import (
	"log"
	"sync"
	"time"
)

func AntispamLimiterPractice() {
	sendMessage := newChatLimiter(3, time.Second*5, flush)

	sendMessage("first", "the1 first>1  ")
	sendMessage("first", "the1 second>>1 ")
	sendMessage("second", "the2 first>2 sample")

	time.Sleep(time.Second * 3)
	sendMessage("first", "the1 third>>>1 ")
	sendMessage("second", "the2 second>>2 ")

	time.Sleep(time.Second * 8)
	log.Println("main is finished")
}

func newChatLimiter(limit int, timeout time.Duration, flush func(content []string)) (sendMessage func(userID, message string)) {
	buffer := make(map[string][]string)
	timers := make(map[string]*time.Timer)

	var mtx sync.Mutex

	checkingAndFlush := func(userID string) {
		mtx.Lock()
		defer mtx.Unlock()

		content := buffer[userID]
		if len(content) > 0 {
			flush(content)
			buffer[userID] = nil
		}
	}

	resetTimer := func(userID string) {
		if timers[userID] != nil {
			timers[userID].Stop()
		}

		timers[userID] = time.AfterFunc(timeout, func() {
			checkingAndFlush(userID)
		})
	}

	return func(userID, message string) {
		mtx.Lock()
		defer mtx.Unlock()

		buffer[userID] = append(buffer[userID], message)
		if len(buffer[userID]) == limit {
			flush(buffer[userID])
			buffer[userID] = nil
		}

		resetTimer(userID)
	}

}
