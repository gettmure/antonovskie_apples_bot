package tgclient

import (
	"antonovskie_apples_bot/complimentr"
	"fmt"
	"log"
	"strings"
)

func handleStartMessage(bot *telegramBot, chatId int64) {
	message := "Привет, солнышко!🍎🍏❤️\n\n" +
		"Я - 🍎🍏❤️-бот (антоновские яблоки бот). Я написан для того, чтобы радовать тебя в любые моменты жизни! :)\n\n" +
		"Я - очень чуткий и милый бот, поэтому буду радовать тебя прекрасными словами. " +
		"Правда пока что я умею это делать только на английском языке, но я обязательно научусь и на русском!!!\n\n" +
		"Сейчас я хочу сказать про тебя вот что: %s!\n\n"

	compliment, err := complimentr.InitClient().GetCompliment()
	if err != nil {
		log.Println("complimentr failed for /start command", err)
		bot.SendMessage(fmt.Sprintf(message, "Блин-блинский, я сломался :("), chatId)
	}

	bot.SendMessage(fmt.Sprintf(message, compliment.Compliment), chatId)
}

func handleComplimentMessage(bot *telegramBot, chatId int64) {
	compliment, err := complimentr.InitClient().GetCompliment()
	if err != nil {
		log.Println(err)

		return
	}

	message := fmt.Sprintf("%s! 🍎🍏❤️", strings.Title(compliment.Compliment))
	bot.SendMessage(message, chatId)
}
