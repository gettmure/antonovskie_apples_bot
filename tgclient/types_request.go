package tgclient

type SendMessageData struct {
	ChatId int64  `json:"chat_id"`
	Text   string `json:"text"`
}

type SendAudioData struct {
	ChatId  int64  `json:"chat_id"`
	AudioId string `json:"audio"`
	Caption string `json:"caption"`
}
