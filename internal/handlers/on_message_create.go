package handlers

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func OnMessageCreate(discord *discordgo.Session, message *discordgo.MessageCreate) {
	log.Printf("On message create!\n%#v", *message)

	if message.Author.Bot {
		return
	}

	if message.Content == "" ||
		!(message.Type == discordgo.MessageTypeDefault ||
			message.Type == discordgo.MessageTypeReply) {
		// メッセージの種類: https://pkg.go.dev/github.com/bwmarrin/discordgo#MessageType
		log.Printf("Skip message type: %v, content: %s", message.Type, message.Content)

		return
	}

	if _, err := discord.ChannelMessageSend(message.ChannelID, "Hello! I'm bot!"); err != nil {
		log.Printf("Failed to send message to channel(%s). error: %s\n", message.ChannelID, err.Error())
	}
}
