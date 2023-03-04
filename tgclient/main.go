package tgclient

import (
	"antonovskie_apples_bot/complimentr"
	"context"
	"log"
	"sync"
	"time"
)

type TelegramBot interface {
	ListenUpdates(ctx context.Context, wg *sync.WaitGroup)
	SendMessage(message string, chatId int64) *MessageResponse
	SendAudio(audioId string, chatId int64, caption string) *MessageResponse
}

type telegramBot struct {
	token        string // https://core.telegram.org/bots/api#authorizing-your-bot
	lastUpdateId int64
	client       ApiClient
}

func InitBot(token string, apiClient ApiClient) (TelegramBot, error) {
	bot := &telegramBot{token: token, lastUpdateId: -1, client: apiClient}

	_, err := bot.client.GetMe(bot.token)
	if err != nil {
		return nil, err
	}

	return bot, nil
}

func (bot *telegramBot) ListenUpdates(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	ticker := time.NewTicker(5 * time.Second)

	complimentrClient := complimentr.InitClient()
	complimentr := complimentr.InitComplimentr(complimentrClient)

	for {
		select {
		case <-ctx.Done():
			log.Println("Graceful shutdown")

			return
		case <-ticker.C:
			response, err := bot.client.GetUpdates(bot.token, bot.lastUpdateId)
			if err != nil {
				log.Println(err)
				continue
			}

			for _, update := range response.Result {
				handleUpdateResponse(bot, complimentr, update)
			}
		}
	}
}

func (bot *telegramBot) SendMessage(message string, chatId int64) *MessageResponse {
	response, err := bot.client.SendMessage(bot.token, message, chatId)
	if err != nil {
		log.Println(err)

		return nil
	}

	return response
}

func (bot *telegramBot) SendAudio(audioId string, chatId int64, caption string) *MessageResponse {
	response, err := bot.client.SendAudio(bot.token, audioId, chatId, caption)
	if err != nil {
		log.Println(err)

		return nil
	}

	return response
}
