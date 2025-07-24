package func_is_argument

import "log"

func FilterPractice() {
	var result []int
	nums := []int{12, 20, 13, 5, 100}

	for i := 0; i < len(nums); i++ {
		isValid := filterWithFilterFunc(nums[i], isChet)

		// если число четное то добавим его в слайс
		if isValid {
			result = append(result, nums[i])
		}

	}
	log.Println(result)

}
func filterWithFilterFunc(num int, predicate func(isValid int) bool) bool {
	// logic before the calling func
	// тут может быть логика до вызова функции, но обычно, в случае с фильтрами эта логика может либо отсутствовать,
	// либо это может быть логика связананная с приведением данных
	return predicate(num)
}

// FILTERS:

func isChet(num int) bool {
	if num%2 == 0 {
		return true
	} else {
		return false
	}
}
