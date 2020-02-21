package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func printJSON() {
	fmt.Println("Testies")
}

func openJSONfileAsByteArr(filepath string) []byte {
	file, err := ioutil.ReadFile(urlsfilepath) // For read access.
	if err != nil {
		fmt.Println("Failed to read file urls.json")
		fmt.Println(err)
		panic(err)
	}
	return file
}

func getJSONitems(data []byte) []interface{} {

	var items []interface{}

	err := json.Unmarshal(data, &items)
	if err != nil {
		fmt.Println("Failed to parse the json")
		fmt.Println(err)
		panic(err)
	}
	return items
}

// GetSourceURLFromJSON takes a filepath to JSON file and type of source as strings, parses the contents, returns true if
// result was found and and URL in string form for the source
func GetSourceURLFromJSON(path string, inSource string) (bool, string) {

	stuff := getJSONitems(openJSONfileAsByteArr(path))

	for k := range stuff {
		fmt.Println("Checking for sources...")
		if strings.ToLower(stuff[k].Name) == strings.ToLower(inSource) {
			fmt.Printf("Found %s source: %s\n", stuff[k].Source, stuff[k].Name)
			return true, stuff[k].(source).URL
		}
	}

	fmt.Printf("Can't find any sources for: %s\n", inSource)
	return false, ""
}

// GetWatchListFromJSON takes a filepath to JSON file, parses the contents, returns true if
// result was found and and array of strings with distro names that are watched
func GetWatchListFromJSON(path string) (bool, []string) {

	stuff := getJSONitems(openJSONfileAsByteArr(path))

	watchedNames := make([]string, 0)

	for k := range stuff {
		fmt.Println("Checking for watched distros...")
		if stuff[k].(watched).Watched {
			fmt.Printf("Found watched distro: %s\n", stuff[k].(watched).Name)
			watchedNames = append(watchedNames, stuff[k].Name)
		}
	}

	fmt.Printf("Could not find watched distros in: %s\n", path)
	return len(watchedNames) > 0, watchedNames
}

// GetQuotesListFromJSON takes a filepath to JSON file and author as strings, parses the contents, returns true if
// result was found and and array of strings with quotes
func GetQuotesListFromJSON(path string, auth string) (bool, []string) {

	stuff := getJSONitems(openJSONfileAsByteArr(path))

	for k := range stuff {
		if strings.ToLower(stuff[k].Author) == strings.ToLower(auth) {
			fmt.Printf("Found %d quotes for author: %s\n", len(stuff[k].Sayings), stuff[k].Author)
			if len(stuff[k].Sayings) > 0 {
				return true, stuff[k].Sayings
			}
			if len(stuff[k].(quote).Sayings) == 0 {
				fmt.Printf("%s exists, but has 0 sayings\n", auth)
				return false, make([]string, 0)
			}
		}
	}

	fmt.Printf("Could not find author: %s\n", auth)
	return false, make([]string, 0)
}

func closeFile(f *os.File) {
	fmt.Printf("Closing %s\n", f.Name())
	err := f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
