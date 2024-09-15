package myfunc

import (
	"fmt"
	"math"
)

func Average(data []int) {
	sum := 0
	count := len(data)

	// Calculate the sum of the numbers
	for _, num := range data {
		sum += num
	}

	// Calculate the average
	average := float64(sum) / float64(count)
	roundedAverage := math.Round(average)
	fmt.Println("Average:", roundedAverage)
}