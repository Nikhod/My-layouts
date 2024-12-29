package practices

import (
	"Nikcase/pkg/models"
	"fmt"
	"log"
	"sync"
)

/*
это является копией функции DecodeTheEncodingDataWithGoroutine with some changes:
В вашем примере код работает, потому что основная горутина (main()) читает из канала и выполняет вывод данных, что
обеспечивает правильное выполнение всей конструкции.Конструкция go func() { wg.Wait(); close(pocket);
close(secondChan) }()
работает, потому что:

1) Основная горутина читает из канала: Основная горутина (main()) ждёт, пока из secondChan придёт сообщение, и это
предотвращает преждевременное завершение программы. Так как основная горутина активно ждёт данных из канала, она не
завершает выполнение раньше времени.

2) Корректное использование WaitGroup: Горутины, выполняющие EncodingPractice и DecodingPractice, корректно уменьшают
счётчик WaitGroup. Ваша третья горутина, использующая wg.Wait(), завершает своё выполнение, когда все остальные горутины
завершают свою работу, что предотвращает взаимоблокировку.

ВЫВОД:
Таким образом, если основная горутина читает из канала и ожидает завершения всех горутин, использование конструкции
go func() { wg.Wait(); close(pocket); close(secondChan) }() является корректным и предотвращает преждевременное
завершение программы.

Если бы основной поток бы не читал бы данные с канала (то есть не ждал бы пока в этот канал поступят данные), или не было
бы другой логики, которая бы остановила бы основной поток на какоето время, то конструкция go func() { wg.Wait();
close(pocket); close(secondChan) }() - была бы неверной, и основной поток завершался бы раньше чем горутины
*/
func CorrectWayOfWgUsage() {
	var user models.Users
	var wg sync.WaitGroup
	var secondChan = make(chan string)

	pocket := make(chan []byte)

	wg.Add(3)
	go EncodingPractice(pocket, &wg)
	go DecodingPractice(pocket, &wg, &user)
	go DoSomething(&wg, secondChan)

	go func() {
		wg.Wait()
		close(pocket)
		close(secondChan)
	}()

	log.Println("message from something goroutine func:", <-secondChan)
	fmt.Println(user)
}

func DoSomething(wg *sync.WaitGroup, oneChan chan string) {
	defer wg.Done()
	oneChan <- "hello, Nikolas"
}
