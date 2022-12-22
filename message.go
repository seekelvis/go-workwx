package workwx

import (
	"errors"
)

type ButtonInteractionContent struct {
	// Required
	CardType        string           `json:"card_type"`
	TaskId          string           `json:"task_id"`
	Source          *Source          `json:"source"`
	MainTitle       *MainTitle       `json:"main_title"`
	ButtonSelection *ButtonSelection `json:"button_selection"`
	ButtonList      []*ButtonList    `json:"button_list"`
	// Optional
	ActionMenu            *ActionMenu              `json:"action_menu"`
	QuoteArea             *QuoteArea               `json:"quote_area"`
	SubTitleText          string                   `json:"sub_title_text"`
	HorizontalContentList []*HorizontalContentList `json:"horizontal_content_list"`
	CardAction            *CardAction              `json:"card_action"`
}

type Source struct {
	IconUrl   string `json:"icon_url"`
	Desc      string `json:"desc"`
	DescColor int    `json:"desc_color"` // 0(默认) 灰色，1 黑色，2 红色，3 绿色
}

type MainTitle struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

type ButtonSelection struct {
	QuestionKey string                      `json:"question_key"`
	Title       string                      `json:"title"`
	OptionList  []ButtonSelectionOptionList `json:"option_list"`
	SelectedId  string                      `json:"selected_id"`
}

type ButtonSelectionOptionList struct {
	Id   string `json:"id"`
	Text string `json:"text"`
}

type ButtonList struct {
	Type  int    `json:"type"` // 0 或不填代表回调点击事件，1 代表跳转url
	Text  string `json:"text"`
	Key   string `json:"key"`
	Style int    `json:"style"` // 按钮样式，目前可填1~4，不填或错填默认1
	Url   string `json:"url"`   // 跳转事件的url，button_list.type是1时必填
}

type ActionMenu struct {
	Desc       string                 `json:"desc"`
	ActionList []ActionMenuActionList `json:"action_list"`
}

type ActionMenuActionList struct {
	Text string `json:"text"`
	Key  string `json:"key"`
}

type QuoteArea struct {
	Type      int    `json:"type"` // 0或不填代表没有点击事件，1 代表跳转url，2 代表跳转小程序
	Title     string `json:"title"`
	QuoteText string `json:"quote_text"`
	URL       string `json:"url"`      // 点击跳转的url，quote_area.type是1时必填
	Appid     string `json:"appid"`    // 点击跳转的小程序的appid，必须是与当前应用关联的小程序，quote_area.type是2时必填
	Pagepath  string `json:"pagepath"` // 点击跳转的小程序的pagepath，quote_area.type是2时选填
}

type HorizontalContentList struct {
	Type    int    `json:"type"` // 链接类型，0或不填代表不是链接，1 代表跳转url，2 代表下载附件，3 代表点击跳转成员详情
	Keyname string `json:"keyname"`
	Value   string `json:"value"`
	Url     string `json:"url"`      // 链接跳转的url，horizontal_content_list.type是1时必填
	MediaId string `json:"media_id"` // 附件的media_id，horizontal_content_list.type是2时必填
	Userid  string `json:"userid"`   // 成员详情的userid，horizontal_content_list.type是3时必填
}

type CardAction struct {
	Type     int    `json:"type"`     // 跳转事件类型，0或不填代表不是链接，1 代表跳转url，2 代表打开小程序
	Url      string `json:"url"`      // 跳转事件的url，card_action.type是1时必填
	AppId    string `json:"app_id"`   // 跳转事件的小程序的appid，必须是与当前应用关联的小程序，card_action.type是2时必填
	Pagepath string `json:"pagepath"` // 跳转事件的小程序的pagepath，card_action.type是2时选填
}

// SendTextMessage 发送文本消息
//
// 收件人参数如果仅设置了 `ChatID` 字段，则为【发送消息到群聊会话】接口调用；
// 否则为单纯的【发送应用消息】接口调用。
func (c *WorkwxApp) SendTextMessage(
	recipient *Recipient,
	content string,
	isSafe bool,
) error {
	return c.sendMessage(recipient, "text", map[string]interface{}{"content": content}, isSafe)
}

// SendImageMessage 发送图片消息
//
// 收件人参数如果仅设置了 `ChatID` 字段，则为【发送消息到群聊会话】接口调用；
// 否则为单纯的【发送应用消息】接口调用。
func (c *WorkwxApp) SendImageMessage(
	recipient *Recipient,
	mediaID string,
	isSafe bool,
) error {
	return c.sendMessage(
		recipient,
		"image",
		map[string]interface{}{
			"media_id": mediaID,
		}, isSafe,
	)
}

// SendVoiceMessage 发送语音消息
//
// 收件人参数如果仅设置了 `ChatID` 字段，则为【发送消息到群聊会话】接口调用；
// 否则为单纯的【发送应用消息】接口调用。
func (c *WorkwxApp) SendVoiceMessage(
	recipient *Recipient,
	mediaID string,
	isSafe bool,
) error {
	return c.sendMessage(
		recipient,
		"voice",
		map[string]interface{}{
			"media_id": mediaID,
		}, isSafe,
	)
}

// SendVideoMessage 发送视频消息
//
// 收件人参数如果仅设置了 `ChatID` 字段，则为【发送消息到群聊会话】接口调用；
// 否则为单纯的【发送应用消息】接口调用。
func (c *WorkwxApp) SendVideoMessage(
	recipient *Recipient,
	mediaID string,
	description string,
	title string,
	isSafe bool,
) error {
	return c.sendMessage(
		recipient,
		"video",
		map[string]interface{}{
			"media_id":    mediaID,
			"description": description, // TODO: 零值
			"title":       title,       // TODO: 零值
		}, isSafe,
	)
}

// SendFileMessage 发送文件消息
//
// 收件人参数如果仅设置了 `ChatID` 字段，则为【发送消息到群聊会话】接口调用；
// 否则为单纯的【发送应用消息】接口调用。
func (c *WorkwxApp) SendFileMessage(
	recipient *Recipient,
	mediaID string,
	isSafe bool,
) error {
	return c.sendMessage(
		recipient,
		"file",
		map[string]interface{}{
			"media_id": mediaID,
		}, isSafe,
	)
}

// SendTextCardMessage 发送文本卡片消息
//
// 收件人参数如果仅设置了 `ChatID` 字段，则为【发送消息到群聊会话】接口调用；
// 否则为单纯的【发送应用消息】接口调用。
func (c *WorkwxApp) SendTextCardMessage(
	recipient *Recipient,
	title string,
	description string,
	url string,
	buttonText string,
	isSafe bool,
) error {
	return c.sendMessage(
		recipient,
		"textcard",
		map[string]interface{}{
			"title":       title,
			"description": description,
			"url":         url,
			"btntxt":      buttonText, // TODO: 零值
		}, isSafe,
	)
}

// SendNewsMessage 发送图文消息
//
// 收件人参数如果仅设置了 `ChatID` 字段，则为【发送消息到群聊会话】接口调用；
// 否则为单纯的【发送应用消息】接口调用。
func (c *WorkwxApp) SendNewsMessage(
	recipient *Recipient,
	title string,
	description string,
	url string,
	picURL string,
	isSafe bool,
) error {
	return c.sendMessage(
		recipient,
		"news",
		map[string]interface{}{
			// TODO: 支持发送多条图文
			"articles": []interface{}{
				map[string]interface{}{
					"title":       title,
					"description": description, // TODO: 零值
					"url":         url,
					"picurl":      picURL, // TODO: 零值
				},
			},
		}, isSafe)
}

// SendMPNewsMessage 发送 mpnews 类型的图文消息
//
// 收件人参数如果仅设置了 `ChatID` 字段，则为【发送消息到群聊会话】接口调用；
// 否则为单纯的【发送应用消息】接口调用。
func (c *WorkwxApp) SendMPNewsMessage(
	recipient *Recipient,
	title string,
	thumbMediaID string,
	author string,
	sourceContentURL string,
	content string,
	digest string,
	isSafe bool,
) error {
	return c.sendMessage(
		recipient,
		"mpnews",
		map[string]interface{}{
			// TODO: 支持发送多条图文
			"articles": []interface{}{
				map[string]interface{}{
					"title":              title,
					"thumb_media_id":     thumbMediaID,
					"author":             author,           // TODO: 零值
					"content_source_url": sourceContentURL, // TODO: 零值
					"content":            content,
					"digest":             digest,
				},
			},
		}, isSafe,
	)
}

// SendMarkdownMessage 发送 Markdown 消息
//
// 仅支持 Markdown 的子集，详见[官方文档](https://work.weixin.qq.com/api/doc#90002/90151/90854/%E6%94%AF%E6%8C%81%E7%9A%84markdown%E8%AF%AD%E6%B3%95)。
//
// 收件人参数如果仅设置了 `ChatID` 字段，则为【发送消息到群聊会话】接口调用；
// 否则为单纯的【发送应用消息】接口调用。
func (c *WorkwxApp) SendMarkdownMessage(
	recipient *Recipient,
	content string,
	isSafe bool,
) error {
	return c.sendMessage(recipient, "markdown", map[string]interface{}{"content": content}, isSafe)
}

// SendTemplateCardButtonMessage 发送模板卡片消息之按钮交互型
//

func (c *WorkwxApp) SendTemplateCardButtonMessage(
	recipient *Recipient,
	content ButtonInteractionContent,

) error {
	body := map[string]interface{}{
		"card_type":        "button_interaction",
		"source":           content.Source,
		"main_title":       content.MainTitle,
		"task_id":          content.TaskId,
		"button_selection": content.ButtonSelection,
		"button_list":      content.ButtonList,
	}
	if content.ActionMenu != nil {
		body["action_menu"] = content.ActionMenu
	}
	if content.QuoteArea != nil {
		body["quote_area"] = content.QuoteArea
	}
	if content.SubTitleText != "" {
		body["sub_title_text"] = content.SubTitleText
	}
	if content.HorizontalContentList != nil {
		body["horizontal_content_list"] = content.HorizontalContentList
	}
	if content.CardAction != nil {
		body["card_action"] = content.CardAction
	}

	return c.sendMessage(
		recipient,
		"template_card",
		body,
		false,
	)
}

// sendMessage 发送消息底层接口
//
// 收件人参数如果仅设置了 `ChatID` 字段，则为【发送消息到群聊会话】接口调用；
// 否则为单纯的【发送应用消息】接口调用。
func (c *WorkwxApp) sendMessage(
	recipient *Recipient,
	msgtype string,
	content map[string]interface{},
	isSafe bool,
) error {
	isApichatSendRequest := false
	if !recipient.isValidForMessageSend() {
		if !recipient.isValidForAppchatSend() {
			// TODO: better error
			return errors.New("recipient invalid for message sending")
		}

		// 发送给群聊
		isApichatSendRequest = true
	}

	req := reqMessage{
		ToUser:  recipient.UserIDs,
		ToParty: recipient.PartyIDs,
		ToTag:   recipient.TagIDs,
		ChatID:  recipient.ChatID,
		AgentID: c.AgentID,
		MsgType: msgtype,
		Content: content,
		IsSafe:  isSafe,
	}

	var resp respMessageSend
	var err error
	if isApichatSendRequest {
		resp, err = c.execAppchatSend(req)
	} else {
		resp, err = c.execMessageSend(req)
	}

	if err != nil {
		return err
	}

	// TODO: what to do with resp?
	_ = resp
	return nil
}
