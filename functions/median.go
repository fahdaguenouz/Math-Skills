package myfunc

import (
	"fmt"
	"math"
	
)


func Median(data []int) {
	count := len(data)


	// Sort the data
	for i := 0; i < count-1; i++ {
		for j := 0; j < count-i-1; j++ {
			if data[j] > data[j+1] {
				// Swap the elements
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}

	var median float64

	if count%2 == 1 {
		// Odd number of elements: middle element is the median
		median = float64(data[count/2])
	} else {
		// Even number of elements: average of the two middle elements
		mid1 := data[(count/2)-1]
		mid2 := data[count/2]
		median = float64(mid1+mid2) / 2.0
	}

	// Round the median
	roundedMedian := math.Round(median)

	// Print the rounded median
	fmt.Println("Median:", roundedMedian)
}