package controllers

import models "github.com/nikzayn/golang-lld/parking-lot/models"

/*
 * Description - To create new parking lot
  * @Input() - numLevels integer, numSportsPerLevel - integer
  * @Output() - ParkingLot struct
  * TC - O(1) | SC - O(1)
*/
func NewParkingLot(numLevels, numSpotsPerLevel int) *models.ParkingLot {
	//Initialize a levels map
	levels := make([]*models.Level, numLevels)
	//Traverse over no of levels
	for i := 0; i < numLevels; i++ {
		levelSize := models.Large
		if i%2 == 0 {
			levelSize = models.Compact
		}
		//update each level with numSpotsPerLevel and levelSize
		levels[i] = NewLevel(numSpotsPerLevel, levelSize)
	}
	return &models.ParkingLot{Levels: levels}
}
