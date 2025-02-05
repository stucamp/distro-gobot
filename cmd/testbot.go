package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
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
	fmt.Println("Bot is now running...")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}
<<<<<<< HEAD

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	paul := "746974797698695258"
	stu := "401429986411675658"
	urlRSSMap := getRSSURLmap()

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.Contains(m.Content, "mint") && (m.Author.ID == stu || m.Author.ID == paul) {
		f, err := os.Open("./media/facepalm.gif")
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()
		s.ChannelFileSend(m.ChannelID, "facepalm.gif", f)
	}
	
	if strings.Contains(strings.ToLower(m.Content), "mint") {
		f, err := os.Open("./media/dataNO.gif")
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()
		s.ChannelFileSend(m.ChannelID, "dataNO.gif", f)
	}

	if strings.Contains(strings.ToLower(m.Content), "wait what?") && (m.Author.ID == stu || m.Author.ID == paul) {
		f, err := os.Open("./media/bubbles.gif")
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()
		s.ChannelFileSend(m.ChannelID, "bubbles.gif", f)
	}
	
	if strings.Contains(strings.ToLower(m.Content), "stu is awesome") {
		f, err := os.Open("./media/stu.gif")
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()
		s.ChannelFileSend(m.ChannelID, "stu.gif", f)
	}
	
	if strings.Contains(strings.ToLower(m.Content), "sucks") {
		f, err := os.Open("./media/wtf.gif")
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()
		s.ChannelFileSend(m.ChannelID, "wtf.gif", f)
	}
	
	if strings.Contains(m.Content, "ubuntu") && (m.Author.ID == stu || m.Author.ID == paul) {
		f, err := os.Open("./media/gud.gif")
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()
		s.ChannelFileSend(m.ChannelID, "gud.gif", f)
	}

	if strings.Contains(m.Content, "nomachine") && (m.Author.ID == paul || m.Author.ID == stu) {
		s.ChannelMessageSend(m.ChannelID, "Real pros, only use CLI")
	}

	if m.Content == "PAVLOS!!!" && (m.Author.ID == paul || m.Author.ID == stu) {
		s.ChannelMessageSend(m.ChannelID, "Stop tinkering...")
	}

	if strings.Contains(m.Content, "παύλος") && (m.Author.ID == paul || m.Author.ID == stu) {
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

	if strings.Contains(strings.ToLower(m.Content), "todd") {
		s.ChannelMessageSend(m.ChannelID, "Fix my car... for free!")
	}

	if strings.Contains(strings.ToLower(m.Content), "arch") {
		s.ChannelMessageSend(m.ChannelID, "My grandma can install arch... it's not hard.")
	}
	
	if strings.Contains(strings.ToLower(m.Content), "kali") {
		s.ChannelMessageSend(m.ChannelID, "Script Kiddie Detected")
	}
}
=======
>>>>>>> f3348f023988bc45ddbdf68c2f1139a3c1d2dc75
