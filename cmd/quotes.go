package main

import (
	"fmt"
	"math/rand"
	"time"
)

type quote struct {
	Author  string
	Sayings []string
}

const quotesfilepath = "./json/quotes.json"

func randQuote(quotes []string) string {

	randseed := time.Now().Unix()
	fmt.Printf("Random seed: %d\n", randseed)
	rand.Seed(randseed)
	quote := quotes[rand.Intn(len(quotes))]
	fmt.Printf("Select quote: %s\n", quote)
	return quote
}

// GetRandQuote takes an authore name as string and returns a randome quote stored
// in JSON if one such quote exists
func GetRandQuote(auth string) (bool, string) {

	hasQuotes, quotes := GetQuotesListFromJSON(quotesfilepath, auth)

	var randomQ string
	if hasQuotes {
		randomQ = randQuote(quotes)
		fmt.Printf("%s\n Selected as quote", randomQ)
	} else {
		fmt.Println("Didn't find anything in quote array")
		randomQ = ""
	}
	return hasQuotes, randomQ
}
