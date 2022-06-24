package handlers

import (
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"

	"github.com/byakheee/my-discord-bot/internal/extensions/logex"
)

func OnVoiceStateUpdate(discord *discordgo.Session, stateUpdate *discordgo.VoiceStateUpdate) {
	log.Info().Msg("On voice state update!")
	logex.DebugRawJSON(stateUpdate)
}
