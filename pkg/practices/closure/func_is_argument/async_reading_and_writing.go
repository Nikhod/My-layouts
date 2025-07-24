package func_is_argument

import (
	"fmt"
	"log"
	"sync"
)

// Storage - a storage of slice of int
type Storage struct {
	store []int
}

//-=--=-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=-=-=-=-=-=-BUFFERED CHANNEL-=-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=-=-=-=-=
// 🧠 Вывод
//❌ Твой код не работает асинхронно, потому что всё выполняется в одном потоке.
//✅ Закрытие канала до чтения — нормальная практика, если ты гарантируешь, что все записи завершились до чтения
//(как в твоём случае).
//✅ Чтение из закрытого канала — безопасно, и range завершит итерацию корректно.

func (c *Storage) eachElementOfBufferedCup(pocket chan int, read func(pocketInRead chan int)) {
	for i := 0; i < len(c.store); i++ {
		pocket <- c.store[i]
	}

	close(pocket)
	read(pocket)

}

func BufferedCupPractice() {
	from := Storage{store: []int{12, 20, 13, 5, 100}}
	var pocketMain = make(chan int, len(from.store))

	from.eachElementOfBufferedCup(pocketMain, func(pocketInClosure chan int) {
		for value := range pocketInClosure {
			fmt.Println(value)
		}
	})

	log.Println("finish")
}

//-=--=-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=-=-=-=-=-=-UNBUFFERED CHANNEL-=-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=-=-=-=-=

//// 🧠 Вывод
//Чтение и запись через небуферизованный канал должны происходить одновременно в разных горутинах.
//Используй defer close(chan) после всех записей, и range chan для чтения.
//Для синхронизации — sync.WaitGroup.

func (c *Storage) eachElementOfUnbufferedCup(wg *sync.WaitGroup, pocket chan int, read func(pocketInRead chan int)) {
	defer wg.Done()
	defer close(pocket)

	go read(pocket)
	for i := 0; i < len(c.store); i++ {
		pocket <- c.store[i]
	}

}

func UnbufferedCupPractice() {
	from := Storage{store: []int{12, 20, 13, 5, 100}}
	var wg sync.WaitGroup
	var pocketMain = make(chan int)

	wg.Add(1)
	from.eachElementOfUnbufferedCup(&wg, pocketMain, func(pocketInClosure chan int) {
		for value := range pocketInClosure {
			fmt.Println(value)
		}
	})

	wg.Wait()
	log.Println("finish")
}
