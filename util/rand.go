package util

import (
	"math/rand"
	"time"
)

func GetRand(min int64, max int64) int64 {
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Int63n(max - min)
	randNum = randNum + min
	return randNum
}
