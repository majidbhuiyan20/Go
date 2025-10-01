package main

import "fmt"

func startLearning() {
	fmt.Println("Hello Majid! Started Learning Go")
}

func add(num1 int, num2 int) int {
	sum := num1 + num2
	fmt.Println("Sum is = ", sum)
	return sum
}

func getNumber(num1 int, num2 int) (int, int) {
	sum := num1 + num2
	mul := num1 * num2
	fmt.Println("Sum is = ", sum)
	fmt.Println("Mul is = ", mul)
	return sum, mul
}

func printSomething() {
	fmt.Println("Learn something with Habib vai")
}

func printWithString(name string) {
	fmt.Println("Learn go Lang with ", name)
}

func nameInput() string {
	var name string
	fmt.Print("Enter Your Name - ")
	fmt.Scan(&name)
	fmt.Println("Your name is ", name)
	return name
}

func getTwoNumber() (int, int, int) {
	var num1 int
	var num2 int
	fmt.Print("Enter first number - ")
	fmt.Scan(&num1)
	fmt.Print("Enter Second number - ")
	fmt.Scan(&num2)
	sum := num1 + num2
	fmt.Println("Sumation is - ", sum)
	return num1, num2, sum
}

func main() {
	startLearning()
	add(10, 20)
	getNumber(12, 12)
	printSomething()
	printWithString("Habib vai")
	nameInput()
	getTwoNumber()

}
