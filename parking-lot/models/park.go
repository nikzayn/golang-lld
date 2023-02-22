package models

import (
	"errors"
	"time"
)

/*
 * Description - Park a vehicle and return the parking spot
  * @Input() - v -> Vehicle struct
  * @Output() -> *ParkingSpot, error
  * TC - O(N^2) | SC - O(1)
*/
func (p *ParkingLot) Park(v *Vehicle) (*ParkingSpot, error) {
	for _, level := range p.Levels {
		for _, spot := range level.Spots {
			if spot.IsAvailable && spot.Size >= int(v.Size) {
				spot.IsAvailable = false
				spot.Vehicle = v
				return spot, nil
			}
		}
	}
	return nil, errors.New("parking lot is full")
}

/*
 * Description - Unpark a vehicle and return the parking spot
  * @Input() - spot -> ParkingSpot struct
  * @Output() -> error
  * TC - O(1) | SC - O(1)
*/
func (p *ParkingLot) Unpark(spot *ParkingSpot) error {
	if spot == nil {
		return errors.New("spot is invalid")
	}
	spot.IsAvailable = true
	return nil
}

/*
 * Description - Checkout system will generate ticket fare for a particular vehicle
  * @Input() - spot -> ParkingSpot struct
  * @Output() -> float64, error
  * TC - O(1) | SC - O(1)
*/
func (p *ParkingLot) Checkout(spot *ParkingSpot) (float64, error) {
	//If spot is nil or is it available
	if spot == nil || spot.IsAvailable {
		return 0, errors.New("spot is invalid")
	}
	spot.IsAvailable = true
	v := spot.Vehicle
	spot.Vehicle = nil
	fee := p.calculateFare(v)
	return fee, nil
}

/*
 * Description - Calculate the fare charges accordingly to the vehicle size
  * @Input() - v -> Vehicle struct
  * @Output() -> float64
  * TC - O(1) | SC - O(1)
*/
func (p *ParkingLot) calculateFare(v *Vehicle) float64 {
	duration := time.Since(v.EntryTime)
	switch v.Size {
	case Motorcycle:
		return duration.Minutes() * 10
	case Compact:
		return duration.Minutes() * 20
	case Large:
		return duration.Minutes() * 30
	}
	return 0.0
}
