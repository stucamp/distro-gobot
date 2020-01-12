package main

import (
	"strings"

	"github.com/mmcdole/gofeed"
)

func getURLmap() map[string]string {
	var m map[string]string
	m = make(map[string]string)
	m["news"] = "https://distrowatch.com/news/dw.xml"
	m["torrent"] = "https://distrowatch.com/news/torrents.xml"
	m["security"] = "https://linuxsecurity.com/linuxsecurity_advisories.xml"
	m["release"] = "https://distrowatch.com/news/dwd.xml"
	return m
}

func strFormatOut(input *gofeed.Item) string {
	var output string = input.Title + "\n" + input.Link + "\n"
	return output
}

func parseUrlforStu(url string) *gofeed.Feed {
	fp := gofeed.NewParser()
	output, _ := fp.ParseURL(url)
	return output
}

func printReleases(m map[string]string) string {
	release := parseUrlforStu(m["release"])
	var output string
	for _, thing := range release.Items {
		output += strFormatOut(thing)
		output += "\n"
	}
	return output
}

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

func printSecurityNews(m map[string]string) string {
	release := parseUrlforStu(m["security"])
	var output string
	for _, thing := range release.Items {
		output += strFormatOut(thing)
		output += "\n"
	}
	return output
}

func isDesired(wanted string) bool {
	distros := []string{"debian", "neon", "raspbian", "ubuntu", "antergos", "manjaro", "CentOS", "Fedora", "kali", "tails", "popos",
		"ipfire", "elementaryos", "pfsense", "openmediavault", "FreeNAS", "FreeBSD", "Peppermint", "mint",
		"openSUSE", "Zorin", "proxmox", "gparted-live", "systemrescuedc", "OSMC", "FreeBSD", "untangle"}
	for _, distro := range distros {
		if strings.Contains(wanted, distro) {
			return true
		}
	}
	return false
}

func watchedDistros() string {
	distros := []string{"debian", "neon", "raspbian", "ubuntu", "antergos", "manjaro", "CentOS", "Fedora", "kali", "tails", "popos",
		"ipfire", "elementaryos", "pfsense", "openmediavault", "FreeNAS", "FreeBSD", "Peppermint", "mint",
		"openSUSE", "Zorin", "proxmox", "gparted-live", "systemrescuedc", "OSMC", "FreeBSD", "untangle"}

	var output string
	for _, distro := range distros {
		output += distro + "\n"
	}
	return output
}
