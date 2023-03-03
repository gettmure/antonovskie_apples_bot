package tgclient

import (
	"sync/atomic"
)

func handleGetMeResponse(response *GetMeResponse) {
	logGetMe(response)
}

func handleUpdateResponse(bot *telegramBot, update UpdateResponse) {
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
		handleStartMessage(bot, update.Message.Chat.Id)
	case "/story":
		handleStoryMessage(bot, update.Message.Chat.Id)
	}
}
