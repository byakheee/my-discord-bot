package handlers

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func OnMessageCreate(discord *discordgo.Session, message *discordgo.MessageCreate) {
	log.Printf("guildId: %s", message.GuildID)
	log.Printf("channelId: %s", message.ChannelID)
	log.Printf("userId: %s", message.Author.ID)

	connections, err := discord.UserConnections()
	if err != nil {
		log.Printf("Failed to get user connections. error: %s\n", err.Error())
	}

	log.Printf("user connects num: %d", len(connections))

	for i, c := range connections {
		log.Printf("user connectin #%d: %s", i, c.Name)
	}

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
