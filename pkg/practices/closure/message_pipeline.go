package closure

import (
	"gorm.io/gorm"
	"log"
)

func MessagePipeline() {
	messages := getMessages()   // getting messages from source
	fetching := fetch(messages) // one can add some logic before the fetching
	messages = fetching()

	messages = validate(messages, isValid) // validate and adding only valid data to []message

	save := connectAndSaveToDB(messages) // connecting to db

	logPrintToConsole := wrapperSave(save) // one can add logic before wrapperSave
	logPrintToConsole()                    // log to console

}

type message struct {
	msgType string
	body    string
	id      int
}

func fetch(msg []message) func() []message {
	// some logic
	return func() []message {
		log.Println("fetching...")
		return msg
	}

}

func validate(msgs []message, isValid func(msg message) bool) []message {
	var validatedMess []message
	for i := 0; i < len(msgs); i++ {
		iterable := msgs[i]
		if isValid(iterable) {
			validatedMess = append(validatedMess, iterable)
		}
	}

	return validatedMess
}

func connectAndSaveToDB(data []message) (save func() error) {
	//показано для примера, чтобы понять гибкость работы с замыканиями
	db := getDBConnection() // функция которую мы вернем будет работать с этими данными,
	// с заполненными полями этого GORM экземляра

	return func() error {
		//замыкание для транзакции, если будет ошибка,
		// то сделает откат на уровне БД
		err := db.Transaction(func(tx *gorm.DB) error {
			err := tx.Table("messages").Save(data).Error
			if err != nil {
				return err
			}

			return nil
		})

		if err != nil {
			return err
		}

		return nil
	}
}

// обертка для функции которя сохраняет данные в БД. Она возвращается от connectAndSaveToDB
func wrapperSave(save func() error) (logging func()) {
	return func() {
		err := save()
		if err != nil {
			log.Println("my error:", err)
			return
		} else {
			log.Println("save successfully")
			return
		}
	}
}

// helper functions

func isValid(msg message) bool {
	if msg.id < 1 || len(msg.body) < 5 || msg.msgType != "sms" {
		return false
	}

	return true
}

// empty func - emulation
func getDBConnection() *gorm.DB { return &gorm.DB{} }

func getMessages() []message {
	return []message{
		{
			msgType: "sms",
			body:    "first",
			id:      1,
		},
		{
			msgType: "sms",
			body:    "second",
			id:      2,
		},
		{
			msgType: "sms",
			body:    "third",
			id:      3,
		},
		{
			msgType: "sms",
			body:    "fourth",
			id:      4,
		},
		{
			msgType: "sms",
			body:    "fifth",
			id:      5,
		},
		{
			msgType: "sms",
			body:    "sixth",
			id:      6,
		},
		{
			msgType: "sms",
			body:    "seventh",
			id:      7,
		},
		{
			msgType: "sms",
			body:    "eighth",
			id:      8,
		},
		{
			msgType: "sms",
			body:    "ninth",
			id:      9,
		},
		{
			msgType: "sms",
			body:    "tenth",
			id:      10,
		},
		{
			msgType: "sms",
			body:    "eleventh",
			id:      11,
		},
		{
			msgType: "sms",
			body:    "twelve",
			id:      12,
		},
		{
			msgType: "sms",
			body:    "thirteenth",
			id:      13,
		},
	}
}
