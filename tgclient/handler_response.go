package tgclient

import (
	"antonovskie_apples_bot/complimentr"
	"sync/atomic"
)

func handleGetMeResponse(response *GetMeResponse) {
	logGetMe(response)
}

func handleUpdateResponse(bot *telegramBot, complimentr complimentr.Complimentr, update UpdateResponse) {
	if bot.lastUpdateId == update.UpdateId {
		return
	}

	if bot.lastUpdateId == -1 {
		atomic.StoreInt64(&bot.lastUpdateId, update.UpdateId)

		return
	}

	atomic.StoreInt64(&bot.lastUpdateId, update.UpdateId)
	logUpdate(&update)

	switch update.Message.Text {
	case "/start":
		handleStartCommand(bot, complimentr, update.Message.Chat.Id)
	case "/story":
		handleStoryCommand(bot, update.Message.Chat.Id)
	case "/compliment":
		handleComplimentCommand(bot, complimentr, update.Message.Chat.Id)
	}
}
