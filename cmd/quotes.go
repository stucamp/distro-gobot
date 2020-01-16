package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

type quote struct {
	Author  string
	Sayings []string
}

const filepath = "./quotes.json"

func readJSON(auth string) []string {
	file, err0 := ioutil.ReadFile(filepath) // For read access.
	if err0 != nil {
		fmt.Println("Failed to read file quotes.json")
		panic(err0)
	}

	var quotes []quote

	err1 := json.Unmarshal(file, &quotes)
	if err1 != nil {
		fmt.Println("Failed to parse the json")
		panic(err1)
	}

	for k := range quotes {
		fmt.Printf("Comparing %s and %s...\n", strings.ToLower(quotes[k].Author), strings.ToLower(auth))
		if strings.ToLower(quotes[k].Author) == strings.ToLower(auth) {
			return quotes[k].Sayings
		}
	}

	return make([]string, 0)
}

func randQuote(quotes []string) string {
	rand.Seed(time.Now().Unix())
	return quotes[rand.Intn(len(quotes))]
}

// GetRandQuote takes an authore name as string and returns a randome quote stored
// in JSON if one such quote exists
func GetRandQuote(auth string) string {
	quotes := readJSON(auth)
	return (randQuote(quotes))
}

// AuthExists checks if author specified has quotes in the JSON
func AuthExists(auth string) bool {
	return (len(readJSON(auth)) <= 0)
}
