package main

import (
	"math/rand"
	"time"
)

// GetRandNum returns a random num from 0 - 100
func GetRandNum() string {
	rand.Seed(time.Now().Unix())
	var output string
	output = string(rand.Intn(100))
	return output
}
