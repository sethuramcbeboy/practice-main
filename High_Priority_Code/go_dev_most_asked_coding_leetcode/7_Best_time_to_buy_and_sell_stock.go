package godevmostaskedcoding

import (
	"fmt"
	"log"
)

func maxProfit(prices []int) int {

	minPrice := prices[0]
	maxProfit := 0

	for _, p := range prices {
		if p < minPrice {
			minPrice = p
		}
		if p-minPrice > maxProfit {
			maxProfit = p - minPrice
		}
		log.Println(minPrice, maxProfit)
	}
	return maxProfit
}

func max_main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}
