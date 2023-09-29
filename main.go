package main

import "fmt"

func main() {
	// Initialize a variable to store the highest number
	highest := 1

	// Loop from 2 to 20
	for i := 2; i <= 20; i++ {
		// Check if the current number is greater than the current highest
		if i > highest {
			highest = i
		}
	}

	// Print the highest number
	fmt.Println("The highest number from 1 to 20 is:", highest)
}
