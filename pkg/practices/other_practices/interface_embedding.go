package other_practices

import "fmt"

// interface embedding and classic basic structure one

func Practice() {
	var myPhone = phone{
		gadget: gadget{
			brand:            "iphone",
			producer:         "USA",
			dateOfProduction: "yesterday",
			status:           false,
		},
	}

	SomeAction(&myPhone)

}

func SomeAction(device Device) {
	device.TurnON()
	myGadget := device.GetInfo()
	fmt.Printf("info: %+v\n", myGadget)
	fmt.Println("producer:", myGadget.producer)
}

type Starter interface {
	TurnON()
}

type Stopper interface {
	TurnOFF()
}

type Device interface {
	Starter
	Stopper
	GetInfo() *gadget
}

func (g *gadget) GetInfo() *gadget {
	return &gadget{
		brand:            g.brand,
		producer:         g.producer,
		dateOfProduction: g.dateOfProduction,
		status:           g.status,
	}
}

func (g *gadget) TurnON() {
	g.status = true
	fmt.Println("device is turned on")
}
func (g *gadget) TurnOFF() {
	g.status = false
	fmt.Println("device is turned off")
}

type (
	gadget struct {
		brand            string
		producer         string
		dateOfProduction string
		status           bool
	}

	phone struct {
		gadget
	}
	computer struct {
		gadget
	}
	tablet struct {
		gadget
	}
)

/* todo play with second point in explaining of gemeni, change data in method, return self pointer and develop
todo the chain of actions with method like GORM
*/
