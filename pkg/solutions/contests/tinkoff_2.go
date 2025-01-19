package contests

import (
	"strings"
	"unicode"
)

// сколько раз в строке substring встречаются слово "sheriff". Регистр и пробелы игнорирутся
func FindAmountOfUsageOfString(substring string) int {
	sheriffInMap := map[rune]int{
		's': 0,
		'h': 0,
		'e': 0,
		'r': 0,
		'i': 0,
		'f': 0,
	}

	//substring = "sheriff sheriff sheriff"
	substring = strings.ToLower(substring)

	for i := 0; i < len(substring); i++ {
		charRuneInString := rune(substring[i])

		if unicode.IsSpace(charRuneInString) {
			continue
		}
		if value, ok := sheriffInMap[charRuneInString]; ok {
			sheriffInMap[charRuneInString] = value + 1
		}
	}

	sheriffInMap['f'] /= 2
	return findMinValue(sheriffInMap)
}

func findMinValue(sheriffMap map[rune]int) int {
	var min int = sheriffMap['s']

	for _, value := range sheriffMap {
		if value < min {
			min = value
		}
	}
	return min
}
