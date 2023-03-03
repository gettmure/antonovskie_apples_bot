package tgclient

type GetMeResponse struct {
	Ok          bool    `json:"ok"`
	Description *string `json:"description"`
	Result      struct {
		Id                      int64  `json:"id"`
		Firstname               string `json:"first_name"`
		Lastname                string `json:"last_name"`
		Username                string `json:"username"`
		LanguageCode            string `json:"language_code"`
		IsPremium               *bool  `json:"is_premium"`
		AddedToAttachmentMenu   *bool  `json:"added_to_attachment_menu"`
		CanJoinGroups           *bool  `json:"can_join_groups"`
		CanReadAllGroupMessages *bool  `json:"can_read_all_group_messages"`
		SupportsInlineQueries   *bool  `json:"supports_inline_queries"`
	} `json:"result"`
}

type GetUpdatesResponse struct {
	Ok          bool             `json:"ok"`
	Description *string          `json:"description"`
	Result      []UpdateResponse `json:"result"`
}

type UpdateResponse struct {
	UpdateId int64           `json:"update_id"`
	Message  MessageResponse `json:"message"`
}

type MessageResponse struct {
	MessageId int64        `json:"message_id"`
	Text      string       `json:"text"`
	Chat      ChatResponse `json:"chat"`
}

type ChatResponse struct {
	Id int64 `json:"id"`
}
