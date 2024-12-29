package other_practices

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func DecodeSeveralJsonObjects() {
	var person []persons
	queue := map[int]string{
		1: "FIRST",
		2: "SECOND",
		3: "THIRD",
		4: "FOURTH",
		5: "FIFTH",
	}

	bytesFromFile, err := os.ReadFile("example.json")
	if err != nil {
		log.Println(err)
		return
	}

	reader := bytes.NewReader(bytesFromFile)
	err = json.NewDecoder(reader).Decode(&person)
	if err != nil {
		log.Println(err)
		return
	}

	for following, value := range person {
		fmt.Printf("the %s person: %+v\n", queue[following+1], value)
	}
}

type persons struct {
	Id       int
	Name     string
	Login    string
	Password string
	Active   bool
	Token    string
}
