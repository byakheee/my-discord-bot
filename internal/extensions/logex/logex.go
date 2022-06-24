package logex

import (
	"encoding/json"

	"github.com/rs/zerolog/log"
)

func DebugAnyStruct(contents any) {
	c, err := json.Marshal(contents)
	if err != nil {
		log.Error().Err(err).Msg("Failed to marshal contents to json.")
	}

	log.Debug().RawJSON("contents", c).Msg("")
}
