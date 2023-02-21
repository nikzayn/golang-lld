package controllers

import (
	models "github.com/nikzayn/golang-lld/parking-lot/models"
)

/*
 * Description - To create n number of paring spots
  * @Input() - id -> integer, numSports -> integer, size -> integer
  * @Output() - Level struct
  * TC - O(N) | SC - O(N)
*/
func NewLevel(numSpots int, size models.VehicleSize) *models.Level {
	//Initialize a spot map
	spots := make([]*models.ParkingSpot, numSpots)
	//Traverse over no. of spots
	for i := 0; i < numSpots; i++ {
		//update spots with new parking spot
		spots[i] = NewParkingSpot(i, size)
	}
	return &models.Level{Spots: spots}
}
