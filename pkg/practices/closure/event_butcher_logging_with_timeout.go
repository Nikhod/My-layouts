package closure

import (
	"log"
	"sync"
	"time"
)

//⏱️ ЗАДАЧА: Batcher с таймаутом
//📋 Описание:
//Ты пишешь систему, которая собирает события и отправляет их пачкой:
//если собралось N штук — отправляем (flush);
//если за T секунд не собралось N — тоже отправляем (flush "недобитой" партии по таймауту).

func BatcherTimeoutPractice() {
	timeout := time.Second * 5
	limit := 3
	batcher := newButcherWithTimeout(limit, timeout, flush)

	batcher("event_1")
	batcher("event_2")
	batcher("event_3")

	batcher("event_4")
	//time.Sleep(time.Second * 10)
	batcher("event_5")
	//batcher("event_6")

	mainStopTimer := time.NewTimer(time.Second * 6)
	<-mainStopTimer.C
	log.Println("main is finished")

}
func newButcherWithTimeout(limit int, timeout time.Duration, flush func([]string)) (batchWithTimeout func(event string)) {
	var eventStore []string
	var mtx sync.Mutex
	var timer *time.Timer

	afterTimeoutFunc := func() {
		mtx.Lock()
		defer mtx.Unlock()

		if len(eventStore) > 0 {
			flush(eventStore)
			eventStore = eventStore[:0]
		}
	}

	resetTimer := func() {
		// если таймера нет, запускаем новый
		// если таймер уже запущен, значит структурка не пустая, нужно остановить таймер, перезапустить новый

		if timer == nil {
			timer = time.AfterFunc(timeout, afterTimeoutFunc)
		} else {
			timer.Stop()
			timer = time.AfterFunc(timeout, afterTimeoutFunc)
		}

	}

	return func(event string) {
		mtx.Lock()
		defer mtx.Unlock()

		eventStore = append(eventStore, event)
		if len(eventStore) >= limit {
			flush(eventStore)
			eventStore = eventStore[:0]
			if timer != nil { // проверяй, иначе будет паника
				timer.Stop()
			}
			return
		}

		resetTimer()

	}
}

func ChatnewButcherWithTimeout(limit int, timeout time.Duration, flush func([]string)) func(string) {
	var mu sync.Mutex
	var eventStore []string
	var timer *time.Timer

	// Это будет вызываться, когда таймер сработает
	onTimeout := func() {
		mu.Lock()
		defer mu.Unlock()

		if len(eventStore) > 0 {
			flush(eventStore)
			eventStore = nil
		}
	}

	// А это отдельная функция для запуска / перезапуска таймера
	resetTimer := func() {
		if timer != nil {
			timer.Stop()
		}
		timer = time.AfterFunc(timeout, onTimeout)
	}

	return func(event string) {
		mu.Lock()
		defer mu.Unlock()

		eventStore = append(eventStore, event)

		if len(eventStore) >= limit {
			flush(eventStore)
			eventStore = nil
			if timer != nil {
				timer.Stop()
			}
			return
		}

		resetTimer()
	}
}
