package other_practices

import "fmt"

/*
В Go можно встраивать один тип в другой. Это позволяет получить доступ к полям и методам встроенного типа напрямую
через внешний тип. Это часто называют "наследованием" в Go, хотя технически это композиция.

PickUp наследует от Car, таким образом имеет доступ к Методам и полям структуры PICKUP. Будь осторожен с названиями
полей и Методов
*/
func FirstEmbeddingPractice() {
	myPickUp := PickUp{
		WeightCapacity: 3000,
		HasTrailer:     true,
		Car: Car{
			Power:    1000,
			Brand:    "dodge",
			Model:    "trx",
			Producer: "USA",
			Cistern:  0,
		},
	}

	BasicAct(&myPickUp)
}

// funciton with interface in signature
func BasicAct(myPickUp BasicActionOfMachine) {
	myPickUp.GetFuel(133)
	myPickUp.Drive(100)
}

type BasicActionOfMachine interface {
	Drive(litres int)
	GetFuel(litres int)
}

func (c *Car) Drive(litres int) {
	c.Cistern -= litres
	fmt.Println("Cistern is being gotten empty on", litres, "litres")
}

func (c *Car) GetFuel(litres int) {
	c.Cistern += litres
	fmt.Println("Cistern is being filled on", litres, "litres")
}

type (
	Car struct {
		Power    int
		Brand    string
		Model    string
		Producer string
		Cistern  int
	}

	PickUp struct {
		Car
		WeightCapacity int
		HasTrailer     bool
	}
)
