package models

type Rating struct {
	ListOfComments []Comment
	ListOfRatings  []int
	TotalRatingSum int
}
