package main

import "fmt" 
       "strings"


func addLetters() {
	letter1 := "A"
	letter2 := "B"
	result := letter1 + letter2
	fmt.Printf("The result of adding letters %s and %s is: %s\n", letter1, letter2, result)
}

func pallindrome() {
	var input string

	fmt.Print("Enter a string: ")
	fmt.Scanln(&input)

	// Convert the string to lowercase and remove spaces
	input = strings.ToLower(input)
	input = strings.ReplaceAll(input, " ", "")

	// Compare the string with its reverse
	isPalindrome := true
	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-1-i] {
			isPalindrome = false
			break
		}
	}

	if isPalindrome {
		fmt.Println("The input is a palindrome.")
	} else {
		fmt.Println("The input is not a palindrome.")
	}
}


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

	addLetters()
}
