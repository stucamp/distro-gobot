package main

import (
	"math/rand"
	"strconv"
	"time"
)

// GetRandNum returns a random num from 0 - 100
func GetRandNum() string {
	rand.Seed(time.Now().Unix())
	var output string
	output = strconv.Itoa(rand.Intn(100))
	return output
}
