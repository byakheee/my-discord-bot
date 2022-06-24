package gateways

import (
	"github.com/byakheee/my-discord-bot/internal/extensions/logex"
)

// nolint:gochecknoglobals
var (
	stayingVoiceChannel = map[string]map[string]string{}
)

func SetStayngVoiceChannel(guildID, userID, channelID string) {
	if _, ok := stayingVoiceChannel[guildID]; !ok {
		stayingVoiceChannel[guildID] = map[string]string{}
	}

	stayingVoiceChannel[guildID][userID] = channelID
	logex.DebugAnyStruct(stayingVoiceChannel)
}

func DeleteStayngVoiceChannel(guildID, userID string) {
	if _, ok := stayingVoiceChannel[guildID]; ok {
		delete(stayingVoiceChannel[guildID], userID)
	}

	logex.DebugAnyStruct(stayingVoiceChannel)
}

func GetStayngVoiceChannel(guildID, userID string) string {
	logex.DebugAnyStruct(stayingVoiceChannel)

	return stayingVoiceChannel[guildID][userID]
}
