package discord

import (
	"app/store"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

func (b *Bot) generateMessageEmbedFields(session *discordgo.Session, counters []store.Counter) ([]*discordgo.MessageEmbedField, error) {
	messageField := []*discordgo.MessageEmbedField{}

	for _, counter := range counters {
		user, err := session.User(counter.UserID)
		if err != nil {
			return nil, err
		}

		messageField = append(messageField, &discordgo.MessageEmbedField{
			Name:   user.Username,
			Value:  strconv.Itoa(counter.Count) + " pt",
			Inline: true,
		})
	}

	return messageField, nil
}
