package closure

import (
	"encoding/json"
	"errors"
	"log"
	"os"
)

func EtlPractice() {
	fileName := "./sample_file.txt"

	tidyData := extract(fileName) // коннект с источником

	transactions := tidyData() // выгрузка данных с источника

	transactions = transform(transactions, filter, modifier) // трансформация данных в нужный нам формат и фильтер
	upload := load(transactions)                             // может быть какая то логика
	err := upload()                                          // выгрузка данных
	if err != nil {
		log.Println(err)
		return
	}

}

// ✅ Сценарий (ETL pipeline)
//Мы имитируем поток финансовых транзакций:
//
//Extract (забрать сырые данные) →
//Transform (очистить и пересчитать суммы) →
//Load (записать в хранилище — имитируем БД)

type transaction struct {
	ID     int
	Amount float64
	Valid  bool
}

//💡 Где использовать замыкания
//Этап	 	|||Что делаем							|||Closure тип	|||Почему?
//Extract	|||Получаем данные из источника			|||Возвращаем функцию (closure)		|||Можно замкнуть источник данных
//Transform	|||Передаём очистители и трансформеры	|||Принимаем функцию как аргумент	|||Гибкость обработки
//Load		|||Записываем куда-то					|||Возвращаем функцию (closure)		|||Можно замкнуть "хранилище"

// 🔹 Что значит «замкнуть источник данных» (в Extract)?
//	💡 Идея:
//	Ты создаёшь функцию, которая «запоминает» источник (данные, параметры подключения, файл, указатель, счётчик),
//	и при каждом вызове работает с этим источником.

func extract(filepath string) (tidyData func() []transaction) {
	file, _ := os.Open(filepath)
	defer func() {
		// todo file.close()
	}()

	return func() []transaction {
		var source []transaction
		_ = json.NewDecoder(file).Decode(&source)

		return nil
	}
}

func transform(transactions []transaction, filter func(transaction) bool, modifier func(*transaction) *transaction) []transaction {
	var result []transaction
	log.Println("transforming...")

	for _, trx := range transactions {
		if filter(trx) {
			result = append(result, *modifier(&trx))
		}
	}
	return result
}

func load(data []transaction) (upload func() error) {
	// можно добавить какую-то логику
	return func() error {
		for _, value := range data {
			ok := loadToDb(value)
			if !ok { // эмуляция ошибки
				return errors.New("couldn't load to db")
			}
		}
		return nil
	}
}

// helper functions:
func modifier(tx *transaction) *transaction {
	tx.Valid = true
	tx.ID += 10
	tx.Amount += 500
	log.Println("was modified")
	return tx
}

func filter(t transaction) bool {
	if t.ID < 0 || t.Amount < 10 || t.Valid == false {
		return false
	}
	log.Println("filter_Ok")
	return true
}

func removeElement(tx []transaction, index int) []transaction {
	if index+1 < len(tx) {
		right := tx[index+1:]
		tx = tx[:index]
		tx = append(tx, right...)
		log.Println("was removed")
		return tx
	} else {
		log.Println("was removed")
		return tx[:index]
	}
}

// пустышка
func loadToDb(tx transaction) bool { return true }
