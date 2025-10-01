package main

import "fmt"

func add(num1 int, num2 int) int {

	sum := num1 + num2
	return sum
}

func main() {
	fmt.Println("Hello Majid! Started Learning Go")

	sum := add(10, 20)
	fmt.Println(sum)
}
