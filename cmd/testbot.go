package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

// Variables used for command line parameters
var (
	Token string
)

func init() {

	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.Contains(m.Content, "mint") && (m.Author.ID == "665638806732668960" || m.Author.ID == "401429986411675658") {
		s.ChannelMessageSend(m.ChannelID, "https://media2.giphy.com/media/XsUtdIeJ0MWMo/source.gif")
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
		output += "!distronews - Prints release to latest DistroWatch News\n"
		output += "!newreleases - Prints new releases\n"
		output += "!distreleases - Prints new distribution news\n"
		output += "!devreleases - Prints new devleopement news\n"
		output += "!security - Prints recent security news\n"
		output += "!isotorrent - Prints recently release torrent links\n"
		output += "!watched - Prints list of distros that will be returned by !isotorrent if available\n"

		s.ChannelMessageSend(m.ChannelID, output)
	}

	if m.Content == "!distronews" {
		output := printDistroWatchNews(getURLmap())
		s.ChannelMessageSend(m.ChannelID, output)
	}

	if m.Content == "!distreleases" {
		output := printDistReleaseNews(getURLmap())
		s.ChannelMessageSend(m.ChannelID, output)
	}

	if m.Content == "!devreleases" {
		output := printDevReleaseNews(getURLmap())
		s.ChannelMessageSend(m.ChannelID, output)
	}

	if m.Content == "!newreleases" {
		output := printReleases(getURLmap())
		s.ChannelMessageSend(m.ChannelID, output)
	}

	if m.Content == "!security" {
		output := printSecurityNews(getURLmap())
		s.ChannelMessageSend(m.ChannelID, output)
	}

	if m.Content == "!isotorrent" {
		output := printTorrents(getURLmap())
		s.ChannelMessageSend(m.ChannelID, output)
	}

	if m.Content == "!watched" {
		s.ChannelMessageSend(m.ChannelID, watchedDistros())
	}

}
