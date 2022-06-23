package handlers

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func OnVoiceStateUpdate(discord *discordgo.Session, stateUpdate *discordgo.VoiceStateUpdate) {
	log.Printf("On voice state update!\n%#v", *stateUpdate)
}
