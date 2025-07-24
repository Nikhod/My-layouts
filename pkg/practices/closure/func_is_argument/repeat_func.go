package func_is_argument

import "log"

func Repeat(kolvo int, afresh func(n int)) {

	for i := 1; i <= kolvo; i++ {
		afresh(i)
	}

}

func RepeatPractice() {
	amountOfTries := 5

	Repeat(amountOfTries, func(n int) {
		log.Println(">>>attempt #", n)
	})

}
