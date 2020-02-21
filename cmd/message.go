package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// This function will be called (due to AddHandler in main) every time a new
// message is created on any channel that the autenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	urlRSSMap := getRSSURLmap()

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.Contains(m.Content, "mint") && (m.Author.ID == "665638806732668960" || m.Author.ID == "401429986411675658") {
		f, err := os.Open("./media/source.gif")
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()
		s.ChannelFileSend(m.ChannelID, "facepalm.gif", f)
	}

	if strings.Contains(m.Content, "nomachine") && (m.Author.ID == "665638806732668960" || m.Author.ID == "401429986411675658") {
		s.ChannelMessageSend(m.ChannelID, "Real pros, only use CLI")
	}

	if m.Content == "PAVLOS!!!" && (m.Author.ID == "665638806732668960" || m.Author.ID == "401429986411675658") {
		s.ChannelMessageSend(m.ChannelID, "Stop tinkering...")
	}

	if strings.Contains(m.Content, "παύλος") && (m.Author.ID == "665638806732668960" || m.Author.ID == "401429986411675658") {
		s.ChannelMessageSend(m.ChannelID, "Stop tinkering...")
	}

	if m.Content == "!distrobot" {
		var output string
		output = "Commands are as follows:\n\n"
		output += "!distronews - Prints link to latest DistroWatch News\n"
		output += "!newreleases - Prints new releases\n"
		output += "!distreleases - Prints new distribution news\n"
		output += "!devreleases - Prints new devleopement news\n"
		output += "!security - Prints recent security news\n"
		output += "!isotorrent - Prints recently release torrent links\n"
		output += "!watched - Prints list of distros that will be returned by !isotorrent if available\n"

		s.ChannelMessageSend(m.ChannelID, output)
		fmt.Printf("Parsing %s\n", m.Content)
	}

	if strings.Contains(m.Content, "cookie") {
		s.MessageReactionAdd(m.ChannelID, m.ID, "🍪")
		fmt.Printf("Parsing %s\n", m.Content)
	}

	if m.Content == "!distronews" {
		output := printDistroWatchNews(urlRSSMap)
		s.ChannelMessageSend(m.ChannelID, output)
		fmt.Printf("Parsing %s\n", m.Content)
	}

	if m.Content == "!distreleases" {
		output := printDistReleaseNews(urlRSSMap)
		s.ChannelMessageSend(m.ChannelID, output)
		fmt.Printf("Parsing %s\n", m.Content)
	}

	if m.Content == "!devreleases" {
		output := printDevReleaseNews(urlRSSMap)
		s.ChannelMessageSend(m.ChannelID, output)
		fmt.Printf("Parsing %s\n", m.Content)
	}

	if m.Content == "!newreleases" {
		output := printReleases(urlRSSMap)
		s.ChannelMessageSend(m.ChannelID, output)
		fmt.Printf("Parsing %s\n", m.Content)
	}

	if m.Content == "!security" {
		output := printSecurityNews(urlRSSMap)
		s.ChannelMessageSend(m.ChannelID, output)
		fmt.Printf("Parsing %s\n", m.Content)
	}

	if m.Content == "!isotorrent" {
		output := printTorrents(urlRSSMap)
		s.ChannelMessageSend(m.ChannelID, output)
		fmt.Printf("Parsing %s\n", m.Content)
	}

	if m.Content == "!watched" {
		s.ChannelMessageSend(m.ChannelID, watchedDistros())
		fmt.Printf("Parsing %s\n", m.Content)
	}

	if strings.Contains(m.Content, "!quote") {
		fmt.Printf("Parsing %s\n", m.Content)
		auth := strings.Split(m.Content, " ")
		isThere, quote := GetRandQuote(auth[1])
		if isThere {
			s.ChannelMessageSend(m.ChannelID, quote)
		} else {
			neg := auth[1] + " has no quotes in the db"
			s.ChannelMessageSend(m.ChannelID, neg)
		}
	}

	if strings.Contains(m.Content, "!roll") {
		s.ChannelMessageSend(m.ChannelID, GetRandNumStr())
	}
}
