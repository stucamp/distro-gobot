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

func getStringFromJSON(inSource string) (bool, string) {
	file, err := ioutil.ReadFile(urlsfilepath) // For read access.
	if err != nil {
		fmt.Println("Failed to read file urls.json")
		fmt.Println(err)
		panic(err)
	}

	//defer closeFile(file)
	stuff := getJSONitems(file)

	for k := range stuff {
		if strings.ToLower(stuff[k].(source).Name) == strings.ToLower(inSource) {
			fmt.Printf("Found %s source: %s\n", stuff[k].(source).Source, stuff[k].(source).Name)
			return true, stuff[k].(source).URL
		}
	}

	fmt.Printf("Can't find any sources for: %s\n", inSource)
	return false, ""
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

func closeFile(f *os.File) {
	fmt.Printf("Closing %s\n", f.Name())
	err := f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
