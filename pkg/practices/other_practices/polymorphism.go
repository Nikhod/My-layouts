package other_practices

import (
	"fmt"
	"log"
)

/*
Полиморфизм - это один из принципов ООП, который гласит,  и он позволяет объектам разных типов быть использованными
как объекты одного типа. Это значит, что по основному свойству Интерфейсов - это контракт где описаны действия
структур. Те структуры которые имеют такой метод как в интерфейсе (с такой же сигнатурой), это значит что эти
структуры можно объединить в один целый тип. Интерфейс передается в func(i interface), но так же func()
могут возвращать интерфейс (Exm: error - кастом эррор в этом же пакете реализован). В интерфейсе могут быть описаны
другие интерфейсы. Утиная типизация связан с полиморфизмом: "duck typing" ("если это ходит как утка и крякает
как утка, то это, вероятно, утка"). Это означает, что тип определяется не по имени, а по наличию определенных методов.
Если тип реализует нужные методы, он удовлетворяет интерфейсу.
*/
func PolymorphismPractice() {
	onlinePayment := alifMobi{
		balance: balance{
			money: 5000,
		},
		number: "44448888",
	}

	cardPayment := creditcard{
		number:     "1234567",
		expiration: "11-23-2027",
		pass:       "cat_dog",
		balance: balance{
			money: 5000,
		},
	}

	payment := []PaymentMethod{&onlinePayment, &cardPayment}
	err := payment[1].pay(1000)
	if err != nil {
		log.Println(err)
		return
	}
}

type PaymentMethod interface {
	pay(amount int) error
}

// one can develop the individual method for 'balance' and embedding to other structures
type (
	//правильнее будет, если добавить методы (ругулирующие финансы, платеж, поступление итд) к этой структуре, чтобы
	//легче редактировать код и масштабировать, а не создаавать отельные структуры для структур методов оплаты.
	//Это классно выразит идею инкапсуляции и встраивания
	balance struct {
		money int
	}

	customErr struct {
		message    string
		statusCode int
	}

	creditcard struct {
		number     string
		expiration string
		pass       string
		balance
	}

	alifMobi struct {
		balance
		number string
	}
)

// there are differences between method "pay" of creditcard and alifMobdi structures.

func (c *creditcard) pay(amount int) error {
	walletWithMoney := &c.balance.money
	if amount > *walletWithMoney {
		return customErr{
			message:    "you have no money to pay",
			statusCode: 500,
		}
	}

	*walletWithMoney -= amount
	fmt.Println("payment with CREDIT CARD is processing....")
	fmt.Printf("%v dollars in wallet\nc.balance.money = %v\n", *walletWithMoney, c.balance.money)
	fmt.Println(*walletWithMoney == c.balance.money)
	fmt.Printf("residue: %v\n", c.balance.money)
	return nil
}

func (a *alifMobi) pay(amount int) error {
	walletWithMoney := &a.balance.money
	if amount > *walletWithMoney {
		return customErr{
			message:    "you have no money to pay",
			statusCode: 500,
		}
	}

	*walletWithMoney -= amount
	fmt.Println("alif mobi payment is processing")
	fmt.Printf("%v dollars in wallet\nc.balance.money = %v\n", *walletWithMoney, a.balance.money)
	fmt.Println(*walletWithMoney == a.balance.money)
	fmt.Printf("residue: %v\n", a.balance.money)
	return nil
}

// embedded interface "error" implements method Error() string - custom Error
func (c customErr) Error() string {
	return fmt.Sprintf("error message: %v\nstatus code: %v\n",
		c.message, c.statusCode)
}

// other interface topics for learning
// -- Пустой интерфейс interface{}
// -- Утверждение типа (Type Assertion)
// -- Переключение типа (Type Switch) (есть в тетрадке, посмотри)
