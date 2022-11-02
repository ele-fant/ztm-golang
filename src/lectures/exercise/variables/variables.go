//Summary:
//  Print basic information to the terminal using various variable
//  creation techniques. The information may be printed using any
//  formatting you like.
//
//Requirements:
//* Store your favorite color in a variable using the `var` keyword
//* Store your birth year and age (in years) in two variables using
//  compound assignment
//* Store your first & last initials in two variables using block assignment
//* Declare (but don't assign!) a variable for your age (in days),
//  then assign it on the next line by multiplying 365 with the age
// 	variable created earlier
//
//Notes:
//* Use fmt.Println() to print out information
//* Basic math operations are:
//    Subtraction    -
// 	  Addition       +
// 	  Multiplication *
// 	  Division       /

package main

import "fmt"

func main() {
	//Requirements:
	//* Store your favorite color in a variable using the `var` keyword
	var myFavColor = "green"
	fmt.Println(myFavColor)
	//* Store your birth year and age (in years) in two variables using
	//  compound assignment
	type year int
	var birthYear, myAge year = 1990, 32
	fmt.Println("born in", birthYear, ", age:", myAge)
	//* Store your first & last initials in two variables using block assignment
	var (
		firstInitial = "F" // runes need a special formatting to print them as runes
		lastInitial  = "T"
	)
	fmt.Println("first initial:", firstInitial)
	fmt.Println("last initial:", lastInitial)
	//* Declare (but don't assign!) a variable for your age (in days),
	//  then assign it on the next line by multiplying 365 with the age
	// 	variable created earlier
	type days int
	var age days

	age = days(myAge * 365)
	fmt.Println("age in days:", age)
}
