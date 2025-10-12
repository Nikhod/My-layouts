package my_context

import (
	"context"
	"fmt"
	"time"
)

var pigsFly = false

// процесс с неопределенным временем работы
func longProcess(ctx context.Context, params ...interface{}) {

	// канал по которому придет отмена контекста
	killed := ctx.Done()

Loop1:
	for {
		// проверка отмены контекста
		select {
		case <-killed:
			break Loop1 // завершение работы, если контекст отменен
		default:
		}

		// процесс с определенным временем работы
		time.Sleep(time.Second)
		fmt.Printf("Im alive %v\n", time.Now())

		// проверка выполнения задачи
		if pigsFly {
			fmt.Printf("Pigs fly! %v\n", time.Now())
			break Loop1
		}
	}
}
