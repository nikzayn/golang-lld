package models

import "time"

//Define a Parking lot with levels embedded
type ParkingLot struct {
	Levels []*Level
}

//Define a Level with Parking spot embedded
type Level struct {
	Spots []*ParkingSpot
}

//Define a Parking Spot according to size, availability and vehicle details
type ParkingSpot struct {
	Id          int
	Size        int
	IsAvailable bool
	Vehicle     *Vehicle
}

//Define a vehicle size, entryTime and exitTime
type Vehicle struct {
	Size      VehicleSize
	EntryTime time.Time
	ExitTime  time.Time
	LicenseNo int
}

//Define a vehicle size
type VehicleSize int

const (
	Motorcycle VehicleSize = iota
	Compact
	Large
)
