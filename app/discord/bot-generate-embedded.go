package discord

import (
	"app/store"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

func (b *Bot) generateEmbeddedMessage(counters []store.Counter) (discordgo.MessageEmbed, error) {
	messageField := []*discordgo.MessageEmbedField{}

	for _, counter := range counters {
		user, err := b.session.User(counter.UserID)
		if err != nil {
			return discordgo.MessageEmbed{}, err
		}

		messageField = append(messageField, &discordgo.MessageEmbedField{
			Name:   user.Username,
			Value:  strconv.Itoa(counter.Count) + " pt",
			Inline: true,
		})
	}

	return discordgo.MessageEmbed{
		Title:       "Counter",
		Description: "Increment :arrow_up:, Decrement :arrow_down:, Reset :zero:",
		Fields:      messageField,
	}, nil
}
