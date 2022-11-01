package main

import "fmt"

func main() {
	var myName = "Fabio"
	fmt.Println("my name is", myName)

	var name string = "Kathy"
	fmt.Println("name =", name)

	userName := "admin"
	fmt.Println("username =", userName)

	var sum int
	fmt.Println("the sum is", sum)

	part1, other := 1, 5
	fmt.Println("part1 is", part1, "; other is", other)

	part2, other := 1, 0
	fmt.Println("part1 is", part2, "; other is", other)

	sum = part1 + part2
	fmt.Println("sum =", sum)

	var (
		lessonName = "Variables"
		lessonType = "Demo"
	)
	fmt.Println("lesson name:", lessonName)
	fmt.Println("lesson type:", lessonType)

	word1, word2, _ := "hello", "world", "!"
	fmt.Println(word1, word2)
}
