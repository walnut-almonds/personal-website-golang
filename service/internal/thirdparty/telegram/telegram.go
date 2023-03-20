package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type ITelegram interface {
	SendMessage(chatID int64, msg string) error
	SendImage(chatID int64, imagePath string) error
}

func NewTelegram(token string) ITelegram {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		panic(err)
	}

	return &telegram{
		bot: bot,
	}
}

type telegram struct {
	bot *tgbotapi.BotAPI
}

func (tg *telegram) SendMessage(chatID int64, msg string) error {
	m := tgbotapi.NewMessage(chatID, msg)
	_, err := tg.bot.Send(m)
	if err != nil {
		return err
	}
	return nil
}

func (tg *telegram) SendImage(chatID int64, imagePath string) error {
	msgP := tgbotapi.NewPhoto(chatID, tgbotapi.FilePath(imagePath))
	_, err := tg.bot.Send(msgP)
	if err != nil {
		return err
	}
	return nil
}
