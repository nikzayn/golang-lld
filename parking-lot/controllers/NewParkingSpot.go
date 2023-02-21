package controllers

import models "github.com/nikzayn/golang-lld/parking-lot/models"

/*
 * Description - To create new parking spot
  * @Input() - id integer, size - integer
  * @Output() - ParkingSpot struct
  * TC - O(1) | SC - O(1)
*/
func NewParkingSpot(id int, size models.VehicleSize) *models.ParkingSpot {
	return &models.ParkingSpot{Id: id, Size: int(size), IsAvailable: true}
}
