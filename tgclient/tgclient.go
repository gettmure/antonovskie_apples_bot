package tgclient

import (
	"context"
	"log"
	"sync"
	"time"
)

type TelegramBot interface {
	ListenUpdates(ctx context.Context, wg *sync.WaitGroup)
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

	for {
		select {
		case <-ctx.Done():
			log.Println("Graceful shutdown")

			return
		case <-ticker.C:
			updates, err := bot.client.GetUpdates(bot.token, -1)
			if err != nil {
				log.Println(err)
				continue
			}

			handleUpdateResponse(&bot.lastUpdateId, updates)
		}
	}
}
