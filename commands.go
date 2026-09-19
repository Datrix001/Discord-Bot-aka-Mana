package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var comicName = []string{"solo leveling", "nano machine"}

func Normalize(s string) string {
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, "_", " ")
	s = strings.ReplaceAll(s, "-", " ")
	return s
}

func Greet(s *discordgo.Session, m *discordgo.MessageCreate) {
	usr := m.Author.GlobalName
	s.ChannelMessageSend(m.ChannelID, "Hello "+usr)
	s.MessageReactionAdd(m.ChannelID, m.ID, "👋")
}

//PART ONE SEARCH

// func Search(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
// 	arg := strings.Join(args[2:], " ")
// 	arg = Normalize(arg)
// 	isThere := false
// 	var comic string

// 	for _, value := range comicName {
// 		if strings.Contains(arg, value) {
// 			isThere = true
// 			comic = value
// 		}
// 	}
// 	if isThere {
// 		s.ChannelMessageSend(m.ChannelID, "Searching for "+comic)
// 	} else {
// 		s.ChannelMessageSend(m.ChannelID, "No Comic Like that exist")

// 	}
// }

// PART 2 SEARCH
// SEARCH MANHWA

func Search(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	// isThere := false
	arg := strings.Join(args[2:], " ")
	arg = Normalize(arg)
	comicDetails, err := GetManga(arg)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Bro The manhwa you mentioned doesn't exist")
		return
	}
	details := fmt.Sprintf(" Author for %s is %s.\n Total chapters are %s", comicDetails.Name, comicDetails.Author, comicDetails.LatestChapter)
	messageEmbed := discordgo.MessageEmbed{
		Title:       strings.ToUpper(comicDetails.Name),
		URL:         comicDetails.CoverURL,
		Description: details,
		Color:       0x5865F2,
	}
	s.ChannelMessageSendEmbed(m.ChannelID, &messageEmbed)

}

func Dm(s *discordgo.Session, m *discordgo.MessageCreate) {
	userName := m.Author.GlobalName
	userChannerId := m.Author.ID
	dmChannel, err := s.UserChannelCreate(userChannerId)
	if err != nil {
		log.Panic(err)
	}
	s.ChannelMessageSend(dmChannel.ID, "Hello "+userName)
}

func About(s *discordgo.Session, m *discordgo.MessageCreate) {
	userName := m.Author.GlobalName
	isPremium := ""
	if m.Author.PremiumType != 0 {
		isPremium = "You have premium account!"
	} else {
		isPremium = "You don't have a premium account!!"
	}
	message := fmt.Sprintf("Your Name is %s.\n %v", userName, isPremium)
	nameField := &discordgo.MessageEmbedField{
		Name:   "Name",
		Value:  userName,
		Inline: true,
	}

	premiumField := &discordgo.MessageEmbedField{
		Name:   "Premium",
		Value:  isPremium,
		Inline: true,
	}

	chapterField := &discordgo.MessageEmbedField{
		Name:   "Chapter",
		Value:  "220",
		Inline: true,
	}
	embedMessage := discordgo.MessageEmbed{
		Title:       "User Information!",
		Description: message,
		Color:       0x5865F2,
		Fields: []*discordgo.MessageEmbedField{
			nameField,
			premiumField,
			chapterField,
		},
	}

	s.ChannelMessageSendEmbed(m.ChannelID, &embedMessage)
	// s.ChannelMessageSendEmbed(m.ChannelID, *discordgo.Em)

}
