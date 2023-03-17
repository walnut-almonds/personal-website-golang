package telegram

import (
	"go.uber.org/dig"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"personal-website-golang/service/internal/thirdparty/logger"
)

type ITelegram interface {
	SendMessage(chatID int64, msg string) error
	SendImage(chatID int64, imagePath string) error
}

func NewTelegram(in digIn) ITelegram {
	bot, err := tgbotapi.NewBotAPI(in.AppConf.GetTelegramConfig().Token)
	if err != nil {
		panic(err)
	}

	return &telegram{
		digIn: in,

		bot: bot,
	}
}

type telegram struct {
	digIn

	bot *tgbotapi.BotAPI
}

type digIn struct {
	dig.In

	AppConf config.IAppConfig
	Logger  logger.ILogger `name:"appLogger"`
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
