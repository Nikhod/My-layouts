package solutions

import (
	"Nikcase/pkg/models"
	"errors"
	"fmt"
	"strings"
)

/* data for example checking
list := make(map[string]int)
	list["one"] = 111
	list["two"] = 222
	list["three"] = 333
	list["four"], list["five"], list["six"] = 4444, 5555, 6666

	fmt.Printf("original map: %v\n", list)
	var keys = []string{"seven", "eight"}
	var values = []int{7, 8}


	fmt.Printf("the result map: %v\n", list)
*/

func RemoveFromMap(table map[string]int, keys ...string) {
	fmt.Printf("original map: %v\n", table)
	for index := range keys {
		delete(table, keys[index])
	}
	fmt.Printf("the result map: %v\n", table)
}

func AddToMap(table map[string]int, keys []string, values []int) error {
	if len(keys) != len(values) {
		return errors.New("the quantity of keys and values don't match")
	}
	for i := 0; i < len(keys); i++ {
		table[keys[i]] = values[i]
	}

	return nil
}

func CountAmountOfWord(text, word string) int {
	text = strings.ToLower(text)
	text = strings.Replace(text, ",", "", -1)

	splitedText := strings.Split(text, " ")
	counter := make(map[string]int)
	counter[word] = 0

	for _, value := range splitedText {
		if value == word {
			counter[word] = counter[word] + 1
		}
	}
	return counter[word]
}

func CountUniqueValueInSlice(sliceOfStrings []string) []string {
	counter := make(map[string]int)
	var resultSlice = make([]string, 2, 2)

	for _, value := range sliceOfStrings {
		counter[value] = counter[value] + 1
	}

	for word := range counter {
		if counter[word] == 1 {
			resultSlice = append(resultSlice, word)
		}
	}

	return resultSlice
}

func UniteTwoMaps(firstMap, secondMap map[int]string) map[int]string {
	if len(firstMap) < len(secondMap) {
		for key, value := range secondMap {
			firstMap[key] = value

		}
		return firstMap

	} else {
		for key, value := range firstMap {
			secondMap[key] = value
		}
		return secondMap
	}

}

func UniteAllMaps(maps ...map[string][]string) (unitedMap map[string][]string) {
	unitedMap = make(map[string][]string)
	for i := 0; i < len(maps); i++ {
		for key, value := range maps[i] {
			theSlice := unitedMap[key]
			if _, ok := unitedMap[key]; ok {
				theSlice = append(theSlice, value...)
				unitedMap[key] = theSlice
			} else {
				unitedMap[key] = value
			}
		}
	}
	return unitedMap
}

func ChangeConfigServer(config models.Server) models.Server {
	var resultServerConfig models.Server
	resultServerConfig.Host = config.Host + "_new"
	resultServerConfig.Port = config.Port + "_new"

	return resultServerConfig
}

func FindSquare(number int) int {
	return number * number
}
