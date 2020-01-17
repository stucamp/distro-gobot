package main

import (
	"strings"

	"github.com/mmcdole/gofeed"
)

type source struct {
	Name   string
	Source string
	URL    string
}

const urlsfilepath = "./json/urls.json"

// Creates a map of relavent links to various RSS feeds
func getRSSURLmap() map[string]string {
	var m map[string]string
	m = make(map[string]string)
	m["news"] = "https://distrowatch.com/news/dw.xml"
	m["torrent"] = "https://distrowatch.com/news/torrents.xml"
	m["security"] = "https://linuxsecurity.com/linuxsecurity_advisories.xml"
	m["release"] = "https://distrowatch.com/news/dwd.xml"
	return m
}

// Given a rss item, returns a formated string with distro's name and url
func strFormatOut(input *gofeed.Item) string {
	var output string = input.Title + "\n" + input.Link + "\n"
	return output
}

// Takes a URL to an RSS feed and returns a *gofeed.Feed struct
func parseUrlforStu(url string) *gofeed.Feed {
	fp := gofeed.NewParser()
	output, _ := fp.ParseURL(url)
	return output
}

// Takes a specified RSS xml file url and returns a formated string of the items.
func printReleases(m map[string]string) string {
	release := parseUrlforStu(m["release"])
	var output string
	for _, thing := range release.Items {
		output += strFormatOut(thing)
		output += "\n"
	}
	return output
}

// Takes a specified RSS xml file url and returns a formated string of the distro names and torrent urls.
func printTorrents(m map[string]string) string {
	release := parseUrlforStu(m["torrent"])
	var output string
	for _, thing := range release.Items {
		if isDesired(thing.Title) {
			output += strFormatOut(thing)
			output += "\n"
		}
	}
	return output
}

// Takes a specified RSS xml file url and returns a formated string of DistroWatch's New page
func printDistroWatchNews(m map[string]string) string {
	release := parseUrlforStu(m["news"])
	var output string
	for _, thing := range release.Items {
		if strings.Contains(thing.Title, "DistroWatch Weekly") {
			output += strFormatOut(thing)
			output += "\n"
		}
	}
	return output
}

// Takes a specified RSS xml file url and returns a formated string of the development releases
// and links to the resepctive page on distrowatch
func printDevReleaseNews(m map[string]string) string {
	release := parseUrlforStu(m["news"])
	var output string
	for _, thing := range release.Items {
		if strings.Contains(thing.Title, "Development") {
			output += strFormatOut(thing)
			output += "\n"
		}
	}
	return output
}

// Takes a specified RSS xml file url and returns a formated string of the distrobution releases
// and links to the resepctive page on distrowatch
func printDistReleaseNews(m map[string]string) string {
	release := parseUrlforStu(m["news"])
	var output string
	for _, thing := range release.Items {
		if strings.Contains(thing.Title, "Distribution") {
			output += strFormatOut(thing)
			output += "\n"
		}
	}
	return output
}

// Takes a specified RSS xml file url and returns a formated string of the recently posted security
// alerts found on Linux Security's RSS feed
func printSecurityNews(m map[string]string) string {
	release := parseUrlforStu(m["security"])
	var output string
	for _, thing := range release.Items {
		output += strFormatOut(thing)
		output += "\n"
	}
	return output
}
