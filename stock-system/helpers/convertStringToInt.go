package helpers

import "fmt"

func ConvertStringToInt(str string) int {
	var i int
	fmt.Sscanf(str, "%d", &i)
	return i
}
