//--Summary:
//  Create a program that can perform dice rolls using various configurations
//  of number of dice, number of rolls, and number of sides on the dice.
//  The program should report the results as detailed in the requirements.
//
//--Requirements:
//* Print the sum of the dice roll
//* Print additional information in these cirsumstances:
//  - "Snake eyes": when the total roll is 2, and total dice is 2
//  - "Lucky 7": when the total roll is 7
//  - "Even": when the total roll is even
//  - "Odd": when the total roll is odd
//* The program must use variables to configure:
//  - number of times to roll the dice
//  - number of dice used in the rolls
//  - number of sides of all the dice (6-sided, 10-sided, etc determined
//    with a variable). All dice must use the same variable for number
//    of sides, so they all have the same number of sides.
//
//--Notes:
//* Use packages from the standard library to complete the project
//* Try using different values for your variables to check the results

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func roll_dice(sides int) int {
	return rand.Intn(sides) + 1
}

func main() {
	rand.Seed(time.Now().UnixNano())

	numRolls := 10
	numSides := 6
	numDice := 2

	for i := 1; i <= numRolls; i++ {
		fmt.Println("Roll number:", i)
		sum := 0

		for j := 1; j <= numDice; j++ {
			roll := roll_dice(numSides)
			fmt.Println("... dice number", j, "=", roll)
			sum += roll
		}

		fmt.Println("Sum of dice:", sum)

		switch sum := sum; {
		case sum == 2 && i == 2:
			fmt.Println("Snake Eyes!")
		case i == 7:
			fmt.Println("Lucky 7!")
		case i%2 == 0:
			fmt.Println("Even!")
		case i%2 != 0:
			fmt.Println("Odd!")
		}

		fmt.Println()

	}
	//* Print the sum of the dice roll
	//* Print additional information in these cirsumstances:

	//* The program must use variables to configure:
	//  - number of times to roll the dice
	//  - number of dice used in the rolls
	//  - number of sides of all the dice (6-sided, 10-sided, etc determined
	//    with a variable). All dice must use the same variable for number
	//    of sides, so they all have the same number of sides.
}
