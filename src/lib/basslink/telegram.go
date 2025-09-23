package basslink

import (
	"fmt"
	"github.com/vannleonheart/goutil"
)

type TelegramBotConfig struct {
	ApiUrl string `json:"api_url"`
	Token  string `json:"token"`
}

type TelegramBotClient struct {
	Config TelegramBotConfig `json:"config"`
}

type TelegramBotResponse struct {
	Ok          bool    `json:"ok"`
	ErrorCode   *int    `json:"error_code,omitempty"`
	Description *string `json:"description,omitempty"`
}

type TelegramBotUser struct {
	Id                      int     `json:"id"`
	IsBot                   bool    `json:"is_bot"`
	FirstName               *string `json:"first_name,omitempty"`
	LastName                *string `json:"last_name,omitempty"`
	Username                *string `json:"username,omitempty"`
	LanguageCode            *string `json:"language_code,omitempty"`
	CanJoinGroups           *bool   `json:"can_join_groups,omitempty"`
	CanReadAllGroupMessages *bool   `json:"can_read_all_group_messages,omitempty"`
	SupportsInlineQueries   *bool   `json:"supports_inline_queries,omitempty"`
	CanConnectToBusiness    *bool   `json:"can_connect_to_business,omitempty"`
	HasMainWebApp           *bool   `json:"has_main_web_app,omitempty"`
}

type TelegramBotUpdate struct {
	UpdateId     int                   `json:"update_id"`
	Message      *TelegramBotMessage   `json:"message,omitempty"`
	MyChatMember *TelegramMyChatMember `json:"my_chat_member,omitempty"`
}

type TelegramBotMessage struct {
	MessageId          int                  `json:"message_id"`
	From               TelegramBotUser      `json:"from"`
	Chat               TelegramBotChat      `json:"chat"`
	Date               int64                `json:"date"`
	Text               *string              `json:"text,omitempty"`
	Entities           *[]TelegramBotEntity `json:"entities,omitempty"`
	NewChatParticipant *TelegramBotUser     `json:"new_chat_participant,omitempty"`
	NewChatMember      *TelegramBotUser     `json:"new_chat_member,omitempty"`
	NewChatMembers     *[]TelegramBotUser   `json:"new_chat_members,omitempty"`
}

type TelegramBotChat struct {
	Id                          int     `json:"id"`
	FirstName                   *string `json:"first_name,omitempty"`
	LastName                    *string `json:"last_name,omitempty"`
	Username                    *string `json:"username,omitempty"`
	Title                       *string `json:"title,omitempty"`
	Type                        string  `json:"type"`
	AllMembersAreAdministrators *bool   `json:"all_members_are_administrators,omitempty"`
}

type TelegramBotEntity struct {
	Type   string `json:"type"`
	Offset int    `json:"offset"`
	Length int    `json:"length"`
}

type TelegramMyChatMember struct {
	Chat          TelegramBotChat `json:"chat"`
	From          TelegramBotUser `json:"from"`
	Date          int64           `json:"date"`
	OldChatMember struct {
		User   TelegramBotUser `json:"user"`
		Status string          `json:"status"`
	} `json:"old_chat_member"`
	NewChatMember struct {
		User   TelegramBotUser `json:"user"`
		Status string          `json:"status"`
	} `json:"new_chat_member"`
}

type TelegramBotGetMeResponse struct {
	TelegramBotResponse
	Result TelegramBotUser `json:"result"`
}

type TelegramBotGetUpdatesResponse struct {
	TelegramBotResponse
	Result []TelegramBotUpdate `json:"result"`
}

func NewTelegramBotClient(config TelegramBotConfig) *TelegramBotClient {
	return &TelegramBotClient{
		Config: config,
	}
}

func (t *TelegramBotClient) GetMe() (*TelegramBotGetMeResponse, error) {
	var result TelegramBotGetMeResponse

	url := fmt.Sprintf("%s%s/getMe", t.Config.ApiUrl, t.Config.Token)
	_, err := goutil.SendHttpGet(url, nil, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (t *TelegramBotClient) GetUpdates() (*TelegramBotGetUpdatesResponse, error) {
	var result TelegramBotGetUpdatesResponse

	url := fmt.Sprintf("%s%s/getUpdates", t.Config.ApiUrl, t.Config.Token)
	_, err := goutil.SendHttpGet(url, nil, nil, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (t *TelegramBotClient) SendMessage(chatId int, text string) (interface{}, error) {
	var result interface{}

	url := fmt.Sprintf("%s%s/sendMessage", t.Config.ApiUrl, t.Config.Token)
	data := map[string]interface{}{
		"chat_id":    chatId,
		"text":       text,
		"parse_mode": "HTML",
	}
	_, err := goutil.SendHttpPost(url, data, &map[string]string{
		"Content-Type": "application/json",
	}, &result, nil)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
