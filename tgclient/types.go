package tgclient

type Response[T Fetchable] struct {
	Ok          bool    `json:"ok"`
	Description *string `json:"description"`
	Result      T       `json:"result"`
}

type Fetchable interface {
	GetMeResponse | int
}

type GetMeResponse struct {
	Id                      int     `json:"id"`
	Firstname               string  `json:"first_name"`
	Lastname                *string `json:"last_name"`
	Username                *string `json:"username"`
	LanguageCode            *string `json:"language_code"`
	IsPremium               *bool   `json:"is_premium"`
	AddedToAttachmentMenu   *bool   `json:"added_to_attachment_menu"`
	CanJoinGroups           *bool   `json:"can_join_groups"`
	CanReadAllGroupMessages *bool   `json:"can_read_all_group_messages"`
	SupportsInlineQueries   *bool   `json:"supports_inline_queries"`
}
