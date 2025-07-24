package func_is_argument

import (
	"fmt"
	"log"
)

func applyToEach(nums []int, action func(num int) bool) {
	//здесь может быть НЕБОЛЬШАЯ логика привидения типов, не более
	for i := 0; i < len(nums); i++ {
		isValid := action(nums[i])
		if isValid {
			log.Println("<<<четное>>>>", isValid)
		}

	}

	log.Println("after calling action func")
}

// find the square of num and log to console
func findSquare(nums []int) {
	for _, v := range nums {
		log.Println(v * v)
	}

	fmt.Println("\n\n")
}

func ApplyToEachElementPractice() {
	store := []int{12, 67, 243, 323, 23, 4, 32, 43}

	applyToEach(store, func(num int) bool {
		if num%2 == 0 {
			return true
		} else {
			return false
		}
	})
}

//❗Рекомендации по стилю:
//Старайся, чтобы applyToEach не делала лишнего, например, findSquare лучше вызвать отдельно:
//Не пихай логику внутрь applyToEach, если она может быть сделана в вызывающем коде. Это делает функцию
//гибче и переиспользуемой.
//
