package ParkingLot

import "time"

//Define a Parking lot with levels embedded
type ParkingLot struct {
	levels []*Level
}

//Define a Level with Parking spot embedded
type Level struct {
	spots []*ParkingSpot
}

//Define a Parking Spot according to size, availability and vehicle details
type ParkingSpot struct {
	id          int
	size        int
	isAvailable bool
	vehicle     *Vehicle
}

//Define a vehicle size, entryTime and exitTime
type Vehicle struct {
	size      VehicleSize
	entryTime time.Time
	exitTime  time.Time
}

//Define a vehicle size
type VehicleSize int

const (
	Motorcycle VehicleSize = iota
	Car
	Truck
)
