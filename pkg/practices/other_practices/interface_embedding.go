package other_practices

import "fmt"

// interface embedding and classic basic structure one

func FirstInterfaceEmbeddingPractice() {
	myComputer := computer{
		gadget: gadget{
			brand:            "McBook",
			producer:         "Taiwan",
			dateOfProduction: "yesterday",
			status:           false,
		},
	}
	myTablet := tablet{
		gadget: gadget{
			brand:            "ipad",
			producer:         "FAXCON",
			dateOfProduction: "yesterday",
			status:           false,
		},
	}
	myPhone := phone{
		gadget: gadget{
			brand:            "iphone",
			producer:         "USA",
			dateOfProduction: "yesterday",
			status:           false,
		},
	}

	myDevices := []Devices{
		&myPhone,
		&myComputer,
		&myTablet,
	}

	//one can use the cycle "for"
	SomeAction(myDevices[1])

}

func SomeAction(device Devices) {
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

type Devices interface {
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
