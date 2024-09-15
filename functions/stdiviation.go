package myfunc

import (
	"fmt"
	"math"
)

func StandardDeviation(data []int) {
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

	// Step 4: Calculate the standard deviation (square root of variance)
	standardDeviation := math.Sqrt(variance)
	roundedStandardDeviation := int64(math.Round(standardDeviation))

	// Step 5: Print the result rounded to an integer
	fmt.Println("Standard Deviation:", roundedStandardDeviation)
}