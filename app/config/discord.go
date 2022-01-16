package config

import (
	"errors"
	"os"
)

type DiscordConfig struct {
	BotID string
}

func GetDiscordConfig() (*DiscordConfig, error) {
	botID := os.Getenv("DISCORD_BOT_ID")
	if botID == "" {
		return nil, errors.New("[Config] DISCORD_BOT_ID is empty")
	}

	return &DiscordConfig{
		BotID: botID,
	}, nil
}
