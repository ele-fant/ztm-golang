//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import "fmt"

// --Requirements:
//   - Write a function that accepts a person's name as a function
//     parameter and displays a greeting to that person.
func greetPerson(name string) {
	fmt.Println("Hello,", name, "!")
}

//   - Write a function that returns any message, and call it from within
//     fmt.Println()
func returnAnyMessage(message string) string {
	return message
}

//   - Write a function to add 3 numbers together, supplied as arguments, and
//     return the answer
func addThreeNumbers(num1, num2, num3 int) int {
	return num1 + num2 + num3
}

// * Write a function that returns any number
func returnAnyNumber() int {
	return 4
}

// * Write a function that returns any two numbers
func returnAnyTwoNums() (int, int) {
	return 1, 6
}

func main() {
	greetPerson("Fabio")

	fmt.Println(returnAnyMessage("Random message..."))

	sumOfTriple := addThreeNumbers(1, 2, 3)
	fmt.Println("Sum of triple is:", sumOfTriple)

	anyNumber := returnAnyNumber()
	fmt.Println("Random number:", anyNumber)

	a, b := returnAnyTwoNums()
	fmt.Println(a, b)

	//* Add three numbers together using any combination of the existing functions.
	//  * Print the result
	anotherThree := addThreeNumbers(returnAnyNumber(), a, b)
	fmt.Println(anotherThree)

	//* Call every function at least once
}
