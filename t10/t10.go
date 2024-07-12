package main

import (
	"fmt"
	"math"
)

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	mapTemps := make(map[int][]float64)

	for _, temp := range temperatures {
		key := int(math.Floor(temp/10) * 10)
		mapTemps[key] = append(mapTemps[key], math.Floor(temp*2+0.5)/2)
	}
	for key, value := range mapTemps {
		fmt.Printf("Значение: %d:%.1f\n", key, value)
	}
}
