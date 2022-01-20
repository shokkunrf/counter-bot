package config

import (
	"errors"
	"os"
)

type DiscordConfig struct {
	BotID        string
	MessageTitle string
}

const BOT_MESSAGE_TITLE = "Counter"

func GetDiscordConfig() (*DiscordConfig, error) {
	botID := os.Getenv("DISCORD_BOT_ID")
	if botID == "" {
		return nil, errors.New("[Config] DISCORD_BOT_ID is empty")
	}

	messageTitle := os.Getenv("DISCORD_BOT_MESSAGE_TITLE")
	if messageTitle == "" {
		messageTitle = BOT_MESSAGE_TITLE
	}

	return &DiscordConfig{
		BotID:        botID,
		MessageTitle: messageTitle,
	}, nil
}
