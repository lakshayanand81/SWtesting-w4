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
func reverseString(input string) string {
	runes := []rune(input)
	length := len(runes)

	for i := 0; i < length/2; i++ {
		runes[i], runes[length-i-1] = runes[length-i-1], runes[i]
	}

	return string(runes)
}

func reverse() {
	original := "MAYANK"
	reversed := reverseString(original)
	fmt.Println("Original:", original)
	fmt.Println("Reversed:", reversed)
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
        fmt.Println(reverse())
	addLetters()
}


// fibonacci returns a function that generates Fibonacci numbers.
func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	// Create a Fibonacci number generator function.
	fib := fibonacci()

	// Generate and print the first 10 Fibonacci numbers.
	for i := 0; i < 10; i++ {
		fmt.Println(fib())
	}
}

