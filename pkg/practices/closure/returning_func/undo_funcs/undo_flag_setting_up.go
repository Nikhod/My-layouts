package undo_funcs

import (
	"errors"
	"os"
)

var flag any = "NOT_EMPTY"

func changeFlag(newValue interface{}) (undo func(), err error) {
	oldValue := flag

	flag = newValue

	undo = func() {
		flag = oldValue
	}

	//предположим что тут будет ошибка, если 'flag' будет пустым
	if flag == nil {
		err = errors.New("there is no element in flag")
		return undo, err
	}

	return undo, nil
}

func ChangingFlagPractice() {
	newValue := os.Getenv("sample_key")

	undo, err := changeFlag(newValue)
	if err != nil {
		undo()
		return
	}

}
