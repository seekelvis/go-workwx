package workwx

import (
	"errors"
)

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
	content string,
) error {
	return c.sendMessage(
		recipient,
		"template_card",
		map[string]interface{}{
			// TODO: 支持发送多条图文
			"template_card": map[string]interface{}{
				"card_type": "card_type",
				"source": map[string]interface{}{
					"icon_url":   "icon_url",
					"desc":       "desc",
					"desc_color": 1,
				},
				"action_menu": map[string]interface{}{
					"desc": "desc",
					"action_list": []interface{}{
						map[string]interface{}{
							"text": "接受推送",
							"key":  "A",
						},
						map[string]interface{}{
							"text": "拒绝推送",
							"key":  "B",
						},
					},
				},
				"main_title": map[string]interface{}{
					"title": "欢迎使用企业微信",
					"desc":  "您的好友正在邀请您加入企业微信",
				},
				"quote_area": map[string]interface{}{
					"type":       1,
					"url":        "https://work.weixin.qq.com",
					"title":      "企业微信的引用样式",
					"quote_text": "企业微信真好用呀真好用",
				},
				"sub_title_text": "下载企业微信还能抢红包！",
				"horizontal_content_list": []interface{}{
					map[string]interface{}{
						"keyname": "邀请人",
						"value":   "张三",
					},
					map[string]interface{}{
						"type":    1,
						"keyname": "企业微信官网",
						"value":   "点击访问",
						"url":     "https://work.weixin.qq.com",
					},
					map[string]interface{}{
						"type":     2,
						"keyname":  "企业微信下载",
						"value":    "企业微信.apk",
						"media_id": "文件的media_id",
					},
					map[string]interface{}{
						"type":    3,
						"keyname": "员工信息",
						"value":   "点击查看",
						"userid":  "zhangsan",
					},
				},
				"card_action": map[string]interface{}{
					"type":     2,
					"url":      "https://work.weixin.qq.com",
					"appid":    "小程序的appid",
					"pagepath": "/index.html",
				},
				"task_id": "task_id",
				"button_selection": map[string]interface{}{
					"question_key": "btn_question_key1",
					"title":        "企业微信评分",
					"option_list": []interface{}{
						map[string]interface{}{
							"id":   "btn_selection_id1",
							"text": "100分",
						},
						map[string]interface{}{
							"id":   "btn_selection_id2",
							"text": "101分",
						},
					},
					"selected_id": "btn_selection_id1",
				},
				"button_list": []interface{}{
					map[string]interface{}{
						"text":  "按钮1",
						"style": 1,
						"key":   "button_key_1",
					},
					map[string]interface{}{
						"text":  "按钮2",
						"style": 2,
						"key":   "button_key_2",
					},
				},
			},
		}, false,
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
