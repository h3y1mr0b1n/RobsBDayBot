package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/jasonlvhit/gocron"
)

var (
	Token string
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()

}

func task() {
	discord, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("[CRITICAL] Session unable to start", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	discord.AddHandler(messageCreate)

	// In this example, we only care about receiving message events. (have no friends tho)
	discord.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)

	now := time.Now().Format("Jan 2, 2006")

	if now == "Jan 14, 2021" {
		discord.ChannelMessageSend("788552340428554240", "Ich habe heute Geburtstag :O ")
		discord.Close()
	} else {
		discord.ChannelMessageSend("788552340428554240", "Keiner hat heute Geburtstag")
		discord.Close()
	}

	//Open Websocket
	print(Token)
	err = discord.Open()
	if err != nil {
		fmt.Println("[CRITICAL] Session unable to start", err)
		return
	}

	discord.Close()
}

func main() {
	gocron.Every(1).Day().At("16:15").Do(task)
	fmt.Println("Task scheduled")
	_, time := gocron.NextRun()
	fmt.Println(time)
	<-gocron.Start()

}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	//Make Bot ignore itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "!dateadd" {
		now := time.Now()
		s.ChannelMessageSend(m.ChannelID, now.Format("Jan 2, 2006"))
	}
}
