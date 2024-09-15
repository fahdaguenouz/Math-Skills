package main

import (
	"bufio"
	"fmt"
	"mathskills/functions"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args[1:]

	if len(args) == 1{
		input := args[0]
		

		file, err := os.Open(input)
		if err != nil {
			fmt.Println("Error opening the file:", err)
			return
		}
		defer file.Close()

		var numbers []int
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			
			num, err := strconv.Atoi(line)
			if err != nil {
				fmt.Println("Skipping invalid number:", line)
				continue
			}
			// Append the valid numbers to the slice
			numbers = append(numbers, num)
		}
		// Check for scanning errors
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading the file:", err)
			return
		}

		myfunc.Average(numbers)
		myfunc.Median(numbers)
		myfunc.Variance(numbers)
		myfunc.StandardDeviation(numbers)





	} else if len(args) > 2 {
		fmt.Println("Too much arguments")
		os.Exit(1)
	} else if len(args) < 1 {
		fmt.Println("Less arguments please enter the data name")
		os.Exit(1)
	}
}