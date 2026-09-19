package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
// 	"os/signal"
// 	"strings"
// 	"syscall"

// 	"github.com/bwmarrin/discordgo"
// 	"github.com/joho/godotenv"
// )

// const prefix = "!gobot"

// type Answers struct {
// 	OriginChannelId string
// 	FavFood         string
// 	FavGame         string
// }

// func (a *Answers) ToMessagesEmbeded() discordgo.MessageEmbed {
// 	field := []*discordgo.MessageEmbedField{{
// 		Name:  "Favourite Food",
// 		Value: a.FavFood,
// 	},
// 		{
// 			Name:  "Favourite Game",
// 			Value: a.FavGame,
// 		},
// 	}
// 	return discordgo.MessageEmbed{Title: "New Responses!", Fields: field}
// }

// var responses map[string]Answers = map[string]Answers{}

// func main() {
// 	godotenv.Load()
// 	token := os.Getenv("BOT_TOKEN")
// 	sess, err := discordgo.New("Bot " + token)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	sess.AddHandler(func(s *discordgo.Session, n *discordgo.MessageCreate) {

// 		// bot will ignore its own messages
// 		if n.Author.ID == s.State.User.ID {
// 			return
// 		}

// 		//dm logic
// 		if n.GuildID == "" {
// 			answers, ok := responses[n.ChannelID]
// 			if !ok {
// 				return
// 			}

// 			if answers.FavFood == "" {
// 				answers.FavFood = n.Content

// 				s.ChannelMessageSend(n.ChannelID, "What's ur fav game?")

// 				responses[n.ChannelID] = answers
// 				return
// 			} else {
// 				answers.FavGame = n.Content
// 				// log.Printf("Answers: %v,%v", answers.FavFood, answers.FavGame)
// 				embed := answers.ToMessagesEmbeded()
// 				s.ChannelMessageSendEmbed(answers.OriginChannelId, &embed)

// 				delete(responses, n.ChannelID)

// 			}
// 		}

// 		args := strings.Split(n.Content, " ")
// 		if args[0] != prefix {
// 			return
// 		}

// 		if args[1] == "Hello" {
// 			s.ChannelMessageSend(n.ChannelID, "world!")
// 		}
// 		if strings.Join(args[1:], " ") == "Mayank fuck you" {
// 			s.ChannelMessageDelete(n.ChannelID, n.Message.ID)
// 			s.ChannelMessageSend(n.ChannelID, "You have Been given warning not to send messgae like that to mayank.")
// 		}

// 		if args[1] == "prompt" {
// 			UserPromptHandler(s, n)
// 		}

// 	})

// 	sess.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

// 	err = sess.Open()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer sess.Close()

// 	fmt.Println("Bot is Online!!")

// 	sc := make(chan os.Signal, 1)
// 	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
// 	<-sc
// }

// func UserPromptHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
// 	channel, err := s.UserChannelCreate(m.Author.ID)
// 	if err != nil {
// 		log.Panic(err)
// 	}

// 	if _, ok := responses[channel.ID]; !ok {
// 		responses[channel.ID] = Answers{
// 			OriginChannelId: m.ChannelID,
// 			FavFood:         "",
// 			FavGame:         "",
// 		}
// 		s.ChannelMessageSend(channel.ID, "Hey there! Here are some questions")
// 		s.ChannelMessageSend(channel.ID, "What's ur favourite food?")

// 	} else {
// 		s.ChannelMessageSend(channel.ID, "We are still waiting...")
// 	}
// }
