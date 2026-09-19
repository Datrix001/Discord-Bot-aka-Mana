package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
// 	"os/signal"
// 	"syscall"

// 	"github.com/bwmarrin/discordgo"
// 	"github.com/joho/godotenv"
// )

// var commands = []*discordgo.ApplicationCommand{{
// 	Name:        "greet",
// 	Description: "Hello World"},
// }

// func main() {

// 	godotenv.Load()
// 	token := os.Getenv("BOT_TOKEN")

// 	ssn, err := discordgo.New("Bot " + token)

// 	if err != nil {
// 		log.Panic(err)
// 	}

// 	ssn.AddHandler(func(s *discordgo.Session, n *discordgo.InteractionCreate) {
// 		if n.ApplicationCommandData().Name == "greet" {
// 			err = s.InteractionRespond(n.Interaction, &discordgo.InteractionResponse{
// 				Type: discordgo.InteractionResponseChannelMessageWithSource,
// 				Data: &discordgo.InteractionResponseData{
// 					Content: "Hello World",
// 				},
// 			})

// 			if err != nil {
// 				log.Println(err)
// 			}
// 		}
// 	})
// 	// Connect to Discord
// 	err = ssn.Open()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer ssn.Close()

// 	fmt.Println("Bot is online!")

// 	// Register commands
// 	for _, command := range commands {

// 		_, err := ssn.ApplicationCommandCreate(
// 			ssn.State.User.ID,
// 			"", // guild ID
// 			command,
// 		)

// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		fmt.Println("Registered:", command.Name)
// 	}

// 	// Keep program running
// 	sc := make(chan os.Signal, 1)

// 	signal.Notify(
// 		sc,
// 		syscall.SIGINT,
// 		syscall.SIGTERM,
// 		os.Interrupt,
// 	)

// 	<-sc
// }
