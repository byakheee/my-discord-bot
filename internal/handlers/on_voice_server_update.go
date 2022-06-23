package handlers

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func OnVoiceServerUpdate(discord *discordgo.Session, update *discordgo.VoiceServerUpdate) {
	log.Printf("On voice server update!\n%#v", *update)
}
