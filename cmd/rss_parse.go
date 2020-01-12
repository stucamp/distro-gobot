package main

import (
	"fmt"
	"strings"

	"github.com/mmcdole/gofeed"
)

func main() {

	//_, month, day := time.Now().Date()
	//fmt.Printf("%02d/%02d\n", month, day)

	print_dist_release_news(get_URL_map())
}

func print_title_link(input *gofeed.Item) {
	fmt.Println(input.Title)
	fmt.Println(input.Link)
	fmt.Print("\n")
}

func parse_url(url string) *gofeed.Feed {
	fp := gofeed.NewParser()
	output, _ := fp.ParseURL(url)
	return output
}

func get_URL_map() map[string]string {
	var m map[string]string
	m = make(map[string]string)
	m["news"] = "https://distrowatch.com/news/dw.xml"
	m["torrent"] = "https://distrowatch.com/news/torrents.xml"
	m["security"] = "https://linuxsecurity.com/linuxsecurity_advisories.xml"
	m["release"] = "https://distrowatch.com/news/dwd.xml"
	return m
}

func print_dev_release_news(m map[string]string) {
	release := parse_url(m["news"])
	for _, thing := range release.Items {
		if !strings.Contains(thing.Title, "Distribution") {
			print_title_link(thing)
		}
	}
}

func print_dist_release_news(m map[string]string) {
	release := parse_url(m["news"])
	for _, thing := range release.Items {
		if strings.Contains(thing.Title, "Distribution") {
			print_title_link(thing)

		}
	}
}
