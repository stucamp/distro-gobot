package main

import (
	"math/rand"
	"time"
)

// GetRandNum returns a random num from 0 - 100
func GetRandNum() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(100)
}
