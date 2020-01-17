package main

import (
	"math/rand"
	"strconv"
	"time"
)

// GetRandNumStr returns a random num from 0 - 100
func GetRandNumStr() string {
	rand.Seed(time.Now().Unix())
	var output string
	output = strconv.Itoa(rand.Intn(100))
	return output
}

// GetRandNumInt returns a random num from 0 - 100
func GetRandNumInt() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(100)
}
