package main

import (
	"fmt"
)

func numWaterBottles(numBottles int, numExchange int) int {
    var emptyBottles int
    var result int

    // first drink
    emptyBottles = numBottles
    result = numBottles

    // Exchange
    for emptyBottles >= numExchange {
		// exchange
        numBottles = emptyBottles / numExchange
		emptyBottles = emptyBottles % numExchange

		// drink
		emptyBottles += numBottles
		result += numBottles
    }
	return result
}

func main() {
	var numBottles int = 9
	var numExchange int = 3

	fmt.Println(numWaterBottles(numBottles, numExchange))
}

// Độ phức tạp thuật toán: O(logN)
// Độ phức tạp lưu trữ: O(1)