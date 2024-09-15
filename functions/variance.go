package myfunc

import (
	"fmt"
	"math"
)

func Variance(data []int) {
	count := len(data)



	// Step 1: Calculate the mean
	sum := 0
	for _, value := range data {
		sum += value
	}
	mean := float64(sum) / float64(count)

	// Step 2: Calculate the squared differences from the mean
	sumOfSquaredDiffs := 0.0
	for _, value := range data {
		diff := float64(value) - mean
		sumOfSquaredDiffs += diff * diff
	}

	// Step 3: Calculate the variance
	variance := sumOfSquaredDiffs / float64(count)
	roundedVariance := int64(math.Round(variance))
	// Step 4: Print the result
	fmt.Println("Variance:", roundedVariance)
}