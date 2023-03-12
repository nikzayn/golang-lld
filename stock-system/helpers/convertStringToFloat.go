package helpers

import "fmt"

func ConvertStringToFloat(str string) float64 {
	var f float64
	fmt.Sscanf(str, "%f", &f)
	return f
}
