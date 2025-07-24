package returning_func

import (
	"log"
)

func CreateDivider(divider int) func(num int) int {
	dividerFunc := func(num int) int {

		return num / divider
	}
	return dividerFunc
}

type HubOfDivisor struct {
	// key - divisor
	// value - func, where num (dividend) can be divided by divisor
	// "func(num)" - "num" is the dividend
	Divider map[int]func(num int) int
}

func PracticeInputCustomFuncHub() {
	hub := HubOfDivisor{Divider: make(map[int]func(dividend int) int)}

	divider := []int{2, 3, 5, 10}
	for i := 0; i < len(divider); i++ {
		key := divider[i]

		divfunc := CreateDivider(key)
		hub.Divider[key] = divfunc
	}

	log.Println(hub.Divider[5](250))

	return

	resultOfDivision := hub.Divider[2](20)
	log.Println("20 / 2 =", resultOfDivision)

	//fmt.Printf("%+v\n", hub.Divider[2])

}
