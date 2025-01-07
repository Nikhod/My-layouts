package other_practices

import (
	"fmt"
	"time"
)

// there is an example work of encapsulation and  embedding

func ChainExecution() {
	pet := NewPet()
	pet.SetMaster("Nikolas", "Baker street, 55").
		SetAge(23).
		SetAnimal("dog").
		SetAnimalName("Rex").
		SetBreed("Alibi")

	pet.GetInfo()
}

type (
	// the names of this structure is hidden. The access to these fields can be taken by the methods
	Pet struct {
		animalName  string
		age         int
		dateOfBirth time.Time
		breed       string
		animal      string
		master
	}

	master struct {
		name    string
		address string
	}
)

func NewPet() *Pet { return &Pet{} }

func (p *Pet) NewMaster(name, address string) *Pet {
	p.master.name = name
	p.address = address
	return p
}

func (p *Pet) SetAnimalName(name string) *Pet {
	p.animalName = name
	return p
}

func (p *Pet) SetAnimal(typeOf string) *Pet {
	p.animal = typeOf
	return p
}

func (p *Pet) SetMaster(name, address string) *Pet {
	p.master.name = name
	p.master.address = address
	return p
}

func (p *Pet) SetAge(age int) *Pet {
	p.age = age
	return p
}

func (p *Pet) SetBreed(breed string) *Pet {
	p.breed = breed
	return p
}

func (p *Pet) GetInfo() {
	fmt.Printf("animalName: %v\nage: %v\nbreed: %v\nanimal: %v\nmster name: %v\naddress: %v\n",
		p.animalName, p.age, p.breed, p.animal, p.master.name, p.address)
}
