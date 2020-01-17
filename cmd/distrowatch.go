package main

import (
	"strings"
)

type watched struct {
	Name    string
	Watched bool
}

const watchedfilepath = "./json/watched.json"

// Returns true if the distro is both listed in the JSON and has watched status
func isDesired(wanted string) bool {
	hasWatched, distros := GetWatchListFromJSON(watchedfilepath)
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
	hasWatched, distros := GetWatchListFromJSON(watchedfilepath)
	var output string
	if hasWatched {
		for _, distro := range distros {
			output += distro + "\n"
		}
	}
	return output
}

//TODO: allow users to add/remove watched distros
