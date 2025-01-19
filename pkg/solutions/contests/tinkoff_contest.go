package contests

// n - amount of guns	4
// s - money			7
// list := []int{3, 12, 5, 6, 8}
//output: [6]

func SecondTinkoff(money, amountOfGuns int, list []int) int {
	var maxPossiblePrice int
	var priceInList int

	for i := 0; i < amountOfGuns; i++ {
		priceInList = list[i]
		if priceInList <= money {
			if maxPossiblePrice < priceInList {
				maxPossiblePrice = priceInList
			}
		}
	}
	return maxPossiblePrice
}
