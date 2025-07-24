package undo_funcs

//🧩 Задача: временная замена значения в карте с Undo
//Реализуй функцию SetTempValue, которая:
//
//Принимает карту map[string]int, ключ string, и новое значение int.
//
//Устанавливает новое значение по ключу.
//
//Возвращает функцию undo(), которая:
//
//Восстанавливает старое значение, если оно было.
//
//Удаляет ключ, если его не было изначально.

func SetTempValue(store map[string]int, key string, value int) (undo func()) {
	oldValue, ok := store[key]

	store[key] = value

	return func() {
		if ok {
			store[key] = oldValue
		} else {
			delete(store, key)
		}
	}

}

func SetTmpValueInMapPractice() {
	store := map[string]int{
		"first":  1,
		"second": 2,
		"third":  3,
	}

	undo := SetTempValue(store, "third", 999)
	undo()

}
