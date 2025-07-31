package func_is_argument

import (
	"fmt"
	"log"
)

func WrapperPractice() {

	notif := []string{"first", "23", "second", "le", "third", "four"}

	for i := 0; i < len(notif); i++ {
		wrapper(notif[i], sendMessage)
	}

}

// обертка для главной функции SendMessage. Также можно вернуть
// какую-то функцию, для иной (продвинутой) логики
func wrapper(msg string, f func(string)) {
	log.Println("start")
	f(msg)
	log.Println("finish")

	fmt.Println()
}

func sendMessage(msg string) {
	log.Println("message:", msg)
}
