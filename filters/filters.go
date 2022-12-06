package filters

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/KeralaBots/GoTGramBot/types"
)

type FilterResponse struct {
	Type     string
	Data     string
	Prefixes []rune
}

var All FilterResponse = FilterResponse{Type: "all"}
var Document FilterResponse = FilterResponse{Type: "document"}
var Audio FilterResponse = FilterResponse{Type: "audio"}
var Video FilterResponse = FilterResponse{Type: "video"}
var VideoNote FilterResponse = FilterResponse{Type: "video_note"}
var Voice FilterResponse = FilterResponse{Type: "voice"}
var Sticker FilterResponse = FilterResponse{Type: "sticker"}
var Animation FilterResponse = FilterResponse{Type: "animation"}
var ViaBot FilterResponse = FilterResponse{Type: "via_bot"}
var Poll FilterResponse = FilterResponse{Type: "poll"}
var Caption FilterResponse = FilterResponse{Type: "caption"}
var Dice FilterResponse = FilterResponse{Type: "dice"}
var Game FilterResponse = FilterResponse{Type: "game"}
var Venue FilterResponse = FilterResponse{Type: "venue"}
var Location FilterResponse = FilterResponse{Type: "location"}
var NewChatTitle FilterResponse = FilterResponse{Type: "new_chat_title"}
var NewChatPhoto FilterResponse = FilterResponse{Type: "new_chat_photo"}
var Invoice FilterResponse = FilterResponse{Type: "invoice"}
var HasProtectedContent FilterResponse = FilterResponse{Type: "has_protected_content"}
var VideoChatScheduled FilterResponse = FilterResponse{Type: "video_chat_scheduled"}
var VideoChatStarted FilterResponse = FilterResponse{Type: "video_chat_started"}
var VideoChatEnded FilterResponse = FilterResponse{Type: "video_chat_ended"}
var VideoChatParticipantsInvited FilterResponse = FilterResponse{Type: "video_chat_participants_invited"}
var SuccessfulPayment FilterResponse = FilterResponse{Type: "successful_payment"}
var Private FilterResponse = FilterResponse{Type: "chat", Data: "private"}
var Group FilterResponse = FilterResponse{Type: "chat", Data: "group"}
var SuperGroup FilterResponse = FilterResponse{Type: "chat", Data: "supergroup"}
var Channel FilterResponse = FilterResponse{Type: "chat", Data: "channel"}

func Command(command string, prefixes []rune) FilterResponse {
	if prefixes == nil {
		prefixes = []rune{'/'}
	}
	data := FilterResponse{
		Type:     "command",
		Data:     command,
		Prefixes: prefixes,
	}

	return data
}

func Regex(regex string) FilterResponse {
	data := FilterResponse{
		Type: "regex",
		Data: regex,
	}

	return data
}

func CallbackData(data string) FilterResponse {
	return FilterResponse{
		Type: "callback_data",
		Data: data,
	}
}

func (f *FilterResponse) CheckMessage(m *types.Message) bool {
	rawUpdate, err := json.Marshal(m)
	if err != nil {
		return false
	}

	u := make(map[string]interface{})
	err = json.Unmarshal(rawUpdate, &u)
	if err != nil {
		return false
	}

	res := false
	if u["text"] != nil {
		text := fmt.Sprintf("%v", u["text"])

		if f.Type == "command" {
			for _, p := range f.Prefixes {
				if strings.HasPrefix(text, string(p)) {
					re, _ := regexp.MatchString(f.Data, text)
					if re {
						res = true
					}
				}
			}
		}

		if f.Type == "regex" {
			re, _ := regexp.MatchString(f.Data, text)
			if re {
				res = true
			}
		}
	} else if u[f.Type] != nil {
		res = true
	} else {
		res = false
	}

	if f.Type == "all" {
		res = true
	}

	if f.Type == "chat" {
		if m.Chat.Type == f.Data {
			res = true
		} else {
			res = false
		}
	}

	return res
}

func (f *FilterResponse) CheckCallback(m *types.CallbackQuery) bool {
	res := false

	if f.Type == "callback_data" {
		re, _ := regexp.MatchString(f.Data, m.Data)
		if re {
			res = true
		}
	} else if f.Type == "regex" {
		re, _ := regexp.MatchString(f.Data, m.Data)
		if re {
			res = true
		}
	} else {
		res = false
	}

	if f.Type == "all" {
		res = true
	}

	if f.Type == "chat" {
		if m.Message.Chat.Type == f.Data {
			res = true
		} else {
			res = true
		}
	}

	return res
}
