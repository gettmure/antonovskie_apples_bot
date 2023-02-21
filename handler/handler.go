package handler

import (
	"antonovskie_apples_bot/tgclient"
	"fmt"
)

func HandleGetMeResponse(ch chan any) {
	response := <-ch

	switch resp := response.(type) {
	case *tgclient.Response[tgclient.GetMeResponse]:
		fmt.Println("Bot info:")
		fmt.Println("- id:", resp.Result.Id)
		fmt.Println("- public name:", resp.Result.Firstname)
		fmt.Println("- username:", *resp.Result.Username)
	}
}

func HandleUpdateResponse(bot *tgclient.TelegramBot, ch chan any) {
	response := <-ch

	switch resp := response.(type) {
	case *tgclient.Response[tgclient.GetUpdatesResponse]:
		for _, update := range resp.Result {
			if bot.LastUpdateId == nil {
				bot.LastUpdateId = &update.UpdateId
			}

			if *bot.LastUpdateId == update.UpdateId {
				return
			} else {
				bot.LastUpdateId = &update.UpdateId
			}

			fmt.Println("Received update:")
			fmt.Println("- update id:", update.UpdateId)
			fmt.Println("- message id:", update.Message.MessageId)
			fmt.Println("- message text:", update.Message.Text)
		}
	}
}
