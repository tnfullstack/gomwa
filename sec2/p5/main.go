package main

import (
	"log"
	"time"
)

type Car struct {
	Make      string
	Model     string
	VIN       string
	Type      string
	YearBuild time.Time
}

func (c *Car) printCarMake() string {
	return c.Make
}

func (c *Car) printCarModel() string {
	return c.Model
}

func main() {
	var myCar Car
	myCar.Make = "Toyota"
	myCar.Model = "2020 Tundra TRD Pro"
	myCar.Type = "Full Size Pickup Truck"

	log.Println("My car's make is:", myCar.printCarMake())
	log.Println("My car's model is", myCar.printCarModel())
}
