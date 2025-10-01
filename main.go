package main

import "fmt"

func add(num1 int, num2 int) int {

	sum := num1 + num2
	return sum
}

func getNumber(num1 int, num2 int) (int, int) {
	sum := num1 + num2
	mul := num1 * num2

	return sum, mul
}

func printSomething() {
	fmt.Println("Learn something with Habib vai")
}
func printWithString(name string) {
	fmt.Println("Learn go Lang with ", name)
}

func main() {
	fmt.Println("Hello Majid! Started Learning Go")

	sum := add(10, 20)
	fmt.Println(sum)

	a := 20
	b := 30

	s, m := getNumber(a, b)
	fmt.Println(s, m)

	printSomething()
	printWithString("Habib vai")

}
