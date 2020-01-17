package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

type watched struct {
	Name    string
	Watched bool
}

const watchedfilepath = "./watched.json"

func readWatchedJSON() (bool, []string) {
	file, err0 := ioutil.ReadFile(watchedfilepath) // For read access.
	if err0 != nil {
		fmt.Printf("Failed to read file %s", watchedfilepath)
		fmt.Println(err0)
		panic(err0)
	}

	var distros []watched

	err1 := json.Unmarshal(file, &distros)
	if err1 != nil {
		fmt.Printf("Failed to parse %s\n", watchedfilepath)
		fmt.Println(err1)
		panic(err1)
	}

	watchedNames := make([]string, 0)

	for k := range distros {
		fmt.Printf("Checking for watched distros...")
		if distros[k].Watched {
			watchedNames = append(watchedNames, distros[k].Name)
		}
	}
	return len(watchedNames) > 0, watchedNames
}

// Returns true if the distro is both listed in the JSON and has watched status
func isDesired(wanted string) bool {
	hasWatched, distros := readWatchedJSON()
	if hasWatched {
		for _, distro := range distros {
			if strings.Contains(wanted, distro) {
				return true
			}
		}
	}
	return false
}

// Returns a newline delimited string of distros currently being watched for easy printing to discord
func watchedDistros() string {
	hasWatched, distros := readWatchedJSON()
	var output string
	if hasWatched {
		for _, distro := range distros {
			output += distro + "\n"
		}
	}

	return output
}

//TODO: move to JSON and allow users to add/remove watched distros
