package goroutine_practice

import (
	"fmt"
	"runtime"
)

func PracticeWithScheduleOfGoroutines() {
	runtime.GOMAXPROCS(1) // max cores that can be used during the execution of app

	fmt.Printf("FROM MAIN:::		max CPU cores: %d\n", runtime.NumCPU())
	go CountUntiln(50)

	runtime.Gosched()

	fmt.Println("EXIT")
}

func CountUntiln(n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("FROM GOROUTINE:::	number: %d\n", i)
	}
	fmt.Println("THE COUNT IS OVER")
}
