package utils

import (
	"math/rand"
)

// GetRandomInt returns an integer from 0 to the number - 1
func GetRandomInt(num int) int {
	x := rand.Intn(num)
	return x
}

func GetRandomFlt() float64 {
	return rand.Float64()
}

// GetDiceRoll returns an integer from 1 to the number
func GetDiceRoll(num int) int {
	x := rand.Intn(num)
	return x + 1
}

func GetRandomBetweenTwo(low int, high int) int {
	return GetDiceRoll(high-low) + low
}
