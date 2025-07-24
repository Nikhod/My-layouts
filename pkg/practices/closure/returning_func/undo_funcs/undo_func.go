package undo_funcs

import "log"

type Storage struct {
	store []int
}

func UndoPractice() {
	cup := Storage{
		store: []int{1111, 230, 25},
	}
	undo := cup.addData(777)

	log.Println(cup.store)

	undo()
	undo()
	undo()

	log.Println(cup.store)

}

// for undo_practice
func (s *Storage) addData(num int) (undo func()) {
	s.store = append(s.store, num)

	return func() {
		indexOfLastElement := len(s.store) - 1
		s.store = s.store[:indexOfLastElement]
	}
}
