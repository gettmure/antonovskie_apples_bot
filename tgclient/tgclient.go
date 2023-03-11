package tgclient

import (
	"antonovskie_apples_bot/complimentr"
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	wg sync.WaitGroup
)

type TelegramBot interface {
	ListenUpdates(ctx context.Context)
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

func (bot *telegramBot) ListenUpdates(ctx context.Context) {
	ticker := time.NewTicker(5 * time.Second)

	complimentrClient := complimentr.InitClient()
	complimentr := complimentr.InitComplimentr(complimentrClient)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Graceful shutdown")

			return
		case <-ticker.C:
			response, err := bot.client.GetUpdates(bot.token, bot.lastUpdateId)
			if err != nil {
				fmt.Println(err)

				continue
			}

			for _, update := range response.Result {
				if bot.lastUpdateId == -1 {
					atomic.StoreInt64(&bot.lastUpdateId, update.UpdateId)

					continue
				}

				if bot.lastUpdateId == update.UpdateId {
					continue
				}

				wg.Add(1)
				atomic.StoreInt64(&bot.lastUpdateId, update.UpdateId)

				go func(update UpdateResponse) {
					defer wg.Done()

					handleUpdateResponse(bot, complimentr, update)
				}(update)
			}

			wg.Wait()
		}
	}
}

func (bot *telegramBot) SendMessage(message string, chatId int64) *MessageResponse {
	response, err := bot.client.SendMessage(bot.token, message, chatId)
	if err != nil {
		fmt.Println(err)

		return nil
	}

	return response
}

func (bot *telegramBot) SendAudio(audioId string, chatId int64, caption string) *MessageResponse {
	response, err := bot.client.SendAudio(bot.token, audioId, chatId, caption)
	if err != nil {
		fmt.Println(err)

		return nil
	}

	return response
}
