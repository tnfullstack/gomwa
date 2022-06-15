package main

import "log"

func main() {
	var myString string
	myString = "Green"

	log.Println("myString is set to green", myString)

	updateValueByPointer(&myString)

	log.Println("After call to updateValueByPointer(s *string), myString is now = ", myString)
}

func updateValueByPointer(s *string) {
	*s = "Red"
}
