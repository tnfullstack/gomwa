package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, World!")

	var whatToSay string
	var age int // int64

	fName := "Thanh Nguyen"
	currentYear := time.Now().Year()
	birthYear := 1970
	age = currentYear - birthYear

	whatToSay = "Hello"

	fmt.Println(whatToSay, fName+"!")

	fmt.Printf("Hello %s, are you %d years old?\n", fName, age)

	str, num := saySomething()
	fmt.Printf("Say what? %s and %d\n", str, num)
}

func saySomething() (string, int) {
	return "Something!", 10
}
