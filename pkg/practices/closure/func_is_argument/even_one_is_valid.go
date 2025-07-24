package func_is_argument

import "log"

func AnyElement(nums []int, predicate func(nums []int) bool) {
	isValid := predicate(nums)
	log.Println(isValid)
}

func AnyPractice() {
	nums := []int{12, 4, 2, 44522, 5, 676, 878, -1}
	AnyElement(nums, func(nums []int) bool {
		for _, v := range nums {
			if v < 0 {
				return true
			}
		}
		return false
	})

}
