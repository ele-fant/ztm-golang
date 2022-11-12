package main

import (
	"fmt"
	"strings"
)

func main() {
	myString := "Hello"
	fmt.Println(strings.Contains(myString, "he"))
	fmt.Println(strings.Contains(myString, "He"))
}
