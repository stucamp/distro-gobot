package main

import "fmt"

func main() {
	hasQuotes, quotes := GetQuotesListFromJSON(quotesfilepath, "confucius")
	fmt.Println("Has quotes: %b", hasQuotes)
	if hasQuotes {
		for _, quote := range quotes {
			fmt.Println(quote)
		}
	}

	hasURL, url := GetSourceURLFromJSON(urlsfilepath, "news")
	fmt.Println("Has URL: %b", hasURL)
	if hasURL {
		fmt.Println(url)

	}

	hasList, lists := GetWatchListFromJSON(watchedfilepath)
	fmt.Println("Has list: %b", hasList)
	if hasList {
		for _, list := range lists {
			fmt.Println(list)
		}
	}

}
