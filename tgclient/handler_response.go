package tgclient

import (
	"antonovskie_apples_bot/complimentr"
)

func handleGetMeResponse(response *GetMeResponse) {
	logGetMe(response)
}

func handleUpdateResponse(bot *telegramBot, complimentr complimentr.Complimentr, update UpdateResponse) {
	logUpdate(&update)

	switch update.Message.Text {
	case "/start":
		handleStartCommand(bot, complimentr, update.Message.Chat.Id)
	case "/story":
		handleStoryCommand(bot, update.Message.Chat.Id)
	case "/compliment":
		handleComplimentCommand(bot, complimentr, update.Message.Chat.Id)
	case "/track":
		handleTrackCommand(bot, update.Message.Chat.Id)
	}
}
