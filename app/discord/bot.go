package discord

import (
	"app/config"

	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	session      *discordgo.Session
	message      *discordgo.Message
	messageTitle string
}

func NewBot() (*Bot, error) {
	conf, err := config.GetDiscordConfig()
	if err != nil {
		return nil, err
	}

	session, err := discordgo.New("Bot " + conf.BotID)
	if err != nil {
		return nil, err
	}

	return &Bot{
		session:      session,
		messageTitle: conf.MessageTitle,
	}, nil
}

func (b *Bot) Start() error {
	err := b.session.Open()
	if err != nil {
		return err
	}

	b.session.AddHandler(b.receiveMessage)
	b.session.AddHandler(b.receiveReaction)

	return nil
}

func (b *Bot) Stop() error {
	err := b.session.Close()
	if err != nil {
		return err
	}

	return nil
}
