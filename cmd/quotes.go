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

const quotesfilepath = "./quotes.json"

func readQuotesJSON(auth string) []string {
	file, err0 := ioutil.ReadFile(quotesfilepath) // For read access.
	if err0 != nil {
		fmt.Println("Failed to read file quotes.json")
		fmt.Println(err0)
		panic(err0)
	}

	var quotes []quote

	err1 := json.Unmarshal(file, &quotes)
	if err1 != nil {
		fmt.Println("Failed to parse the json")
		fmt.Println(err1)
		panic(err1)
	}

	for k := range quotes {
		fmt.Printf("Comparing %s and %s...\n", strings.ToLower(quotes[k].Author), strings.ToLower(auth))
		if strings.ToLower(quotes[k].Author) == strings.ToLower(auth) {
			fmt.Println("Match Success")
			arrquotes := quotes[k].Sayings
			fmt.Printf("Found these %d quotes:\n", len(arrquotes))
			for _, quote := range arrquotes {
				fmt.Println(quote)
			}
			return arrquotes
		}
	}

	fmt.Println("Match Failed")
	return make([]string, 0)
}

func randQuote(quotes []string) string {
	randseed := time.Now().Unix()
	fmt.Printf("%d used as random seed\n", randseed)
	rand.Seed(randseed)
	quote := quotes[rand.Intn(len(quotes))]
	fmt.Printf("\"%s\" used as output\n", quote)
	return quote
}

// GetRandQuote takes an authore name as string and returns a randome quote stored
// in JSON if one such quote exists
func GetRandQuote(auth string) (bool, string) {
	quotes := readQuotesJSON(auth)
	var randomQ string
	isThere := false
	fmt.Printf("The quote array is %d long", len(quotes))
	if len(quotes) > 0 {
		fmt.Println("Printing array of quotes")
		for _, quote := range quotes {
			fmt.Println(quote)
		}
		randomQ = randQuote(quotes)
		isThere = true
		fmt.Printf("%s\n Selected as quote", randomQ)
	} else {
		fmt.Println("Didn't find anything in quote array")
		randomQ = ""
	}
	return isThere, randomQ
}
