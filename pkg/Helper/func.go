package Helper

import (
	"Nikcase/pkg/models"
	"encoding/json"
	"golang.org/x/net/context"
	"log"
	"time"
)

/*
	Принимает на вход Текст ответа, добавляет в поле структуры

и сериализует его.
*/
func ResponseAnswer(report string) (myAnswer []byte, err error) {
	answer := models.Answer{
		Date:           time.Now(),
		ResponseAnswer: report,
	}

	myAnswer, err = json.Marshal(answer)
	if err != nil {
		log.Println(err)
		return
	}
	return
}

func GetIdFromContext(ctx context.Context) (id int, err error) {
	return
}
