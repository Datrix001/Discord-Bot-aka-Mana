package main

import (
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

const prefix = "/bot"

func main() {

	godotenv.Load()
	token := os.Getenv("BOT_TOKEN")

	sessn, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Panic(err)
	}

	sessn.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {

		if m.Author.ID == s.State.User.ID {
			return
		}

		args := strings.Fields(m.Content)

		if len(args) < 2 || args[0] != prefix {
			return
		}

		args[1] = strings.ToLower(args[1])

		switch args[1] {
		case "hello":
			Greet(s, m)

		case "search":
			Search(s, m, args)

		case "dm":
			Dm(s, m)

		case "about":
			About(s, m)
		default:
			s.ChannelMessageSend(m.ChannelID, "Bro enter something that's reasonable, Don't just send anything. This is why she left you!!")
		}
	})

	err = sessn.Open()
	if err != nil {
		log.Panic(err)
	}

	defer sessn.Close()
	log.Println("Bot is running...")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

}
