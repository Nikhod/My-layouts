package closure

import "fmt"

// 📥 ЗАДАЧА: Batcher — накопление событий с последующей обработкой
//🧾 Описание:
//Ты пишешь компонент для логгера, который собирает события (например, ошибки или действия пользователей), но отправляет
//их пачкой — не по одному.
//Задача: написать функцию NewBatcher(limit int, flush func([]string)), которая вернёт замыкание func(event string).
//Это замыкание накапливает события, и как только их становится limit — вызывает flush с этой пачкой.

func BatcherPractice() {
	batch := newButcher(3, flush)
	for i := 0; i < 6; i++ {
		batch(fmt.Sprint("event_", i+1))
	}

}

func newButcher(limit int, flush func([]string)) func(string) {
	var eventStore []string // никогда не очищается, а передается по кусочкам в flush(). Моожно будет очистить полностью
	var counter int
	var startIndex int

	return func(event string) {
		eventStore = append(eventStore, event)
		counter++
		if counter%limit == 0 { // нужно чтобы каждый limit-раз
			flush(eventStore[startIndex:counter])
			startIndex = counter
			// после того как мы выдадим пачку,
			// стартовый индекс будет перемещен до следующего актуального значения
			//eventStore = eventStore[:0] <---- reset store для будущей возможной реализации
		}
	}
}

//	func flush(events []string) {
//		fmt.Println(events)
//	}
func flush(events []string) {
	fmt.Println(events)
}
