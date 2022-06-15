package main

import "fmt"

type User struct {
	FName string
	LName string
}

func main() {

	// Standard way to create a map
	carMap := make(map[string]string)

	// Adding data to map
	carMap["Make"] = "Toyota"
	carMap["Model"] = "2020 Tundra TRD Pro"
	carMap["Price"] = "55,000.00"
	fmt.Println(carMap["Model"] + ", " + carMap["Make"])

	// map can also hold composite data type
	peopleMap := make(map[string]User)
	peopleMap["Person1"] = User{"Thanh", "Nguyen"}
	peopleMap["Person1"] = User{"Thanh", "Nguyen"}
	peopleMap["Person2"] = User{"Mike", "McNillin"}

	fmt.Printf("Person1's first name is %s\n", peopleMap["Person1"].FName)
	fmt.Printf("Person2's last name is %s\n", peopleMap["Person2"].LName)
}
