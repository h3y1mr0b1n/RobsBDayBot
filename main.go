package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var (
	Token string
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()

}

func main() {
	discord, err := discordgo.New("Bot " + "authentication token")
	if err != nil {
		fmt.Println("[CRITICAL] Session unable to start", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	discord.AddHandler(messageCreate)

	// In this example, we only care about receiving message events. (have no friends tho)
	discord.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)

	//Open Websocket
	if err != nil {
		fmt.Println("[CRITICAL] Session unable to start", err)
		return
	}

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	discord.Close()

}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	//Make Bot ignore itself (that moody Bitch)
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "!dateadd" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}
}
