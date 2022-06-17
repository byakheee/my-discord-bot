package handlers

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func OnMessageCreate(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.Bot {
		return
	}

	if _, err := session.ChannelMessageSend(message.ChannelID, "Hello! I'm bot!"); err != nil {
		log.Printf("Failed to send message to channel(%s). error: %s", message.ChannelID, err.Error())
	}
}
