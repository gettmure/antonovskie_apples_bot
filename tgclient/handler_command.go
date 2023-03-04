package tgclient

import (
	"antonovskie_apples_bot/complimentr"
	"fmt"
	"log"
	"strings"
)

const errorMessage = "Блин-блинский, я сломался :("

func handleStartCommand(bot *telegramBot, complimentr complimentr.Complimentr, chatId int64) {
	const message = "Привет, солнышко!🍎🍏❤️\n\n" +
		"Я - 🍎🍏❤️-бот (антоновские яблоки бот). Я написан для того, чтобы радовать тебя в любые моменты жизни! :)\n\n" +
		"Я - очень чуткий и милый бот, поэтому буду радовать тебя прекрасными словами. " +
		"Правда пока что я умею это делать только на английском языке, но я обязательно научусь и на русском!!!\n\n" +
		"Сейчас я хочу сказать про тебя вот что: %s!\n\n"

	compliment, err := complimentr.GetCompliment()
	if err != nil {
		log.Println("complimentr failed for /start command", err)
		bot.SendMessage(fmt.Sprintf(message, errorMessage), chatId)

		return
	}

	bot.SendMessage(fmt.Sprintf(message, *compliment), chatId)
}

func handleComplimentCommand(bot *telegramBot, complimentr complimentr.Complimentr, chatId int64) {
	compliment, err := complimentr.GetCompliment()
	if err != nil {
		log.Println("complimentr failed for /compliment command", err)
		bot.SendMessage(errorMessage, chatId)

		return
	}

	message := fmt.Sprintf("%s! 🍎🍏❤️", strings.Title(*compliment))
	bot.SendMessage(message, chatId)
}

func handleStoryCommand(bot *telegramBot, chatId int64) {
	const message = "Родион был талантливым разработчиком и давно искал свой проект, который был бы полезен обществу. Однажды он услышал о антоновских яблоках и решил, что это то, чем он бы занялся.\n\n" +
		"Родион начал разрабатывать специальное приложение, которое помогало бы выращивать антоновские яблоки. Сначала он изучил все информацию об этом виде яблок и выяснил, что они требуют особого ухода. Затем он разработал приложение, которое позволяло бы отслеживать все этапы выращивания яблок, а также давало бы полезные советы о том, как их выращивать.\n\n" +
		"Когда приложение было готово, Родион решил его попробовать. Он купил несколько саженцев антоновских яблок и приступил к выращиванию. Приложение очень помогло ему, и уже через несколько месяцев он смог собрать первый урожай.\n\n" +
		"Яблоки получились замечательные, и Родион решил поделиться своим опытом с другими людьми. Он разместил свое приложение в интернете и начал активно рекламировать его. Скоро о его приложении услышали многие люди, и они начали использовать его для выращивания антоновских яблок.\n\n" +
		"Вскоре Родион стал популярным разработчиком, и его приложение стало незаменимым инструментом для садоводов. Люди по всей стране начали выращивать антоновские яблоки, и они получили огромную популярность.\n\n" +
		"Родион был очень рад, что его проект оказался таким успешным. Он понимал, что многие люди благодарны ему за способность выращивать вкусные и здоровые яблоки своими руками. Он продолжал работать над своим приложением и внедрять в него новые функции, чтобы оно было еще более полезным для людей.\n\n" +
		"Таким образом, Родион не только нашел свой идеальный проект, но и принес множество пользы обществу, помогая людям выращивать вкусные и здоровые антоновские яблоки."

	bot.SendMessage(message, chatId)
}

func handleTrackCommand(bot *telegramBot, chatId int64) {
	const audioId = "CQACAgIAAxkBAAEd9uVkA_xYxa_yw-8UeWzO-6-ir-cNEQACECMAAuFm0UshDfJ0TvfpkC4E"

	const caption = "Этот трек был написан специально для тебя🍎🍏❤️\n\n" +
		"Я давно хотел сделать что-то осознанное и без накуренного кринжа. Ты стала катализатором того, что этот трек появился. " +
		"В этом тексте много отсылок на тебя, например:\n\n" +
		"- \"хочу наслаждаться только тобой\" - отсылка на твои обои на телефоне\n" +
		"- \"ошибался не раз, всё для того, чтоб ты стала моя\" - отсылка на проёбы с моей стороны"

	bot.SendAudio(audioId, chatId, caption)
}