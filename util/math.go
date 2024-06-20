package util

import (
	"math/rand"
)

func RandInt(min, max int) int {
	return min + rand.Intn(max-min)
}
