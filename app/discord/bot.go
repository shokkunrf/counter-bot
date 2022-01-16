package discord

import (
	"app/config"

	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	session *discordgo.Session
}

func MakeBot() (*Bot, error) {
	conf, err := config.GetDiscordConfig()
	if err != nil {
		return nil, err
	}

	session, err := discordgo.New("Bot " + conf.BotID)
	if err != nil {
		return nil, err
	}

	return &Bot{
		session: session,
	}, nil
}

func (b *Bot) Start() error {
	err := b.session.Open()
	if err != nil {
		return err
	}

	b.session.AddHandler(b.receiveMessage)

	err = b.session.UpdateGameStatus(1, "Watching No Channel")
	if err != nil {
		return err
	}

	return nil
}

func (b *Bot) Stop() error {
	err := b.session.Close()
	if err != nil {
		return err
	}

	return nil
}
