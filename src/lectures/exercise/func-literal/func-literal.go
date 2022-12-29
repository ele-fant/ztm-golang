//--Summary:
//  Create a program that can create a report of rune information from
//  lines of text.
//
//--Requirements:
//* Create a single function to iterate over each line of text that is
//  provided in main().
//  - The function must return nothing and must execute a closure
//* Using closures, determine the following information about the text and
//  print a report to the terminal:
//  - Number of letters
//  - Number of digits
//  - Number of spaces
//  - Number of punctuation marks
//
//--Notes:
//* The `unicode` stdlib package provides functionality for rune classification

package main

import (
	"fmt"
	"unicode"
)

func iterate(lines []string, op func(r rune) int) {
	sum := 0
	for _, line := range lines {
		for _, l := range line {
			sum += op(l)
		}
	}
	fmt.Printf("%v", sum)
}

func main() {
	lines := []string{
		"There are",
		"68 letters,",
		"five digits,",
		"12 spaces,",
		"and 4 punctuation marks in these lines of text!",
	}

	numLetters := func(r rune) int {
		return (unicode.IsLetter(r))
	}

	fmt.Printf("Letters - ")
	iterate(lines, numLetters)
	//  - Number of digits
	fmt.Printf("Digits - ")
	iterate(lines, func(r rune) int {
		return unicode.IsDigit(r)
	})
	//  - Number of spaces
	fmt.Printf("Letters - ")
	iterate(lines, numLetters)
	//  - Number of punctuation marks
	fmt.Printf("Letters - ")
	iterate(lines, numLetters)
}
