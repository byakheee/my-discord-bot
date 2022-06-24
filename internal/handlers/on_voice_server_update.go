package handlers

import (
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"

	"github.com/byakheee/my-discord-bot/internal/extensions/logex"
)

func OnVoiceServerUpdate(discord *discordgo.Session, update *discordgo.VoiceServerUpdate) {
	log.Info().Msg("On voice server update!")
	logex.DebugRawJSON(update)
}
