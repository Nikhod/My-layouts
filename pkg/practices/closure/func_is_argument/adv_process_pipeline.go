package func_is_argument

import "log"

const (
	squareIndex      = 0
	additionIndex    = 1
	subtractionIndex = 2
	multiplyIndex    = 3
)

func processPipeline(nums []int, handlers []func(int) int) []ordinaryPipeline {
	var result []ordinaryPipeline

	for i := 0; i < len(nums); i++ {
		node := ordinaryPipeline{
			square:           handlers[squareIndex](nums[i]),
			additionNum:      handlers[additionIndex](nums[i]),
			subtraction:      handlers[subtractionIndex](nums[i]),
			multiplyBy100000: handlers[multiplyIndex](nums[i]),
		}

		result = append(result, node)
	}
	return result
}

func ProcessPipelinePractice() {
	store := []int{12, 67, 243, 323, 23, 4, 32, 43}
	newHandlers := []func(int) int{
		func(num int) int { return num * num },
		func(num int) int { return num + 100_000 },
		func(num int) int { return num - 100_000 },
		func(num int) int { return num * 100_000 },
	}

	result := processPipeline(store, newHandlers)

	log.Printf(">>>>:%+v\n", result[0])

}

type ordinaryPipeline struct {
	square           int
	additionNum      int
	subtraction      int
	multiplyBy100000 int
}
