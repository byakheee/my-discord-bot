package handlers

import (
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"

	"github.com/byakheee/my-discord-bot/internal/extensions/logex"
	"github.com/byakheee/my-discord-bot/internal/gateways"
)

func OnVoiceStateUpdate(discord *discordgo.Session, stateUpdate *discordgo.VoiceStateUpdate) {
	log.Info().Msg("On voice state update!")
	logex.DebugAnyStruct(stateUpdate)

	if stateUpdate.ChannelID != "" {
		gateways.SetStayngVoiceChannel(stateUpdate.GuildID, stateUpdate.UserID, stateUpdate.ChannelID)
	} else {
		gateways.DeleteStayngVoiceChannel(stateUpdate.GuildID, stateUpdate.UserID)
	}
}
