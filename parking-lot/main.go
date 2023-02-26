package main

import (
	"fmt"

	"github.com/nikzayn/golang-lld/parking-lot/controllers"
)

func main() {
	numberOfLevels := 5
	spotsPerLevel := 10

	//Created a parking lot with number of levels and number of spot per level
	generateParkingLot := controllers.NewParkingLot(numberOfLevels, spotsPerLevel)
	generateLevels := controllers.NewLevel(spotsPerLevel, 2)

	fmt.Printf("%v %v", generateParkingLot, *generateLevels)
}
