package main

import (
	"fmt"
)

// Function to check if a number is even or odd
func checkEvenOdd(n int) {
	if n%2 == 0 {
		fmt.Println(n, "is even")
	} else {
		fmt.Println(n, "is odd")
	}
}

// Function to calculate the sum of numbers up to n using a for loop
func sumUpTo(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

// Function to modify a number using a pointer
func doubleNumber(num *int) {
	*num = *num * 2
}

func main() {
	// Defer statement (executes at the end)
	defer fmt.Println("Exiting program. Goodbye!")

	// Get user input
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	// Using if statement
	checkEvenOdd(num)

	// Using for loop
	fmt.Println("Sum of numbers up to", num, "is:", sumUpTo(num))

	// Using pointers
	fmt.Println("Doubling the number...")
	doubleNumber(&num)
	fmt.Println("Doubled value:", num)
}
