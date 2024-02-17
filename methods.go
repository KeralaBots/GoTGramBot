package bot

import (
    "os"
    "fmt"
    "strconv"
    "encoding/json"
    "github.com/KeralaBots/GoTGramBot/types"
)


// GetUpdates methods's optional params
type GetUpdatesOpts struct {
    Offset int64 `json:"offset,omitempty"`
    Limit int64 `json:"limit,omitempty"`
    Timeout int64 `json:"timeout,omitempty"`
    AllowedUpdates []string `json:"allowed_updates,omitempty"`
}

// Use this method to receive incoming updates using long polling (wiki). Returns an Array of Update objects.
func (b *Bot) GetUpdates(opts *GetUpdatesOpts) ([]types.Update, error) {
    params := map[string]string{}
    data_params := map[string]string{}

    if opts != nil {
        params["offset"] = strconv.FormatInt(opts.Offset, 10)
        params["limit"] = strconv.FormatInt(opts.Limit, 10)
        params["timeout"] = strconv.FormatInt(opts.Timeout, 10)

        if opts.AllowedUpdates != nil {
            bs, err := json.Marshal(opts.AllowedUpdates)
            if err != nil {
                return nil, fmt.Errorf("failed to marshal field allowed_updates: %w", err)
            }
            params["allowed_updates"] = string(bs)
        }

    }


    r, err := b.Request("getUpdates", params, data_params)
    if err != nil {
        return nil, err
    }

    
    var res []types.Update
    return res, json.Unmarshal(r, &res) 

}

// SetWebhook methods's optional params
type SetWebhookOpts struct {
    Certificate types.InputFile `json:"certificate,omitempty"`
    IpAddress string `json:"ip_address,omitempty"`
    MaxConnections int64 `json:"max_connections,omitempty"`
    AllowedUpdates []string `json:"allowed_updates,omitempty"`
    DropPendingUpdates bool `json:"drop_pending_updates,omitempty"`
    SecretToken string `json:"secret_token,omitempty"`
}

// Use this method to specify a URL and receive incoming updates via an outgoing webhook. Whenever there is an update for the bot, we will send an HTTPS POST request to the specified URL, containing a JSON-serialized Update. In case of an unsuccessful request, we will give up after a reasonable amount of attempts. Returns True on success.
// If you'd like to make sure that the webhook was set by you, you can specify secret data in the parameter secret_token. If specified, the request will contain a header "X-Telegram-Bot-Api-Secret-Token" with the secret token as content.
func (b *Bot) SetWebhook(url string, opts *SetWebhookOpts) (bool, error) {
    params := map[string]string{}
    data_params := map[string]string{}

    params["url"] = url
    if opts != nil {

        if opts.Certificate != nil {
            switch f := opts.Certificate.(type) {
            case string:
                _, err := os.Stat(f)
                if err != nil {
                    params["certificate"] = f
                } else {
                    params["certificate"] = "attach://certificate"
                    data_params["certificate"] = f
                }
            default:
                return false, fmt.Errorf("unknown type for InputFile: %T", opts.Certificate)
            }
        }
        params["ip_address"] = opts.IpAddress
        params["max_connections"] = strconv.FormatInt(opts.MaxConnections, 10)

        if opts.AllowedUpdates != nil {
            bs, err := json.Marshal(opts.AllowedUpdates)
            if err != nil {
                return false, fmt.Errorf("failed to marshal field allowed_updates: %w", err)
            }
            params["allowed_updates"] = string(bs)
        }

        params["drop_pending_updates"] = strconv.FormatBool(opts.DropPendingUpdates)
        params["secret_token"] = opts.SecretToken
    }


    r, err := b.Request("setWebhook", params, data_params)
    if err != nil {
        return false, err
    }

    
    var res bool
    return res, json.Unmarshal(r, &res) 

}

// DeleteWebhook methods's optional params
type DeleteWebhookOpts struct {
    DropPendingUpdates bool `json:"drop_pending_updates,omitempty"`
}

// Use this method to remove webhook integration if you decide to switch back to getUpdates. Returns True on success.
func (b *Bot) DeleteWebhook(opts *DeleteWebhookOpts) (bool, error) {
    params := map[string]string{}
    data_params := map[string]string{}

    if opts != nil {
        params["drop_pending_updates"] = strconv.FormatBool(opts.DropPendingUpdates)
    }


    r, err := b.Request("deleteWebhook", params, data_params)
    if err != nil {
        return false, err
    }

    
    var res bool
    return res, json.Unmarshal(r, &res) 

}

// Use this method to get current webhook status. Requires no parameters. On success, returns a WebhookInfo object. If the bot is using getUpdates, will return an object with the url field empty.
func (b *Bot) GetWebhookInfo() (*types.WebhookInfo, error) {
    params := map[string]string{}
    data_params := map[string]string{}

    r, err := b.Request("getWebhookInfo", params, data_params)
    if err != nil {
        return nil, err
    }
    
    var res *types.WebhookInfo
    return res, json.Unmarshal(r, &res) 

}

// A simple method for testing your bot's authentication token. Requires no parameters. Returns basic information about the bot in form of a User object.
func (b *Bot) GetMe() (*types.User, error) {
    params := map[string]string{}
    data_params := map[string]string{}

    r, err := b.Request("getMe", params, data_params)
    if err != nil {
        return nil, err
    }
    
    var res *types.User
    return res, json.Unmarshal(r, &res) 

}

// Use this method to log out from the cloud Bot API server before launching the bot locally. You must log out the bot before running it locally, otherwise there is no guarantee that the bot will receive updates. After a successful call, you can immediately log in on a local server, but will not be able to log in back to the cloud Bot API server for 10 minutes. Returns True on success. Requires no parameters.
func (b *Bot) LogOut() (bool, error) {
    params := map[string]string{}
    data_params := map[string]string{}

    r, err := b.Request("logOut", params, data_params)
    if err != nil {
        return false, err
    }
    
    var res bool
    return res, json.Unmarshal(r, &res) 

}

// Use this method to close the bot instance before moving it from one local server to another. You need to delete the webhook before calling this method to ensure that the bot isn't launched again after server restart. The method will return error 429 in the first 10 minutes after the bot is launched. Returns True on success. Requires no parameters.
func (b *Bot) Close() (bool, error) {
    params := map[string]string{}
    data_params := map[string]string{}

    r, err := b.Request("close", params, data_params)
    if err != nil {
        return false, err
    }
    
    var res bool
    return res, json.Unmarshal(r, &res) 

}

// SendMessage methods's optional params
type SendMessageOpts struct {
    MessageThreadId int64 `json:"message_thread_id,omitempty"`
    ParseMode string `json:"parse_mode,omitempty"`
    Entities []types.MessageEntity `json:"entities,omitempty"`
    LinkPreviewOptions *types.LinkPreviewOptions `json:"link_preview_options,omitempty"`
    DisableNotification bool `json:"disable_notification,omitempty"`
    ProtectContent bool `json:"protect_content,omitempty"`
    ReplyParameters *types.ReplyParameters `json:"reply_parameters,omitempty"`
    ReplyMarkup types.ReplyMarkup `json:"reply_markup,omitempty"`
}

// Use this method to send text messages. On success, the sent Message is returned.
func (b *Bot) SendMessage(chatId int64, text string, opts *SendMessageOpts) (*types.Message, error) {
    params := map[string]string{}
    data_params := map[string]string{}

    params["chat_id"] = strconv.FormatInt(chatId, 10)
    params["text"] = text
    if opts != nil {
        params["message_thread_id"] = strconv.FormatInt(opts.MessageThreadId, 10)
        params["parse_mode"] = opts.ParseMode

        if opts.Entities != nil {
            bs, err := json.Marshal(opts.Entities)
            if err != nil {
                return nil, fmt.Errorf("failed to marshal field entities: %w", err)
            }
            params["entities"] = string(bs)
        }


        if opts.LinkPreviewOptions != nil {
            bs, err := json.Marshal(opts.LinkPreviewOptions)
            if err != nil {
                return nil, fmt.Errorf("failed to marshal field link_preview_options: %w", err)
            }
            params["link_preview_options"] = string(bs)
        }

        params["disable_notification"] = strconv.FormatBool(opts.DisableNotification)
        params["protect_content"] = strconv.FormatBool(opts.ProtectContent)

        if opts.ReplyParameters != nil {
            bs, err := json.Marshal(opts.ReplyParameters)
            if err != nil {
                return nil, fmt.Errorf("failed to marshal field reply_parameters: %w", err)
            }
            params["reply_parameters"] = string(bs)
        }


        if opts.ReplyMarkup != nil {
            bs, err := json.Marshal(opts.ReplyMarkup)
            if err != nil {
                return nil, fmt.Errorf("failed to marshal field reply_markup: %w", err)
            }
            params["reply_markup"] = string(bs)
        }

    }


    r, err := b.Request("sendMessage", params, data_params)
    if err != nil {
        return nil, err
    }

    
    var res *types.Message
    return res, json.Unmarshal(r, &res) 

}

// ForwardMessage methods's optional params
type ForwardMessageOpts struct {
    MessageThreadId int64 `json:"message_thread_id,omitempty"`
    DisableNotification bool `json:"disable_notification,omitempty"`
    ProtectContent bool `json:"protect_content,omitempty"`
}

// Use this method to forward messages of any kind. Service messages and messages with protected content can't be forwarded. On success, the sent Message is returned.
func (b *Bot) ForwardMessage(chatId int64, fromChatId int64, messageId int64, opts *ForwardMessageOpts) (*types.Message, error) {
    params := map[string]string{}
    data_params := map[string]string{}

    params["chat_id"] = strconv.FormatInt(chatId, 10)
    params["from_chat_id"] = strconv.FormatInt(fromChatId, 10)
    params["message_id"] = strconv.FormatInt(messageId, 10)
    if opts != nil {
        params["message_thread_id"] = strconv.FormatInt(opts.MessageThreadId, 10)
        params["disable_notification"] = strconv.FormatBool(opts.DisableNotification)
        params["protect_content"] = strconv.FormatBool(opts.ProtectContent)
    }


    r, err := b.Request("forwardMessage", params, data_params)
    if err != nil {
        return nil, err
    }

    
    var res *types.Message
    return res, json.Unmarshal(r, &res) 

}

// ForwardMessages methods's optional params
type ForwardMessagesOpts struct {
    MessageThreadId int64 `json:"message_thread_id,omitempty"`
    DisableNotification bool `json:"disable_notification,omitempty"`
    ProtectContent bool `json:"protect_content,omitempty"`
}

// Use this method to forward multiple messages of any kind. If some of the specified messages can't be found or forwarded, they are skipped. Service messages and messages with protected content can't be forwarded. Album grouping is kept for forwarded messages. On success, an array of MessageId of the sent messages is returned.
func (b *Bot) ForwardMessages(chatId int64, fromChatId int64, messageIds []int64, opts *ForwardMessagesOpts) ([]types.MessageId, error) {
    params := map[string]string{}
    data_params := map[string]string{}

    params["chat_id"] = strconv.FormatInt(chatId, 10)
    params["from_chat_id"] = strconv.FormatInt(fromChatId, 10)

    if messageIds != nil {
        bs, err := json.Marshal(messageIds)
        if err != nil {
            return nil, fmt.Errorf("failed to marshal field message_ids: %w", err)
        }
        params["message_ids"] = string(bs)
    }

    if opts != nil {
        params["message_thread_id"] = strconv.FormatInt(opts.MessageThreadId, 10)
        params["disable_notification"] = strconv.FormatBool(opts.DisableNotification)
        params["protect_content"] = strconv.FormatBool(opts.ProtectContent)
    }


    r, err := b.Request("forwardMessages", params, data_params)
    if err != nil {
        return nil, err
    }

    
    var res []types.MessageId
    return res, json.Unmarshal(r, &res) 

}

// CopyMessage methods's optional params
type CopyMessageOpts struct {
    MessageThreadId int64 `json:"message_thread_id,omitempty"`
    Caption string `json:"caption,omitempty"`
    ParseMode string `json:"parse_mode,omitempty"`
    CaptionEntities []types.MessageEntity `json:"caption_entities,omitempty"`
    DisableNotification bool `json:"disable_notification,omitempty"`
    ProtectContent bool `json:"protect_content,omitempty"`
    ReplyParameters *types.ReplyParameters `json:"reply_parameters,omitempty"`
    ReplyMarkup types.ReplyMarkup `json:"reply_markup,omitempty"`
}

// Use this method to copy messages of any kind. Service messages, giveaway messages, giveaway winners messages, and invoice messages can't be copied. A quiz poll can be copied only if the value of the field correct_option_id is known to the bot. The method is analogous to the method forwardMessage, but the copied message doesn't have a link to the original message. Returns the MessageId of the sent message on success.
func (b *Bot) CopyMessage(chatId int64, fromChatId int64, messageId int64, opts *CopyMessageOpts) (*types.MessageId, error) {
    params := map[string]string{}
    data_params := map[string]string{}

    params["chat_id"] = strconv.FormatInt(chatId, 10)
    params["from_chat_id"] = strconv.FormatInt(fromChatId, 10)
    params["message_id"] = strconv.FormatInt(messageId, 10)
    if opts != nil {
        params["message_thread_id"] = strconv.FormatInt(opts.MessageThreadId, 10)
        params["caption"] = opts.Caption
        params["parse_mode"] = opts.ParseMode

        if opts.CaptionEntities != nil {
            bs, err := json.Marshal(opts.CaptionEntities)
            if err != nil {
                return nil, fmt.Errorf("failed to marshal field caption_entities: %w", err)
            }
            params["caption_entities"] = string(bs)
        }

        params["disable_notification"] = strconv.FormatBool(opts.DisableNotification)
        params["protect_content"] = strconv.FormatBool(opts.ProtectContent)

        if opts.ReplyParameters != nil {
            bs, err := json.Marshal(opts.ReplyParameters)
            if err != nil {
                return nil, fmt.Errorf("failed to marshal field reply_parameters: %w", err)
            }
            params["reply_parameters"] = string(bs)
        }


        if opts.ReplyMarkup != nil {
            bs, err := json.Marshal(opts.ReplyMarkup)
            if err != nil {
                return nil, fmt.Errorf("failed to marshal field reply_markup: %w", err)
            }
            params["reply_markup"] = string(bs)
        }

    }


    r, err := b.Request("copyMessage", params, data_params)
    if err != nil {
        return nil, err
    }

    
    var res *types.MessageId
    return res, json.Unmarshal(r, &res) 

}

// CopyMessages methods's optional params
type CopyMessagesOpts struct {
    MessageThreadId int64 `json:"message_thread_id,omitempty"`
    DisableNotification bool `json:"disable_notification,omitempty"`
    ProtectContent bool `json:"protect_content,omitempty"`
    RemoveCaption bool `json:"remove_caption,omitempty"`
}

// Use this method to copy messages of any kind. If some of the specified messages can't be found or copied, they are skipped. Service messages, giveaway messages, giveaway winners messages, and invoice messages can't be copied. A quiz poll can be copied only if the value of the field correct_option_id is known to the bot. The method is analogous to the method forwardMessages, but the copied messages don't have a link to the original message. Album grouping is kept for copied messages. On success, an array of MessageId of the sent messages is returned.
func (b *Bot) CopyMessages(chatId int64, fromChatId int64, messageIds []int64, opts *CopyMessagesOpts) ([]types.MessageId, error) {
    params := map[string]string{}
    data_params := map[string]string{}

    params["chat_id"] = strconv.FormatInt(chatId, 10)
    params["from_chat_id"] = strconv.FormatInt(fromChatId, 10)

    if messageIds != nil {
        bs, err := json.Marshal(messageIds)
        if err != nil {
            return nil, fmt.Errorf("failed to marshal field message_ids: %w", err)
        }
        params["message_ids"] = string(bs)
    }

    if opts != nil {
        params["message_thread_id"] = strconv.FormatInt(opts.MessageThreadId, 10)
        params["disable_notification"] = strconv.FormatBool(opts.DisableNotification)
        params["protect_content"] = strconv.FormatBool(opts.ProtectContent)
        params["remove_caption"] = strconv.FormatBool(opts.RemoveCaption)
    }


    r, err := b.Request("copyMessages", params, data_params)
    if err != nil {
        return nil, err
    }

    
    var res []types.MessageId
    return res, json.Unmarshal(r, &res) 

}

// SendPhoto methods's optional params
type SendPhotoOpts struct {
    MessageThreadId int64 `json:"message_thread_id,omitempty"`
    Caption string `json:"caption,omitempty"`
    ParseMode string `json:"parse_mode,omitempty"`
    CaptionEntities []types.MessageEntity `json:"caption_entities,omitempty"`
    HasSpoiler bool `json:"has_spoiler,omitempty"`
    DisableNotification bool `json:"disable_notification,omitempty"`
    ProtectContent bool `json:"protect_content,omitempty"`
    ReplyParameters *types.ReplyParameters `json:"reply_parameters,omitempty"`
    ReplyMarkup types.ReplyMarkup `json:"reply_markup,omitempty"`
}

// Use this method to send photos. On success, the sent Message is returned.
func (b *Bot) SendPhoto(chatId int64, photo types.InputFile, opts *SendPhotoOpts) (*types.Message, error) {
    params := map[string]string{}
    data_params := map[string]string{}

    params["chat_id"] = strconv.FormatInt(chatId, 10)

    if photo != nil {
        switch f := photo.(type) {
        case string:
            _, err := os.Stat(f)
            if err != nil {
                params["photo"] = f
            } else {
                params["photo"] = "attach://photo"
                data_params["photo"] = f
            }
        default:
            return nil, fmt.Errorf("unknown type for InputFile: %T", photo)
        }
    }
    if opts != nil {
        params["message_thread_id"] = strconv.FormatInt(opts.MessageThreadId, 10)
        params["caption"] = opts.Caption
        params["parse_mode"] = opts.ParseMode

        if opts.CaptionEntities != nil {
            bs, err := json.Marshal(opts.CaptionEntities)
            if err != nil {
                return nil, fmt.Errorf("failed to marshal field caption_entities: %w", err)
            }
            params["caption_entities"] = string(bs)
        }

        params["has_spoiler"] = strconv.FormatBool(opts.HasSpoiler)
        params["disable_notification"] = strconv.FormatBool(opts.DisableNotification)
        params["protect_content"] = strconv.FormatBool(opts.ProtectContent)

        if opts.ReplyParameters != nil {
            bs, err := json.Marshal(opts.ReplyParameters)
            if err != nil {
                return nil, fmt.Errorf("failed to marshal field reply_parameters: %w", err)
            }
            params["reply_parameters"] = string(bs)
        }


        if opts.ReplyMarkup != nil {
            bs, err := json.Marshal(opts.ReplyMarkup)
            if err != nil {
                return nil, fmt.Errorf("failed to marshal field reply_markup: %w", err)
            }
            params["reply_markup"] = string(bs)
        }

    }


    r, err := b.Request("sendPhoto", params, data_params)
    if err != nil {
        return nil, err
    }

    
    var res *types.Message
    return res, json.Unmarshal(r, &res) 

}

// SendAudio methods's optional params
type SendAudioOpts struct {
    MessageThreadId int64 `json:"message_thread_id,omitempty"`
    Caption string `json:"caption,omitempty"`
    ParseMode string `json:"parse_mode,omitempty"`
    CaptionEntities []types.MessageEntity `json:"caption_entities,omitempty"`
    Duration int64 `json:"duration,omitempty"`
    Performer string `json:"performer,omitempty"`
    Title string `json:"title,omitempty"`
    Thumbnail types.InputFile `json:"thumbnail,omitempty"`
    DisableNotification bool `json:"disable_notification,omitempty"`
    ProtectContent bool `json:"protect_content,omitempty"`
    ReplyParameters *types.ReplyParameters `json:"reply_parameters,omitempty"`
    ReplyMarkup types.ReplyMarkup `json:"reply_markup,omitempty"`
}

// Use this method to send audio files, if you want Telegram clients to display them in the music player. Your audio must be in the .MP3 or .M4A format. On success, the sent Message is returned. Bots can currently send audio files of up to 50 MB in size, this limit may be changed in the future.
// For sending voice messages, use the sendVoice method instead.
func (b *Bot) SendAudio(chatId int64, audio types.InputFile, opts *SendAudioOpts) (*types.Message, error) {
    params := map[string]string{}
    data_params := map[string]string{}

    params["chat_id"] = strconv.FormatInt(chatId, 10)

    if audio != nil {
        switch f := audio.(type) {
        case string:
            _, err := os.Stat(f)
            if err != nil {
                params["audio"] = f
            } else {
                params["audio"] = "attach://audio"
                data_params["audio"] = f
            }
        default:
            return nil, fmt.Errorf("unknown type for InputFile: %T", audio)
        }
    }
    if opts != nil {
        params["message_thread_id"] = strconv.FormatInt(opts.MessageThreadId, 10)
        params["caption"] = opts.Caption
        params["parse_mode"] = opts.ParseMode

        if opts.CaptionEntities != nil {
            bs, err := json.Marshal(opts.CaptionEntities)
            if err != nil {
                return nil, fmt.Errorf("failed to marshal field caption_entities: %w", err)
            }
            params["caption_entities"] = string(bs)
        }

        params["duration"] = strconv.FormatInt(opts.Duration, 10)
        params["performer"] = opts.Performer
        params["title"] = opts.Title

        if opts.Thumbnail != nil {
            switch f := opts.Thumbnail.(type) {
            case string:
                _, err := os.Stat(f)
                if err != nil {
                    params["thumbnail"] = f
                } else {
                    params["thumbnail"] = "attach://thumbnail"
                    data_params["thumbnail"] = f
                }
            default:
                return nil, fmt.Errorf("unknown type for InputFile: %T", opts.Thumbnail)
            }
        }
        params["disable_notification"] = strconv.FormatBool(opts.DisableNotification)
        params["protect_content"] = strconv.FormatBool(opts.ProtectContent)

        if opts.ReplyParameters != nil {
            bs, err := json.Marshal(opts.ReplyParameters)
            if err != nil {
                return nil, fmt.Errorf("failed to marshal field reply_parameters: %w", err)
            }
            params["reply_parameters"] = string(bs)
        }


        if opts.ReplyMarkup != nil {
            bs, err := json.Marshal(opts.ReplyMarkup)
            if err != nil {
                return nil, fmt.Errorf("failed to marshal field reply_markup: %w", err)
            }
            params["reply_markup"] = string(bs)
        }

    }


    r, err := b.Request("sendAudio", params, data_params)
    if err != nil {
        return nil, err
    }

    
    var res *types.Message
    return res, json.Unmarshal(r, &res) 

}

// SendDocument methods's optional params
type SendDocumentOpts struct {
    MessageThreadId int64 `json:"message_thread_id,omitempty"`
    Thumbnail types.InputFile `json:"thumbnail,omitempty"`
    Caption string `json:"caption,omitempty"`
    ParseMode string `json:"parse_mode,omitempty"`
    CaptionEntities []types.MessageEntity `json:"caption_entities,omitempty"`
    DisableContentTypeDetection bool `json:"disable_content_type_detection,omitempty"`
    DisableNotification bool `json:"disable_notification,omitempty"`
    ProtectContent bool `json:"protect_content,omitempty"`
    ReplyParameters *types.ReplyParameters `json:"reply_parameters,omitempty"`
    ReplyMarkup types.ReplyMarkup `json:"reply_markup,omitempty"`
}

// Use this method to send general files. On success, the sent Message is returned. Bots can currently send files of any type of up to 50 MB in size, this limit may be changed in the future.
func (b *Bot) SendDocument(chatId int64, document types.InputFile, opts *SendDocumentOpts) (*types.Message, error) {
    params := map[string]string{}
    data_params := map[string]string{}

    params["chat_id"] = strconv.FormatInt(chatId, 10)

    if document != nil {
        switch f := document.(type) {
        case string:
            _, err := os.Stat(f)
            if err != nil {
                params["document"] = f
            } else {
                params["document"] = "attach://document"
                data_params["document"] = f
            }
        default:
            return nil, fmt.Errorf("unknown type for InputFile: %T", document)
        }
    }
    if opts != nil {
        params["message_thread_id"] = strconv.FormatInt(opts.MessageThreadId, 10)

        if opts.Thumbnail != nil {
            switch f := opts.Thumbnail.(type) {
            case string:
                _, err := os.Stat(f)
                if err != nil {
                    params["thumbnail"] = f
                } else {
                    params["thumbnail"] = "attach://thumbnail"
                    data_params["thumbnail"] = f
                }
            default:
                return nil, fmt.Errorf("unknown type for InputFile: %T", opts.Thumbnail)
            }
        }
        params["caption"] = opts.Caption
        params["parse_mode"] = opts.ParseMode

        if opts.CaptionEntities != nil {
            bs, err := json.Marshal(opts.CaptionEntities)
            if err != nil {
                return nil, fmt.Errorf("failed to marshal field caption_entities: %w", err)
            }
            params["caption_entities"] = string(bs)
        }

        params["disable_content_type_detection"] = strconv.FormatBool(opts.DisableContentTypeDetection)
        params["disable_notification"] = strconv.FormatBool(opts.DisableNotification)
        params["protect_content"] = strconv.FormatBool(opts.ProtectContent)

        if opts.ReplyParameters != nil {
            bs, err := json.Marshal(opts.ReplyParameters)
            if err != nil {
                return nil, fmt.Errorf("failed to marshal field reply_parameters: %w", err)
            }
            params["reply_parameters"] = string(bs)
        }


        if opts.ReplyMarkup != nil {
            bs, err := json.Marshal(opts.ReplyMarkup)
            if err != nil {
                return nil, fmt.Errorf("failed to marshal field reply_markup: %w", err)
            }
            params["reply_markup"] = string(bs)
        }

    }


    r, err := b.Request("sendDocument", params, data_params)
    if err != nil {
        return nil, err
    }

    
    var res *types.Message
    return res, json.Unmarshal(r, &res) 

}

// SendVideo methods's optional params
type SendVideoOpts struct {
    MessageThreadId int64 `json:"message_thread_id,omitempty"`
    Duration int64 `json:"duration,omitempty"`
    Width int64 `json:"width,omitempty"`
    Height int64 `json:"height,omitempty"`
    Thumbnail types.InputFile `json:"thumbnail,omitempty"`
    Caption string `json:"caption,omitempty"`
    ParseMode string `json:"parse_mode,omitempty"`
    CaptionEntities []types.MessageEntity `json:"caption_entities,omitempty"`
    HasSpoiler bool `json:"has_spoiler,omitempty"`
    SupportsStreaming bool `json:"supports_streaming,omitempty"`
    DisableNotification bool `json:"disable_notification,omitempty"`
    ProtectContent bool `json:"protect_content,omitempty"`
    ReplyParameters *types.ReplyParameters `json:"reply_parameters,omitempty"`
    ReplyMarkup types.ReplyMarkup `json:"reply_markup,omitempty"`
}

// Use this method to send video files, Telegram clients support MPEG4 videos (other formats may be sent as Document). On success, the sent Message is returned. Bots can currently send video files of up to 50 MB in size, this limit may be changed in the future.
func (b *Bot) SendVideo(chatId int64, video types.InputFile, opts *SendVideoOpts) (*types.Message, error) {
    params := map[string]string{}
    data_params := map[string]string{}

    params["chat_id"] = strconv.FormatInt(chatId, 10)

    if video != nil {
        switch f := video.(type) {
        case string:
            _, err := os.Stat(f)
            if err != nil {
                params["video"] = f
            } else {
                params["video"] = "attach://video"
                data_params["video"] = f
            }
        default:
            return nil, fmt.Errorf("unknown type for InputFile: %T", video)
        }
    }
    if opts != nil {
        params["message_thread_id"] = strconv.FormatInt(opts.MessageThreadId, 10)
        params["duration"] = strconv.FormatInt(opts.Duration, 10)
        params["width"] = strconv.FormatInt(opts.Width, 10)
        params["height"] = strconv.FormatInt(opts.Height, 10)

        if opts.Thumbnail != nil {
            switch f := opts.Thumbnail.(type) {
            case string:
                _, err := os.Stat(f)
                if err != nil {
                    params["thumbnail"] = f
                } else {
                    params["thumbnail"] = "attach://thumbnail"
                    data_params["thumbnail"] = f
                }
            default:
                return nil, fmt.Errorf("unknown type for InputFile: %T", opts.Thumbnail)
            }
        }
        params["caption"] = opts.Caption
        params["parse_mode"] = opts.ParseMode

        if opts.CaptionEntities != nil {
            bs, err := json.Marshal(opts.CaptionEntities)
            if err != nil {
                return nil, fmt.Errorf("failed to marshal field caption_entities: %w", err)
            }
            params["caption_entities"] = string(bs)
        }

        params["has_spoiler"] = strconv.FormatBool(opts.HasSpoiler)
        params["supports_streaming"] = strconv.FormatBool(opts.SupportsStreaming)
        params["disable_notification"] = strconv.FormatBool(opts.DisableNotification)
        params["protect_content"] = strconv.FormatBool(opts.ProtectContent)

        if opts.ReplyParameters != nil {
            bs, err := json.Marshal(opts.ReplyParameters)
            if err != nil {
                return nil, fmt.Errorf("failed to marshal field reply_parameters: %w", err)
            }
            params["reply_parameters"] = string(bs)
        }


        if opts.ReplyMarkup != nil {
            bs, err := json.Marshal(opts.ReplyMarkup)
            if err != nil {
                return nil, fmt.Errorf("failed to marshal field reply_markup: %w", err)
            }
            params["reply_markup"] = string(bs)
        }

    }


    r, err := b.Request("sendVideo", params, data_params)
    if err != nil {
        return nil, err
    }

    
    var res *types.Message
    return res, json.Unmarshal(r, &res) 

}

// SendAnimation methods's optional params
type SendAnimationOpts struct {
    MessageThreadId int64 `json:"message_thread_id,omitempty"`
    Duration int64 `json:"duration,omitempty"`
    Width int64 `json:"width,omitempty"`
    Height int64 `json:"height,omitempty"`
    Thumbnail types.InputFile `json:"thumbnail,omitempty"`
    Caption string `json:"caption,omitempty"`
    ParseMode string `json:"parse_mode,omitempty"`
    CaptionEntities []types.MessageEntity `json:"caption_entities,omitempty"`
    HasSpoiler bool `json:"has_spoiler,omitempty"`
    DisableNotification bool `json:"disable_notification,omitempty"`
    ProtectContent bool `json:"protect_content,omitempty"`
    ReplyParameters *types.ReplyParameters `json:"reply_parameters,omitempty"`
    ReplyMarkup types.ReplyMarkup `json:"reply_markup,omitempty"`
}

// Use this method to send animation files (GIF or H.264/MPEG-4 AVC video without sound). On success, the sent Message is returned. Bots can currently send animation files of up to 50 MB in size, this limit may be changed in the future.
func (b *Bot) SendAnimation(chatId int64, animation types.InputFile, opts *SendAnimationOpts) (*types.Message, error) {
    params := map[string]string{}
    data_params := map[string]string{}

    params["chat_id"] = strconv.FormatInt(chatId, 10)

    if animation != nil {
        switch f := animation.(type) {
        case string:
            _, err := os.Stat(f)
            if err != nil {
                params["animation"] = f
            } else {
                params["animation"] = "attach://animation"
                data_params["animation"] = f
            }
        default:
            return nil, fmt.Errorf("unknown type for InputFile: %T", animation)
        }
    }
    if opts != nil {
        params["message_thread_id"] = strconv.FormatInt(opts.MessageThreadId, 10)
        params["duration"] = strconv.FormatInt(opts.Duration, 10)
        params["width"] = strconv.FormatInt(opts.Width, 10)
        params["height"] = strconv.FormatInt(opts.Height, 10)

        if opts.Thumbnail != nil {
            switch f := opts.Thumbnail.(type) {
            case string:
                _, err := os.Stat(f)
                if err != nil {
                    params["thumbnail"] = f
                } else {
                    params["thumbnail"] = "attach://thumbnail"
                    data_params["thumbnail"] = f
                }
            default:
                return nil, fmt.Errorf("unknown type for InputFile: %T", opts.Thumbnail)
            }
        }
        params["caption"] = opts.Caption
        params["parse_mode"] = opts.ParseMode

        if opts.CaptionEntities != nil {
            bs, err := json.Marshal(opts.CaptionEntities)
            if err != nil {
                return nil, fmt.Errorf("failed to marshal field caption_entities: %w", err)
            }
            params["caption_entities"] = string(bs)
        }

        params["has_spoiler"] = strconv.FormatBool(opts.HasSpoiler)
        params["disable_notification"] = strconv.FormatBool(opts.DisableNotification)
        params["protect_content"] = strconv.FormatBool(opts.ProtectContent)

        if opts.ReplyParameters != nil {
            bs, err := json.Marshal(opts.ReplyParameters)
            if err != nil {
                return nil, fmt.Errorf("failed to marshal field reply_parameters: %w", err)
            }
            params["reply_parameters"] = string(bs)
        }


        if opts.ReplyMarkup != nil {
            bs, err := json.Marshal(opts.ReplyMarkup)
            if err != nil {
                return nil, fmt.Errorf("failed to marshal field reply_markup: %w", err)
            }
            params["reply_markup"] = string(bs)
        }

    }


    r, err := b.Request("sendAnimation", params, data_params)
    if err != nil {
        return nil, err
    }

    
    var res *types.Message
    return res, json.Unmarshal(r, &res) 

}

// SendVoice methods's optional params
type SendVoiceOpts struct {
    MessageThreadId int64 `json:"message_thread_id,omitempty"`
    Caption string `json:"caption,omitempty"`
    ParseMode string `json:"parse_mode,omitempty"`
    CaptionEntities []types.MessageEntity `json:"caption_entities,omitempty"`
    Duration int64 `json:"duration,omitempty"`
    DisableNotification bool `json:"disable_notification,omitempty"`
    ProtectContent bool `json:"protect_content,omitempty"`
    ReplyParameters *types.ReplyParameters `json:"reply_parameters,omitempty"`
    ReplyMarkup types.ReplyMarkup `json:"reply_markup,omitempty"`
}

// Use this method to send audio files, if you want Telegram clients to display the file as a playable voice message. For this to work, your audio must be in an .OGG file encoded with OPUS (other formats may be sent as Audio or Document). On success, the sent Message is returned. Bots can currently send voice messages of up to 50 MB in size, this limit may be changed in the future.
func (b *Bot) SendVoice(chatId int64, voice types.InputFile, opts *SendVoiceOpts) (*types.Message, error) {
    params := map[string]string{}
    data_params := map[string]string{}

    params["chat_id"] = strconv.FormatInt(chatId, 10)

    if voice != nil {
        switch f := voice.(type) {
        case string:
            _, err := os.Stat(f)
            if err != nil {
                params["voice"] = f
            } else {
                params["voice"] = "attach://voice"
                data_params["voice"] = f
            }
        default:
            return nil, fmt.Errorf("unknown type for InputFile: %T", voice)
        }
    }
    if opts != nil {
        params["message_thread_id"] = strconv.FormatInt(opts.MessageThreadId, 10)
        params["caption"] = opts.Caption
        params["parse_mode"] = opts.ParseMode

        if opts.CaptionEntities != nil {
            bs, err := json.Marshal(opts.CaptionEntities)
            if err != nil {
                return nil, fmt.Errorf("failed to marshal field caption_entities: %w", err)
            }
            params["caption_entities"] = string(bs)
        }

        params["duration"] = strconv.FormatInt(opts.Duration, 10)
        params["disable_notification"] = strconv.FormatBool(opts.DisableNotification)
        params["protect_content"] = strconv.FormatBool(opts.ProtectContent)

        if opts.ReplyParameters != nil {
            bs, err := json.Marshal(opts.ReplyParameters)
            if err != nil {
                return nil, fmt.Errorf("failed to marshal field reply_parameters: %w", err)
            }
            params["reply_parameters"] = string(bs)
        }


        if opts.ReplyMarkup != nil {
            bs, err := json.Marshal(opts.ReplyMarkup)
            if err != nil {
                return nil, fmt.Errorf("failed to marshal field reply_markup: %w", err)
            }
            params["reply_markup"] = string(bs)
        }

    }


    r, err := b.Request("sendVoice", params, data_params)
    if err != nil {
        return nil, err
    }

    
    var res *types.Message
    return res, json.Unmarshal(r, &res) 

}

// SendVideoNote methods's optional params
type SendVideoNoteOpts struct {
    MessageThreadId int64 `json:"message_thread_id,omitempty"`
    Duration int64 `json:"duration,omitempty"`
    Length int64 `json:"length,omitempty"`
    Thumbnail types.InputFile `json:"thumbnail,omitempty"`
    DisableNotification bool `json:"disable_notification,omitempty"`
    ProtectContent bool `json:"protect_content,omitempty"`
    ReplyParameters *types.ReplyParameters `json:"reply_parameters,omitempty"`
    ReplyMarkup types.ReplyMarkup `json:"reply_markup,omitempty"`
}

// As of v.4.0, Telegram clients support rounded square MPEG4 videos of up to 1 minute long. Use this method to send video messages. On success, the sent Message is returned.
func (b *Bot) SendVideoNote(chatId int64, videoNote types.InputFile, opts *SendVideoNoteOpts) (*types.Message, error) {
    params := map[string]string{}
    data_params := map[string]string{}

    params["chat_id"] = strconv.FormatInt(chatId, 10)

    if videoNote != nil {
        switch f := videoNote.(type) {
        case string:
            _, err := os.Stat(f)
            if err != nil {
                params["video_note"] = f
            } else {
                params["video_note"] = "attach://video_note"
                data_params["video_note"] = f
            }
        default:
            return nil, fmt.Errorf("unknown type for InputFile: %T", videoNote)
        }
    }
    if opts != nil {
        params["message_thread_id"] = strconv.FormatInt(opts.MessageThreadId, 10)
        params["duration"] = strconv.FormatInt(opts.Duration, 10)
        params["length"] = strconv.FormatInt(opts.Length, 10)

        if opts.Thumbnail != nil {
            switch f := opts.Thumbnail.(type) {
            case string:
                _, err := os.Stat(f)
                if err != nil {
                    params["thumbnail"] = f
                } else {
                    params["thumbnail"] = "attach://thumbnail"
                    data_params["thumbnail"] = f
                }
            default:
                return nil, fmt.Errorf("unknown type for InputFile: %T", opts.Thumbnail)
            }
        }
        params["disable_notification"] = strconv.FormatBool(opts.DisableNotification)
        params["protect_content"] = strconv.FormatBool(opts.ProtectContent)

        if opts.ReplyParameters != nil {
            bs, err := json.Marshal(opts.ReplyParameters)
            if err != nil {
                return nil, fmt.Errorf("failed to marshal field reply_parameters: %w", err)
            }
            params["reply_parameters"] = string(bs)
        }


        if opts.ReplyMarkup != nil {
            bs, err := json.Marshal(opts.ReplyMarkup)
            if err != nil {
                return nil, fmt.Errorf("failed to marshal field reply_markup: %w", err)
            }
            params["reply_markup"] = string(bs)
        }

    }


    r, err := b.Request("sendVideoNote", params, data_params)
    if err != nil {
        return nil, err
    }

    
    var res *types.Message
    return res, json.Unmarshal(r, &res) 

}

// SendMediaGroup methods's optional params
type SendMediaGroupOpts struct {
    MessageThreadId int64 `json:"message_thread_id,omitempty"`
    DisableNotification bool `json:"disable_notification,omitempty"`
    ProtectContent bool `json:"protect_content,omitempty"`
    ReplyParameters *types.ReplyParameters `json:"reply_parameters,omitempty"`
}

// Use this method to send a group of photos, videos, documents or audios as an album. Documents and audio files can be only grouped in an album with messages of the same type. On success, an array of Messages that were sent is returned.
func (b *Bot) SendMediaGroup(chatId int64, media []types.InputMediaAudio, opts *SendMediaGroupOpts) ([]types.Message, error) {
    params := map[string]string{}
    data_params := map[string]string{}

    params["chat_id"] = strconv.FormatInt(chatId, 10)

    if media != nil {
        bs, err := json.Marshal(media)
        if err != nil {
            return nil, fmt.Errorf("failed to marshal field media: %w", err)
        }
        params["media"] = string(bs)
    }

    if opts != nil {
        params["message_thread_id"] = strconv.FormatInt(opts.MessageThreadId, 10)
        params["disable_notification"] = strconv.FormatBool(opts.DisableNotification)
        params["protect_content"] = strconv.FormatBool(opts.ProtectContent)

        if opts.ReplyParameters != nil {
            bs, err := json.Marshal(opts.ReplyParameters)
            if err != nil {
                return nil, fmt.Errorf("failed to marshal field reply_parameters: %w", err)
            }
            params["reply_parameters"] = string(bs)
        }

    }


    r, err := b.Request("sendMediaGroup", params, data_params)
    if err != nil {
        return nil, err
    }

    
    var res []types.Message
    return res, json.Unmarshal(r, &res) 

}

// SendLocation methods's optional params
type SendLocationOpts struct {
    MessageThreadId int64 `json:"message_thread_id,omitempty"`
    HorizontalAccuracy float64 `json:"horizontal_accuracy,omitempty"`
    LivePeriod int64 `json:"live_period,omitempty"`
    Heading int64 `json:"heading,omitempty"`
    ProximityAlertRadius int64 `json:"proximity_alert_radius,omitempty"`
    DisableNotification bool `json:"disable_notification,omitempty"`
    ProtectContent bool `json:"protect_content,omitempty"`
    ReplyParameters *types.ReplyParameters `json:"reply_parameters,omitempty"`
    ReplyMarkup types.ReplyMarkup `json:"reply_markup,omitempty"`
}

// Use this method to send point on the map. On success, the sent Message is returned.
func (b *Bot) SendLocation(chatId int64, latitude float64, longitude float64, opts *SendLocationOpts) (*types.Message, error) {
    params := map[string]string{}
    data_params := map[string]string{}

    params["chat_id"] = strconv.FormatInt(chatId, 10)
    params["latitude"] = strconv.FormatFloat(latitude, 'E', -1, 64)
    params["longitude"] = strconv.FormatFloat(longitude, 'E', -1, 64)
    if opts != nil {
        params["message_thread_id"] = strconv.FormatInt(opts.MessageThreadId, 10)
        params["horizontal_accuracy"] = strconv.FormatFloat(opts.HorizontalAccuracy, 'E', -1, 64)
        params["live_period"] = strconv.FormatInt(opts.LivePeriod, 10)
        params["heading"] = strconv.FormatInt(opts.Heading, 10)
        params["proximity_alert_radius"] = strconv.FormatInt(opts.ProximityAlertRadius, 10)
        params["disable_notification"] = strconv.FormatBool(opts.DisableNotification)
        params["protect_content"] = strconv.FormatBool(opts.ProtectContent)

        if opts.ReplyParameters != nil {
            bs, err := json.Marshal(opts.ReplyParameters)
            if err != nil {
                return nil, fmt.Errorf("failed to marshal field reply_parameters: %w", err)
            }
            params["reply_parameters"] = string(bs)
        }


        if opts.ReplyMarkup != nil {
            bs, err := json.Marshal(opts.ReplyMarkup)
            if err != nil {
                return nil, fmt.Errorf("failed to marshal field reply_markup: %w", err)
            }
            params["reply_markup"] = string(bs)
        }

    }


    r, err := b.Request("sendLocation", params, data_params)
    if err != nil {
        return nil, err
    }

    
    var res *types.Message
    return res, json.Unmarshal(r, &res) 

}

// SendVenue methods's optional params
type SendVenueOpts struct {
    MessageThreadId int64 `json:"message_thread_id,omitempty"`
    FoursquareId string `json:"foursquare_id,omitempty"`
    FoursquareType string `json:"foursquare_type,omitempty"`
    GooglePlaceId string `json:"google_place_id,omitempty"`
    GooglePlaceType string `json:"google_place_type,omitempty"`
    DisableNotification bool `json:"disable_notification,omitempty"`
    ProtectContent bool `json:"protect_content,omitempty"`
    ReplyParameters *types.ReplyParameters `json:"reply_parameters,omitempty"`
    ReplyMarkup types.ReplyMarkup `json:"reply_markup,omitempty"`
}

// Use this method to send information about a venue. On success, the sent Message is returned.
func (b *Bot) SendVenue(chatId int64, latitude float64, longitude float64, title string, address string, opts *SendVenueOpts) (*types.Message, error) {
    params := map[string]string{}
    data_params := map[string]string{}

    params["chat_id"] = strconv.FormatInt(chatId, 10)
    params["latitude"] = strconv.FormatFloat(latitude, 'E', -1, 64)
    params["longitude"] = strconv.FormatFloat(longitude, 'E', -1, 64)
    params["title"] = title
    params["address"] = address
    if opts != nil {
        params["message_thread_id"] = strconv.FormatInt(opts.MessageThreadId, 10)
        params["foursquare_id"] = opts.FoursquareId
        params["foursquare_type"] = opts.FoursquareType
        params["google_place_id"] = opts.GooglePlaceId
        params["google_place_type"] = opts.GooglePlaceType
        params["disable_notification"] = strconv.FormatBool(opts.DisableNotification)
        params["protect_content"] = strconv.FormatBool(opts.ProtectContent)

        if opts.ReplyParameters != nil {
            bs, err := json.Marshal(opts.ReplyParameters)
            if err != nil {
                return nil, fmt.Errorf("failed to marshal field reply_parameters: %w", err)
            }
            params["reply_parameters"] = string(bs)
        }


        if opts.ReplyMarkup != nil {
            bs, err := json.Marshal(opts.ReplyMarkup)
            if err != nil {
                return nil, fmt.Errorf("failed to marshal field reply_markup: %w", err)
            }
            params["reply_markup"] = string(bs)
        }

    }


    r, err := b.Request("sendVenue", params, data_params)
    if err != nil {
        return nil, err
    }

    
    var res *types.Message
    return res, json.Unmarshal(r, &res) 

}

// SendContact methods's optional params
type SendContactOpts struct {
    MessageThreadId int64 `json:"message_thread_id,omitempty"`
    LastName string `json:"last_name,omitempty"`
    Vcard string `json:"vcard,omitempty"`
    DisableNotification bool `json:"disable_notification,omitempty"`
    ProtectContent bool `json:"protect_content,omitempty"`
    ReplyParameters *types.ReplyParameters `json:"reply_parameters,omitempty"`
    ReplyMarkup types.ReplyMarkup `json:"reply_markup,omitempty"`
}

// Use this method to send phone contacts. On success, the sent Message is returned.
func (b *Bot) SendContact(chatId int64, phoneNumber string, firstName string, opts *SendContactOpts) (*types.Message, error) {
    params := map[string]string{}
    data_params := map[string]string{}

    params["chat_id"] = strconv.FormatInt(chatId, 10)
    params["phone_number"] = phoneNumber
    params["first_name"] = firstName
    if opts != nil {
        params["message_thread_id"] = strconv.FormatInt(opts.MessageThreadId, 10)
        params["last_name"] = opts.LastName
        params["vcard"] = opts.Vcard
        params["disable_notification"] = strconv.FormatBool(opts.DisableNotification)
        params["protect_content"] = strconv.FormatBool(opts.ProtectContent)

        if opts.ReplyParameters != nil {
            bs, err := json.Marshal(opts.ReplyParameters)
            if err != nil {
                return nil, fmt.Errorf("failed to marshal field reply_parameters: %w", err)
            }
            params["reply_parameters"] = string(bs)
        }


        if opts.ReplyMarkup != nil {
            bs, err := json.Marshal(opts.ReplyMarkup)
            if err != nil {
                return nil, fmt.Errorf("failed to marshal field reply_markup: %w", err)
            }
            params["reply_markup"] = string(bs)
        }

    }


    r, err := b.Request("sendContact", params, data_params)
    if err != nil {
        return nil, err
    }

    
    var res *types.Message
    return res, json.Unmarshal(r, &res) 

}

// SendPoll methods's optional params
type SendPollOpts struct {
    MessageThreadId int64 `json:"message_thread_id,omitempty"`
    IsAnonymous bool `json:"is_anonymous,omitempty"`
    Type string `json:"type,omitempty"`
    AllowsMultipleAnswers bool `json:"allows_multiple_answers,omitempty"`
    CorrectOptionId int64 `json:"correct_option_id,omitempty"`
    Explanation string `json:"explanation,omitempty"`
    ExplanationParseMode string `json:"explanation_parse_mode,omitempty"`
    ExplanationEntities []types.MessageEntity `json:"explanation_entities,omitempty"`
    OpenPeriod int64 `json:"open_period,omitempty"`
    CloseDate int64 `json:"close_date,omitempty"`
    IsClosed bool `json:"is_closed,omitempty"`
    DisableNotification bool `json:"disable_notification,omitempty"`
    ProtectContent bool `json:"protect_content,omitempty"`
    ReplyParameters *types.ReplyParameters `json:"reply_parameters,omitempty"`
    ReplyMarkup types.ReplyMarkup `json:"reply_markup,omitempty"`
}

// Use this method to send a native poll. On success, the sent Message is returned.
func (b *Bot) SendPoll(chatId int64, question string, options []string, opts *SendPollOpts) (*types.Message, error) {
    params := map[string]string{}
    data_params := map[string]string{}

    params["chat_id"] = strconv.FormatInt(chatId, 10)
    params["question"] = question

    if options != nil {
        bs, err := json.Marshal(options)
        if err != nil {
            return nil, fmt.Errorf("failed to marshal field options: %w", err)
        }
        params["options"] = string(bs)
    }

    if opts != nil {
        params["message_thread_id"] = strconv.FormatInt(opts.MessageThreadId, 10)
        params["is_anonymous"] = strconv.FormatBool(opts.IsAnonymous)
        params["type"] = opts.Type
        params["allows_multiple_answers"] = strconv.FormatBool(opts.AllowsMultipleAnswers)
        params["correct_option_id"] = strconv.FormatInt(opts.CorrectOptionId, 10)
        params["explanation"] = opts.Explanation
        params["explanation_parse_mode"] = opts.ExplanationParseMode

        if opts.ExplanationEntities != nil {
            bs, err := json.Marshal(opts.ExplanationEntities)
            if err != nil {
                return nil, fmt.Errorf("failed to marshal field explanation_entities: %w", err)
            }
            params["explanation_entities"] = string(bs)
        }

        params["open_period"] = strconv.FormatInt(opts.OpenPeriod, 10)
        params["close_date"] = strconv.FormatInt(opts.CloseDate, 10)
        params["is_closed"] = strconv.FormatBool(opts.IsClosed)
        params["disable_notification"] = strconv.FormatBool(opts.DisableNotification)
        params["protect_content"] = strconv.FormatBool(opts.ProtectContent)

        if opts.ReplyParameters != nil {
            bs, err := json.Marshal(opts.ReplyParameters)
            if err != nil {
                return nil, fmt.Errorf("failed to marshal field reply_parameters: %w", err)
            }
            params["reply_parameters"] = string(bs)
        }


        if opts.ReplyMarkup != nil {
            bs, err := json.Marshal(opts.ReplyMarkup)
            if err != nil {
                return nil, fmt.Errorf("failed to marshal field reply_markup: %w", err)
            }
            params["reply_markup"] = string(bs)
        }

    }


    r, err := b.Request("sendPoll", params, data_params)
    if err != nil {
        return nil, err
    }

    
    var res *types.Message
    return res, json.Unmarshal(r, &res) 

}

// SendDice methods's optional params
type SendDiceOpts struct {
    MessageThreadId int64 `json:"message_thread_id,omitempty"`
    Emoji string `json:"emoji,omitempty"`
    DisableNotification bool `json:"disable_notification,omitempty"`
    ProtectContent bool `json:"protect_content,omitempty"`
    ReplyParameters *types.ReplyParameters `json:"reply_parameters,omitempty"`
    ReplyMarkup types.ReplyMarkup `json:"reply_markup,omitempty"`
}

// Use this method to send an animated emoji that will display a random value. On success, the sent Message is returned.
func (b *Bot) SendDice(chatId int64, opts *SendDiceOpts) (*types.Message, error) {
    params := map[string]string{}
    data_params := map[string]string{}

    params["chat_id"] = strconv.FormatInt(chatId, 10)
    if opts != nil {
        params["message_thread_id"] = strconv.FormatInt(opts.MessageThreadId, 10)
        params["emoji"] = opts.Emoji
        params["disable_notification"] = strconv.FormatBool(opts.DisableNotification)
        params["protect_content"] = strconv.FormatBool(opts.ProtectContent)

        if opts.ReplyParameters != nil {
            bs, err := json.Marshal(opts.ReplyParameters)
            if err != nil {
                return nil, fmt.Errorf("failed to marshal field reply_parameters: %w", err)
            }
            params["reply_parameters"] = string(bs)
        }


        if opts.ReplyMarkup != nil {
            bs, err := json.Marshal(opts.ReplyMarkup)
            if err != nil {
                return nil, fmt.Errorf("failed to marshal field reply_markup: %w", err)
            }
            params["reply_markup"] = string(bs)
        }

    }


    r, err := b.Request("sendDice", params, data_params)
    if err != nil {
        return nil, err
    }

    
    var res *types.Message
    return res, json.Unmarshal(r, &res) 

}

// SendChatAction methods's optional params
type SendChatActionOpts struct {
    MessageThreadId int64 `json:"message_thread_id,omitempty"`
}

// Use this method when you need to tell the user that something is happening on the bot's side. The status is set for 5 seconds or less (when a message arrives from your bot, Telegram clients clear its typing status). Returns True on success.
// We only recommend using this method when a response from the bot will take a noticeable amount of time to arrive.
func (b *Bot) SendChatAction(chatId int64, action string, opts *SendChatActionOpts) (bool, error) {
    params := map[string]string{}
    data_params := map[string]string{}

    params["chat_id"] = strconv.FormatInt(chatId, 10)
    params["action"] = action
    if opts != nil {
        params["message_thread_id"] = strconv.FormatInt(opts.MessageThreadId, 10)
    }


    r, err := b.Request("sendChatAction", params, data_params)
    if err != nil {
        return false, err
    }

    
    var res bool
    return res, json.Unmarshal(r, &res) 

}

// SetMessageReaction methods's optional params
type SetMessageReactionOpts struct {
    Reaction []types.ReactionType `json:"reaction,omitempty"`
    IsBig bool `json:"is_big,omitempty"`
}

// Use this method to change the chosen reactions on a message. Service messages can't be reacted to. Automatically forwarded messages from a channel to its discussion group have the same available reactions as messages in the channel. Returns True on success.
func (b *Bot) SetMessageReaction(chatId int64, messageId int64, opts *SetMessageReactionOpts) (bool, error) {
    params := map[string]string{}
    data_params := map[string]string{}

    params["chat_id"] = strconv.FormatInt(chatId, 10)
    params["message_id"] = strconv.FormatInt(messageId, 10)
    if opts != nil {

        if opts.Reaction != nil {
            bs, err := json.Marshal(opts.Reaction)
            if err != nil {
                return false, fmt.Errorf("failed to marshal field reaction: %w", err)
            }
            params["reaction"] = string(bs)
        }

        params["is_big"] = strconv.FormatBool(opts.IsBig)
    }


    r, err := b.Request("setMessageReaction", params, data_params)
    if err != nil {
        return false, err
    }

    
    var res bool
    return res, json.Unmarshal(r, &res) 

}

// GetUserProfilePhotos methods's optional params
type GetUserProfilePhotosOpts struct {
    Offset int64 `json:"offset,omitempty"`
    Limit int64 `json:"limit,omitempty"`
}

// Use this method to get a list of profile pictures for a user. Returns a UserProfilePhotos object.
func (b *Bot) GetUserProfilePhotos(userId int64, opts *GetUserProfilePhotosOpts) (*types.UserProfilePhotos, error) {
    params := map[string]string{}
    data_params := map[string]string{}

    params["user_id"] = strconv.FormatInt(userId, 10)
    if opts != nil {
        params["offset"] = strconv.FormatInt(opts.Offset, 10)
        params["limit"] = strconv.FormatInt(opts.Limit, 10)
    }


    r, err := b.Request("getUserProfilePhotos", params, data_params)
    if err != nil {
        return nil, err
    }

    
    var res *types.UserProfilePhotos
    return res, json.Unmarshal(r, &res) 

}

// Use this method to get basic information about a file and prepare it for downloading. For the moment, bots can download files of up to 20MB in size. On success, a File object is returned. The file can then be downloaded via the link https://api.telegram.org/file/bot<token>/<file_path>, where <file_path> is taken from the response. It is guaranteed that the link will be valid for at least 1 hour. When the link expires, a new one can be requested by calling getFile again.
// Note: This function may not preserve the original file name and MIME type. You should save the file's MIME type and name (if available) when the File object is received.
func (b *Bot) GetFile(fileId string) (*types.File, error) {
    params := map[string]string{}
    data_params := map[string]string{}
    params["file_id"] = fileId

    r, err := b.Request("getFile", params, data_params)
    if err != nil {
        return nil, err
    }
    
    var res *types.File
    return res, json.Unmarshal(r, &res) 

}

// BanChatMember methods's optional params
type BanChatMemberOpts struct {
    UntilDate int64 `json:"until_date,omitempty"`
    RevokeMessages bool `json:"revoke_messages,omitempty"`
}

// Use this method to ban a user in a group, a supergroup or a channel. In the case of supergroups and channels, the user will not be able to return to the chat on their own using invite links, etc., unless unbanned first. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns True on success.
func (b *Bot) BanChatMember(chatId int64, userId int64, opts *BanChatMemberOpts) (bool, error) {
    params := map[string]string{}
    data_params := map[string]string{}

    params["chat_id"] = strconv.FormatInt(chatId, 10)
    params["user_id"] = strconv.FormatInt(userId, 10)
    if opts != nil {
        params["until_date"] = strconv.FormatInt(opts.UntilDate, 10)
        params["revoke_messages"] = strconv.FormatBool(opts.RevokeMessages)
    }


    r, err := b.Request("banChatMember", params, data_params)
    if err != nil {
        return false, err
    }

    
    var res bool
    return res, json.Unmarshal(r, &res) 

}

// UnbanChatMember methods's optional params
type UnbanChatMemberOpts struct {
    OnlyIfBanned bool `json:"only_if_banned,omitempty"`
}

// Use this method to unban a previously banned user in a supergroup or channel. The user will not return to the group or channel automatically, but will be able to join via link, etc. The bot must be an administrator for this to work. By default, this method guarantees that after the call the user is not a member of the chat, but will be able to join it. So if the user is a member of the chat they will also be removed from the chat. If you don't want this, use the parameter only_if_banned. Returns True on success.
func (b *Bot) UnbanChatMember(chatId int64, userId int64, opts *UnbanChatMemberOpts) (bool, error) {
    params := map[string]string{}
    data_params := map[string]string{}

    params["chat_id"] = strconv.FormatInt(chatId, 10)
    params["user_id"] = strconv.FormatInt(userId, 10)
    if opts != nil {
        params["only_if_banned"] = strconv.FormatBool(opts.OnlyIfBanned)
    }


    r, err := b.Request("unbanChatMember", params, data_params)
    if err != nil {
        return false, err
    }

    
    var res bool
    return res, json.Unmarshal(r, &res) 

}

// RestrictChatMember methods's optional params
type RestrictChatMemberOpts struct {
    UseIndependentChatPermissions bool `json:"use_independent_chat_permissions,omitempty"`
    UntilDate int64 `json:"until_date,omitempty"`
}

// Use this method to restrict a user in a supergroup. The bot must be an administrator in the supergroup for this to work and must have the appropriate administrator rights. Pass True for all permissions to lift restrictions from a user. Returns True on success.
func (b *Bot) RestrictChatMember(chatId int64, userId int64, permissions *types.ChatPermissions, opts *RestrictChatMemberOpts) (bool, error) {
    params := map[string]string{}
    data_params := map[string]string{}

    params["chat_id"] = strconv.FormatInt(chatId, 10)
    params["user_id"] = strconv.FormatInt(userId, 10)

    if permissions != nil {
        bs, err := json.Marshal(permissions)
        if err != nil {
            return false, fmt.Errorf("failed to marshal field permissions: %w", err)
        }
        params["permissions"] = string(bs)
    }

    if opts != nil {
        params["use_independent_chat_permissions"] = strconv.FormatBool(opts.UseIndependentChatPermissions)
        params["until_date"] = strconv.FormatInt(opts.UntilDate, 10)
    }


    r, err := b.Request("restrictChatMember", params, data_params)
    if err != nil {
        return false, err
    }

    
    var res bool
    return res, json.Unmarshal(r, &res) 

}

// PromoteChatMember methods's optional params
type PromoteChatMemberOpts struct {
    IsAnonymous bool `json:"is_anonymous,omitempty"`
    CanManageChat bool `json:"can_manage_chat,omitempty"`
    CanDeleteMessages bool `json:"can_delete_messages,omitempty"`
    CanManageVideoChats bool `json:"can_manage_video_chats,omitempty"`
    CanRestrictMembers bool `json:"can_restrict_members,omitempty"`
    CanPromoteMembers bool `json:"can_promote_members,omitempty"`
    CanChangeInfo bool `json:"can_change_info,omitempty"`
    CanInviteUsers bool `json:"can_invite_users,omitempty"`
    CanPostStories bool `json:"can_post_stories,omitempty"`
    CanEditStories bool `json:"can_edit_stories,omitempty"`
    CanDeleteStories bool `json:"can_delete_stories,omitempty"`
    CanPostMessages bool `json:"can_post_messages,omitempty"`
    CanEditMessages bool `json:"can_edit_messages,omitempty"`
    CanPinMessages bool `json:"can_pin_messages,omitempty"`
    CanManageTopics bool `json:"can_manage_topics,omitempty"`
}

// Use this method to promote or demote a user in a supergroup or a channel. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Pass False for all boolean parameters to demote a user. Returns True on success.
func (b *Bot) PromoteChatMember(chatId int64, userId int64, opts *PromoteChatMemberOpts) (bool, error) {
    params := map[string]string{}
    data_params := map[string]string{}

    params["chat_id"] = strconv.FormatInt(chatId, 10)
    params["user_id"] = strconv.FormatInt(userId, 10)
    if opts != nil {
        params["is_anonymous"] = strconv.FormatBool(opts.IsAnonymous)
        params["can_manage_chat"] = strconv.FormatBool(opts.CanManageChat)
        params["can_delete_messages"] = strconv.FormatBool(opts.CanDeleteMessages)
        params["can_manage_video_chats"] = strconv.FormatBool(opts.CanManageVideoChats)
        params["can_restrict_members"] = strconv.FormatBool(opts.CanRestrictMembers)
        params["can_promote_members"] = strconv.FormatBool(opts.CanPromoteMembers)
        params["can_change_info"] = strconv.FormatBool(opts.CanChangeInfo)
        params["can_invite_users"] = strconv.FormatBool(opts.CanInviteUsers)
        params["can_post_stories"] = strconv.FormatBool(opts.CanPostStories)
        params["can_edit_stories"] = strconv.FormatBool(opts.CanEditStories)
        params["can_delete_stories"] = strconv.FormatBool(opts.CanDeleteStories)
        params["can_post_messages"] = strconv.FormatBool(opts.CanPostMessages)
        params["can_edit_messages"] = strconv.FormatBool(opts.CanEditMessages)
        params["can_pin_messages"] = strconv.FormatBool(opts.CanPinMessages)
        params["can_manage_topics"] = strconv.FormatBool(opts.CanManageTopics)
    }


    r, err := b.Request("promoteChatMember", params, data_params)
    if err != nil {
        return false, err
    }

    
    var res bool
    return res, json.Unmarshal(r, &res) 

}

// Use this method to set a custom title for an administrator in a supergroup promoted by the bot. Returns True on success.
func (b *Bot) SetChatAdministratorCustomTitle(chatId int64, userId int64, customTitle string) (bool, error) {
    params := map[string]string{}
    data_params := map[string]string{}
    params["chat_id"] = strconv.FormatInt(chatId, 10)
    params["user_id"] = strconv.FormatInt(userId, 10)
    params["custom_title"] = customTitle

    r, err := b.Request("setChatAdministratorCustomTitle", params, data_params)
    if err != nil {
        return false, err
    }
    
    var res bool
    return res, json.Unmarshal(r, &res) 

}

// Use this method to ban a channel chat in a supergroup or a channel. Until the chat is unbanned, the owner of the banned chat won't be able to send messages on behalf of any of their channels. The bot must be an administrator in the supergroup or channel for this to work and must have the appropriate administrator rights. Returns True on success.
func (b *Bot) BanChatSenderChat(chatId int64, senderChatId int64) (bool, error) {
    params := map[string]string{}
    data_params := map[string]string{}
    params["chat_id"] = strconv.FormatInt(chatId, 10)
    params["sender_chat_id"] = strconv.FormatInt(senderChatId, 10)

    r, err := b.Request("banChatSenderChat", params, data_params)
    if err != nil {
        return false, err
    }
    
    var res bool
    return res, json.Unmarshal(r, &res) 

}

// Use this method to unban a previously banned channel chat in a supergroup or channel. The bot must be an administrator for this to work and must have the appropriate administrator rights. Returns True on success.
func (b *Bot) UnbanChatSenderChat(chatId int64, senderChatId int64) (bool, error) {
    params := map[string]string{}
    data_params := map[string]string{}
    params["chat_id"] = strconv.FormatInt(chatId, 10)
    params["sender_chat_id"] = strconv.FormatInt(senderChatId, 10)

    r, err := b.Request("unbanChatSenderChat", params, data_params)
    if err != nil {
        return false, err
    }
    
    var res bool
    return res, json.Unmarshal(r, &res) 

}

// SetChatPermissions methods's optional params
type SetChatPermissionsOpts struct {
    UseIndependentChatPermissions bool `json:"use_independent_chat_permissions,omitempty"`
}

// Use this method to set default chat permissions for all members. The bot must be an administrator in the group or a supergroup for this to work and must have the can_restrict_members administrator rights. Returns True on success.
func (b *Bot) SetChatPermissions(chatId int64, permissions *types.ChatPermissions, opts *SetChatPermissionsOpts) (bool, error) {
    params := map[string]string{}
    data_params := map[string]string{}

    params["chat_id"] = strconv.FormatInt(chatId, 10)

    if permissions != nil {
        bs, err := json.Marshal(permissions)
        if err != nil {
            return false, fmt.Errorf("failed to marshal field permissions: %w", err)
        }
        params["permissions"] = string(bs)
    }

    if opts != nil {
        params["use_independent_chat_permissions"] = strconv.FormatBool(opts.UseIndependentChatPermissions)
    }


    r, err := b.Request("setChatPermissions", params, data_params)
    if err != nil {
        return false, err
    }

    
    var res bool
    return res, json.Unmarshal(r, &res) 

}

// Use this method to generate a new primary invite link for a chat; any previously generated primary link is revoked. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns the new invite link as String on success.
func (b *Bot) ExportChatInviteLink(chatId int64) (string, error) {
    params := map[string]string{}
    data_params := map[string]string{}
    params["chat_id"] = strconv.FormatInt(chatId, 10)

    r, err := b.Request("exportChatInviteLink", params, data_params)
    if err != nil {
        return "", err
    }
    
    var res string
    return res, json.Unmarshal(r, &res) 

}

// CreateChatInviteLink methods's optional params
type CreateChatInviteLinkOpts struct {
    Name string `json:"name,omitempty"`
    ExpireDate int64 `json:"expire_date,omitempty"`
    MemberLimit int64 `json:"member_limit,omitempty"`
    CreatesJoinRequest bool `json:"creates_join_request,omitempty"`
}

// Use this method to create an additional invite link for a chat. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. The link can be revoked using the method revokeChatInviteLink. Returns the new invite link as ChatInviteLink object.
func (b *Bot) CreateChatInviteLink(chatId int64, opts *CreateChatInviteLinkOpts) (*types.ChatInviteLink, error) {
    params := map[string]string{}
    data_params := map[string]string{}

    params["chat_id"] = strconv.FormatInt(chatId, 10)
    if opts != nil {
        params["name"] = opts.Name
        params["expire_date"] = strconv.FormatInt(opts.ExpireDate, 10)
        params["member_limit"] = strconv.FormatInt(opts.MemberLimit, 10)
        params["creates_join_request"] = strconv.FormatBool(opts.CreatesJoinRequest)
    }


    r, err := b.Request("createChatInviteLink", params, data_params)
    if err != nil {
        return nil, err
    }

    
    var res *types.ChatInviteLink
    return res, json.Unmarshal(r, &res) 

}

// EditChatInviteLink methods's optional params
type EditChatInviteLinkOpts struct {
    Name string `json:"name,omitempty"`
    ExpireDate int64 `json:"expire_date,omitempty"`
    MemberLimit int64 `json:"member_limit,omitempty"`
    CreatesJoinRequest bool `json:"creates_join_request,omitempty"`
}

// Use this method to edit a non-primary invite link created by the bot. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns the edited invite link as a ChatInviteLink object.
func (b *Bot) EditChatInviteLink(chatId int64, inviteLink string, opts *EditChatInviteLinkOpts) (*types.ChatInviteLink, error) {
    params := map[string]string{}
    data_params := map[string]string{}

    params["chat_id"] = strconv.FormatInt(chatId, 10)
    params["invite_link"] = inviteLink
    if opts != nil {
        params["name"] = opts.Name
        params["expire_date"] = strconv.FormatInt(opts.ExpireDate, 10)
        params["member_limit"] = strconv.FormatInt(opts.MemberLimit, 10)
        params["creates_join_request"] = strconv.FormatBool(opts.CreatesJoinRequest)
    }


    r, err := b.Request("editChatInviteLink", params, data_params)
    if err != nil {
        return nil, err
    }

    
    var res *types.ChatInviteLink
    return res, json.Unmarshal(r, &res) 

}

// Use this method to revoke an invite link created by the bot. If the primary link is revoked, a new link is automatically generated. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns the revoked invite link as ChatInviteLink object.
func (b *Bot) RevokeChatInviteLink(chatId int64, inviteLink string) (*types.ChatInviteLink, error) {
    params := map[string]string{}
    data_params := map[string]string{}
    params["chat_id"] = strconv.FormatInt(chatId, 10)
    params["invite_link"] = inviteLink

    r, err := b.Request("revokeChatInviteLink", params, data_params)
    if err != nil {
        return nil, err
    }
    
    var res *types.ChatInviteLink
    return res, json.Unmarshal(r, &res) 

}

// Use this method to approve a chat join request. The bot must be an administrator in the chat for this to work and must have the can_invite_users administrator right. Returns True on success.
func (b *Bot) ApproveChatJoinRequest(chatId int64, userId int64) (bool, error) {
    params := map[string]string{}
    data_params := map[string]string{}
    params["chat_id"] = strconv.FormatInt(chatId, 10)
    params["user_id"] = strconv.FormatInt(userId, 10)

    r, err := b.Request("approveChatJoinRequest", params, data_params)
    if err != nil {
        return false, err
    }
    
    var res bool
    return res, json.Unmarshal(r, &res) 

}

// Use this method to decline a chat join request. The bot must be an administrator in the chat for this to work and must have the can_invite_users administrator right. Returns True on success.
func (b *Bot) DeclineChatJoinRequest(chatId int64, userId int64) (bool, error) {
    params := map[string]string{}
    data_params := map[string]string{}
    params["chat_id"] = strconv.FormatInt(chatId, 10)
    params["user_id"] = strconv.FormatInt(userId, 10)

    r, err := b.Request("declineChatJoinRequest", params, data_params)
    if err != nil {
        return false, err
    }
    
    var res bool
    return res, json.Unmarshal(r, &res) 

}

// Use this method to set a new profile photo for the chat. Photos can't be changed for private chats. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns True on success.
func (b *Bot) SetChatPhoto(chatId int64, photo types.InputFile) (bool, error) {
    params := map[string]string{}
    data_params := map[string]string{}
    params["chat_id"] = strconv.FormatInt(chatId, 10)

    if photo != nil {
        switch f := photo.(type) {
        case string:
            _, err := os.Stat(f)
            if err != nil {
                params["photo"] = f
            } else {
                params["photo"] = "attach://photo"
                data_params["photo"] = f
            }
        default:
            return false, fmt.Errorf("unknown type for InputFile: %T", photo)
        }
    }

    r, err := b.Request("setChatPhoto", params, data_params)
    if err != nil {
        return false, err
    }
    
    var res bool
    return res, json.Unmarshal(r, &res) 

}

// Use this method to delete a chat photo. Photos can't be changed for private chats. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns True on success.
func (b *Bot) DeleteChatPhoto(chatId int64) (bool, error) {
    params := map[string]string{}
    data_params := map[string]string{}
    params["chat_id"] = strconv.FormatInt(chatId, 10)

    r, err := b.Request("deleteChatPhoto", params, data_params)
    if err != nil {
        return false, err
    }
    
    var res bool
    return res, json.Unmarshal(r, &res) 

}

// Use this method to change the title of a chat. Titles can't be changed for private chats. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns True on success.
func (b *Bot) SetChatTitle(chatId int64, title string) (bool, error) {
    params := map[string]string{}
    data_params := map[string]string{}
    params["chat_id"] = strconv.FormatInt(chatId, 10)
    params["title"] = title

    r, err := b.Request("setChatTitle", params, data_params)
    if err != nil {
        return false, err
    }
    
    var res bool
    return res, json.Unmarshal(r, &res) 

}

// SetChatDescription methods's optional params
type SetChatDescriptionOpts struct {
    Description string `json:"description,omitempty"`
}

// Use this method to change the description of a group, a supergroup or a channel. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns True on success.
func (b *Bot) SetChatDescription(chatId int64, opts *SetChatDescriptionOpts) (bool, error) {
    params := map[string]string{}
    data_params := map[string]string{}

    params["chat_id"] = strconv.FormatInt(chatId, 10)
    if opts != nil {
        params["description"] = opts.Description
    }


    r, err := b.Request("setChatDescription", params, data_params)
    if err != nil {
        return false, err
    }

    
    var res bool
    return res, json.Unmarshal(r, &res) 

}

// PinChatMessage methods's optional params
type PinChatMessageOpts struct {
    DisableNotification bool `json:"disable_notification,omitempty"`
}

// Use this method to add a message to the list of pinned messages in a chat. If the chat is not a private chat, the bot must be an administrator in the chat for this to work and must have the 'can_pin_messages' administrator right in a supergroup or 'can_edit_messages' administrator right in a channel. Returns True on success.
func (b *Bot) PinChatMessage(chatId int64, messageId int64, opts *PinChatMessageOpts) (bool, error) {
    params := map[string]string{}
    data_params := map[string]string{}

    params["chat_id"] = strconv.FormatInt(chatId, 10)
    params["message_id"] = strconv.FormatInt(messageId, 10)
    if opts != nil {
        params["disable_notification"] = strconv.FormatBool(opts.DisableNotification)
    }


    r, err := b.Request("pinChatMessage", params, data_params)
    if err != nil {
        return false, err
    }

    
    var res bool
    return res, json.Unmarshal(r, &res) 

}

// UnpinChatMessage methods's optional params
type UnpinChatMessageOpts struct {
    MessageId int64 `json:"message_id,omitempty"`
}

// Use this method to remove a message from the list of pinned messages in a chat. If the chat is not a private chat, the bot must be an administrator in the chat for this to work and must have the 'can_pin_messages' administrator right in a supergroup or 'can_edit_messages' administrator right in a channel. Returns True on success.
func (b *Bot) UnpinChatMessage(chatId int64, opts *UnpinChatMessageOpts) (bool, error) {
    params := map[string]string{}
    data_params := map[string]string{}

    params["chat_id"] = strconv.FormatInt(chatId, 10)
    if opts != nil {
        params["message_id"] = strconv.FormatInt(opts.MessageId, 10)
    }


    r, err := b.Request("unpinChatMessage", params, data_params)
    if err != nil {
        return false, err
    }

    
    var res bool
    return res, json.Unmarshal(r, &res) 

}

// Use this method to clear the list of pinned messages in a chat. If the chat is not a private chat, the bot must be an administrator in the chat for this to work and must have the 'can_pin_messages' administrator right in a supergroup or 'can_edit_messages' administrator right in a channel. Returns True on success.
func (b *Bot) UnpinAllChatMessages(chatId int64) (bool, error) {
    params := map[string]string{}
    data_params := map[string]string{}
    params["chat_id"] = strconv.FormatInt(chatId, 10)

    r, err := b.Request("unpinAllChatMessages", params, data_params)
    if err != nil {
        return false, err
    }
    
    var res bool
    return res, json.Unmarshal(r, &res) 

}

// Use this method for your bot to leave a group, supergroup or channel. Returns True on success.
func (b *Bot) LeaveChat(chatId int64) (bool, error) {
    params := map[string]string{}
    data_params := map[string]string{}
    params["chat_id"] = strconv.FormatInt(chatId, 10)

    r, err := b.Request("leaveChat", params, data_params)
    if err != nil {
        return false, err
    }
    
    var res bool
    return res, json.Unmarshal(r, &res) 

}

// Use this method to get up to date information about the chat. Returns a Chat object on success.
func (b *Bot) GetChat(chatId int64) (*types.Chat, error) {
    params := map[string]string{}
    data_params := map[string]string{}
    params["chat_id"] = strconv.FormatInt(chatId, 10)

    r, err := b.Request("getChat", params, data_params)
    if err != nil {
        return nil, err
    }
    
    var res *types.Chat
    return res, json.Unmarshal(r, &res) 

}

// Use this method to get a list of administrators in a chat, which aren't bots. Returns an Array of ChatMember objects.
func (b *Bot) GetChatAdministrators(chatId int64) ([]types.ChatMember, error) {
    params := map[string]string{}
    data_params := map[string]string{}
    params["chat_id"] = strconv.FormatInt(chatId, 10)

    r, err := b.Request("getChatAdministrators", params, data_params)
    if err != nil {
        return nil, err
    }
    
    var res []types.ChatMember
    return res, json.Unmarshal(r, &res) 

}

// Use this method to get the number of members in a chat. Returns Int on success.
func (b *Bot) GetChatMemberCount(chatId int64) (int64, error) {
    params := map[string]string{}
    data_params := map[string]string{}
    params["chat_id"] = strconv.FormatInt(chatId, 10)

    r, err := b.Request("getChatMemberCount", params, data_params)
    if err != nil {
        return 0, err
    }
    
    var res int64
    return res, json.Unmarshal(r, &res) 

}

// Use this method to get information about a member of a chat. The method is only guaranteed to work for other users if the bot is an administrator in the chat. Returns a ChatMember object on success.
func (b *Bot) GetChatMember(chatId int64, userId int64) (*types.ChatMember, error) {
    params := map[string]string{}
    data_params := map[string]string{}
    params["chat_id"] = strconv.FormatInt(chatId, 10)
    params["user_id"] = strconv.FormatInt(userId, 10)

    r, err := b.Request("getChatMember", params, data_params)
    if err != nil {
        return nil, err
    }
    
    var res *types.ChatMember
    return res, json.Unmarshal(r, &res) 

}

// Use this method to set a new group sticker set for a supergroup. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Use the field can_set_sticker_set optionally returned in getChat requests to check if the bot can use this method. Returns True on success.
func (b *Bot) SetChatStickerSet(chatId int64, stickerSetName string) (bool, error) {
    params := map[string]string{}
    data_params := map[string]string{}
    params["chat_id"] = strconv.FormatInt(chatId, 10)
    params["sticker_set_name"] = stickerSetName

    r, err := b.Request("setChatStickerSet", params, data_params)
    if err != nil {
        return false, err
    }
    
    var res bool
    return res, json.Unmarshal(r, &res) 

}

// Use this method to delete a group sticker set from a supergroup. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Use the field can_set_sticker_set optionally returned in getChat requests to check if the bot can use this method. Returns True on success.
func (b *Bot) DeleteChatStickerSet(chatId int64) (bool, error) {
    params := map[string]string{}
    data_params := map[string]string{}
    params["chat_id"] = strconv.FormatInt(chatId, 10)

    r, err := b.Request("deleteChatStickerSet", params, data_params)
    if err != nil {
        return false, err
    }
    
    var res bool
    return res, json.Unmarshal(r, &res) 

}

// Use this method to get custom emoji stickers, which can be used as a forum topic icon by any user. Requires no parameters. Returns an Array of Sticker objects.
func (b *Bot) GetForumTopicIconStickers() ([]types.Sticker, error) {
    params := map[string]string{}
    data_params := map[string]string{}

    r, err := b.Request("getForumTopicIconStickers", params, data_params)
    if err != nil {
        return nil, err
    }
    
    var res []types.Sticker
    return res, json.Unmarshal(r, &res) 

}

// CreateForumTopic methods's optional params
type CreateForumTopicOpts struct {
    IconColor int64 `json:"icon_color,omitempty"`
    IconCustomEmojiId string `json:"icon_custom_emoji_id,omitempty"`
}

// Use this method to create a topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights. Returns information about the created topic as a ForumTopic object.
func (b *Bot) CreateForumTopic(chatId int64, name string, opts *CreateForumTopicOpts) (*types.ForumTopic, error) {
    params := map[string]string{}
    data_params := map[string]string{}

    params["chat_id"] = strconv.FormatInt(chatId, 10)
    params["name"] = name
    if opts != nil {
        params["icon_color"] = strconv.FormatInt(opts.IconColor, 10)
        params["icon_custom_emoji_id"] = opts.IconCustomEmojiId
    }


    r, err := b.Request("createForumTopic", params, data_params)
    if err != nil {
        return nil, err
    }

    
    var res *types.ForumTopic
    return res, json.Unmarshal(r, &res) 

}

// EditForumTopic methods's optional params
type EditForumTopicOpts struct {
    Name string `json:"name,omitempty"`
    IconCustomEmojiId string `json:"icon_custom_emoji_id,omitempty"`
}

// Use this method to edit name and icon of a topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have can_manage_topics administrator rights, unless it is the creator of the topic. Returns True on success.
func (b *Bot) EditForumTopic(chatId int64, messageThreadId int64, opts *EditForumTopicOpts) (bool, error) {
    params := map[string]string{}
    data_params := map[string]string{}

    params["chat_id"] = strconv.FormatInt(chatId, 10)
    params["message_thread_id"] = strconv.FormatInt(messageThreadId, 10)
    if opts != nil {
        params["name"] = opts.Name
        params["icon_custom_emoji_id"] = opts.IconCustomEmojiId
    }


    r, err := b.Request("editForumTopic", params, data_params)
    if err != nil {
        return false, err
    }

    
    var res bool
    return res, json.Unmarshal(r, &res) 

}

// Use this method to close an open topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights, unless it is the creator of the topic. Returns True on success.
func (b *Bot) CloseForumTopic(chatId int64, messageThreadId int64) (bool, error) {
    params := map[string]string{}
    data_params := map[string]string{}
    params["chat_id"] = strconv.FormatInt(chatId, 10)
    params["message_thread_id"] = strconv.FormatInt(messageThreadId, 10)

    r, err := b.Request("closeForumTopic", params, data_params)
    if err != nil {
        return false, err
    }
    
    var res bool
    return res, json.Unmarshal(r, &res) 

}

// Use this method to reopen a closed topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights, unless it is the creator of the topic. Returns True on success.
func (b *Bot) ReopenForumTopic(chatId int64, messageThreadId int64) (bool, error) {
    params := map[string]string{}
    data_params := map[string]string{}
    params["chat_id"] = strconv.FormatInt(chatId, 10)
    params["message_thread_id"] = strconv.FormatInt(messageThreadId, 10)

    r, err := b.Request("reopenForumTopic", params, data_params)
    if err != nil {
        return false, err
    }
    
    var res bool
    return res, json.Unmarshal(r, &res) 

}

// Use this method to delete a forum topic along with all its messages in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_delete_messages administrator rights. Returns True on success.
func (b *Bot) DeleteForumTopic(chatId int64, messageThreadId int64) (bool, error) {
    params := map[string]string{}
    data_params := map[string]string{}
    params["chat_id"] = strconv.FormatInt(chatId, 10)
    params["message_thread_id"] = strconv.FormatInt(messageThreadId, 10)

    r, err := b.Request("deleteForumTopic", params, data_params)
    if err != nil {
        return false, err
    }
    
    var res bool
    return res, json.Unmarshal(r, &res) 

}

// Use this method to clear the list of pinned messages in a forum topic. The bot must be an administrator in the chat for this to work and must have the can_pin_messages administrator right in the supergroup. Returns True on success.
func (b *Bot) UnpinAllForumTopicMessages(chatId int64, messageThreadId int64) (bool, error) {
    params := map[string]string{}
    data_params := map[string]string{}
    params["chat_id"] = strconv.FormatInt(chatId, 10)
    params["message_thread_id"] = strconv.FormatInt(messageThreadId, 10)

    r, err := b.Request("unpinAllForumTopicMessages", params, data_params)
    if err != nil {
        return false, err
    }
    
    var res bool
    return res, json.Unmarshal(r, &res) 

}

// Use this method to edit the name of the 'General' topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have can_manage_topics administrator rights. Returns True on success.
func (b *Bot) EditGeneralForumTopic(chatId int64, name string) (bool, error) {
    params := map[string]string{}
    data_params := map[string]string{}
    params["chat_id"] = strconv.FormatInt(chatId, 10)
    params["name"] = name

    r, err := b.Request("editGeneralForumTopic", params, data_params)
    if err != nil {
        return false, err
    }
    
    var res bool
    return res, json.Unmarshal(r, &res) 

}

// Use this method to close an open 'General' topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights. Returns True on success.
func (b *Bot) CloseGeneralForumTopic(chatId int64) (bool, error) {
    params := map[string]string{}
    data_params := map[string]string{}
    params["chat_id"] = strconv.FormatInt(chatId, 10)

    r, err := b.Request("closeGeneralForumTopic", params, data_params)
    if err != nil {
        return false, err
    }
    
    var res bool
    return res, json.Unmarshal(r, &res) 

}

// Use this method to reopen a closed 'General' topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights. The topic will be automatically unhidden if it was hidden. Returns True on success.
func (b *Bot) ReopenGeneralForumTopic(chatId int64) (bool, error) {
    params := map[string]string{}
    data_params := map[string]string{}
    params["chat_id"] = strconv.FormatInt(chatId, 10)

    r, err := b.Request("reopenGeneralForumTopic", params, data_params)
    if err != nil {
        return false, err
    }
    
    var res bool
    return res, json.Unmarshal(r, &res) 

}

// Use this method to hide the 'General' topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights. The topic will be automatically closed if it was open. Returns True on success.
func (b *Bot) HideGeneralForumTopic(chatId int64) (bool, error) {
    params := map[string]string{}
    data_params := map[string]string{}
    params["chat_id"] = strconv.FormatInt(chatId, 10)

    r, err := b.Request("hideGeneralForumTopic", params, data_params)
    if err != nil {
        return false, err
    }
    
    var res bool
    return res, json.Unmarshal(r, &res) 

}

// Use this method to unhide the 'General' topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights. Returns True on success.
func (b *Bot) UnhideGeneralForumTopic(chatId int64) (bool, error) {
    params := map[string]string{}
    data_params := map[string]string{}
    params["chat_id"] = strconv.FormatInt(chatId, 10)

    r, err := b.Request("unhideGeneralForumTopic", params, data_params)
    if err != nil {
        return false, err
    }
    
    var res bool
    return res, json.Unmarshal(r, &res) 

}

// Use this method to clear the list of pinned messages in a General forum topic. The bot must be an administrator in the chat for this to work and must have the can_pin_messages administrator right in the supergroup. Returns True on success.
func (b *Bot) UnpinAllGeneralForumTopicMessages(chatId int64) (bool, error) {
    params := map[string]string{}
    data_params := map[string]string{}
    params["chat_id"] = strconv.FormatInt(chatId, 10)

    r, err := b.Request("unpinAllGeneralForumTopicMessages", params, data_params)
    if err != nil {
        return false, err
    }
    
    var res bool
    return res, json.Unmarshal(r, &res) 

}

// AnswerCallbackQuery methods's optional params
type AnswerCallbackQueryOpts struct {
    Text string `json:"text,omitempty"`
    ShowAlert bool `json:"show_alert,omitempty"`
    Url string `json:"url,omitempty"`
    CacheTime int64 `json:"cache_time,omitempty"`
}

// Use this method to send answers to callback queries sent from inline keyboards. The answer will be displayed to the user as a notification at the top of the chat screen or as an alert. On success, True is returned.
func (b *Bot) AnswerCallbackQuery(callbackQueryId string, opts *AnswerCallbackQueryOpts) (bool, error) {
    params := map[string]string{}
    data_params := map[string]string{}

    params["callback_query_id"] = callbackQueryId
    if opts != nil {
        params["text"] = opts.Text
        params["show_alert"] = strconv.FormatBool(opts.ShowAlert)
        params["url"] = opts.Url
        params["cache_time"] = strconv.FormatInt(opts.CacheTime, 10)
    }


    r, err := b.Request("answerCallbackQuery", params, data_params)
    if err != nil {
        return false, err
    }

    
    var res bool
    return res, json.Unmarshal(r, &res) 

}

// Use this method to get the list of boosts added to a chat by a user. Requires administrator rights in the chat. Returns a UserChatBoosts object.
func (b *Bot) GetUserChatBoosts(chatId int64, userId int64) (*types.UserChatBoosts, error) {
    params := map[string]string{}
    data_params := map[string]string{}
    params["chat_id"] = strconv.FormatInt(chatId, 10)
    params["user_id"] = strconv.FormatInt(userId, 10)

    r, err := b.Request("getUserChatBoosts", params, data_params)
    if err != nil {
        return nil, err
    }
    
    var res *types.UserChatBoosts
    return res, json.Unmarshal(r, &res) 

}

// SetMyCommands methods's optional params
type SetMyCommandsOpts struct {
    Scope *types.BotCommandScope `json:"scope,omitempty"`
    LanguageCode string `json:"language_code,omitempty"`
}

// Use this method to change the list of the bot's commands. See this manual for more details about bot commands. Returns True on success.
func (b *Bot) SetMyCommands(commands []types.BotCommand, opts *SetMyCommandsOpts) (bool, error) {
    params := map[string]string{}
    data_params := map[string]string{}


    if commands != nil {
        bs, err := json.Marshal(commands)
        if err != nil {
            return false, fmt.Errorf("failed to marshal field commands: %w", err)
        }
        params["commands"] = string(bs)
    }

    if opts != nil {

        if opts.Scope != nil {
            bs, err := json.Marshal(opts.Scope)
            if err != nil {
                return false, fmt.Errorf("failed to marshal field scope: %w", err)
            }
            params["scope"] = string(bs)
        }

        params["language_code"] = opts.LanguageCode
    }


    r, err := b.Request("setMyCommands", params, data_params)
    if err != nil {
        return false, err
    }

    
    var res bool
    return res, json.Unmarshal(r, &res) 

}

// DeleteMyCommands methods's optional params
type DeleteMyCommandsOpts struct {
    Scope *types.BotCommandScope `json:"scope,omitempty"`
    LanguageCode string `json:"language_code,omitempty"`
}

// Use this method to delete the list of the bot's commands for the given scope and user language. After deletion, higher level commands will be shown to affected users. Returns True on success.
func (b *Bot) DeleteMyCommands(opts *DeleteMyCommandsOpts) (bool, error) {
    params := map[string]string{}
    data_params := map[string]string{}

    if opts != nil {

        if opts.Scope != nil {
            bs, err := json.Marshal(opts.Scope)
            if err != nil {
                return false, fmt.Errorf("failed to marshal field scope: %w", err)
            }
            params["scope"] = string(bs)
        }

        params["language_code"] = opts.LanguageCode
    }


    r, err := b.Request("deleteMyCommands", params, data_params)
    if err != nil {
        return false, err
    }

    
    var res bool
    return res, json.Unmarshal(r, &res) 

}

// GetMyCommands methods's optional params
type GetMyCommandsOpts struct {
    Scope *types.BotCommandScope `json:"scope,omitempty"`
    LanguageCode string `json:"language_code,omitempty"`
}

// Use this method to get the current list of the bot's commands for the given scope and user language. Returns an Array of BotCommand objects. If commands aren't set, an empty list is returned.
func (b *Bot) GetMyCommands(opts *GetMyCommandsOpts) ([]types.BotCommand, error) {
    params := map[string]string{}
    data_params := map[string]string{}

    if opts != nil {

        if opts.Scope != nil {
            bs, err := json.Marshal(opts.Scope)
            if err != nil {
                return nil, fmt.Errorf("failed to marshal field scope: %w", err)
            }
            params["scope"] = string(bs)
        }

        params["language_code"] = opts.LanguageCode
    }


    r, err := b.Request("getMyCommands", params, data_params)
    if err != nil {
        return nil, err
    }

    
    var res []types.BotCommand
    return res, json.Unmarshal(r, &res) 

}

// SetMyName methods's optional params
type SetMyNameOpts struct {
    Name string `json:"name,omitempty"`
    LanguageCode string `json:"language_code,omitempty"`
}

// Use this method to change the bot's name. Returns True on success.
func (b *Bot) SetMyName(opts *SetMyNameOpts) (bool, error) {
    params := map[string]string{}
    data_params := map[string]string{}

    if opts != nil {
        params["name"] = opts.Name
        params["language_code"] = opts.LanguageCode
    }


    r, err := b.Request("setMyName", params, data_params)
    if err != nil {
        return false, err
    }

    
    var res bool
    return res, json.Unmarshal(r, &res) 

}

// GetMyName methods's optional params
type GetMyNameOpts struct {
    LanguageCode string `json:"language_code,omitempty"`
}

// Use this method to get the current bot name for the given user language. Returns BotName on success.
func (b *Bot) GetMyName(opts *GetMyNameOpts) (*types.BotName, error) {
    params := map[string]string{}
    data_params := map[string]string{}

    if opts != nil {
        params["language_code"] = opts.LanguageCode
    }


    r, err := b.Request("getMyName", params, data_params)
    if err != nil {
        return nil, err
    }

    
    var res *types.BotName
    return res, json.Unmarshal(r, &res) 

}

// SetMyDescription methods's optional params
type SetMyDescriptionOpts struct {
    Description string `json:"description,omitempty"`
    LanguageCode string `json:"language_code,omitempty"`
}

// Use this method to change the bot's description, which is shown in the chat with the bot if the chat is empty. Returns True on success.
func (b *Bot) SetMyDescription(opts *SetMyDescriptionOpts) (bool, error) {
    params := map[string]string{}
    data_params := map[string]string{}

    if opts != nil {
        params["description"] = opts.Description
        params["language_code"] = opts.LanguageCode
    }


    r, err := b.Request("setMyDescription", params, data_params)
    if err != nil {
        return false, err
    }

    
    var res bool
    return res, json.Unmarshal(r, &res) 

}

// GetMyDescription methods's optional params
type GetMyDescriptionOpts struct {
    LanguageCode string `json:"language_code,omitempty"`
}

// Use this method to get the current bot description for the given user language. Returns BotDescription on success.
func (b *Bot) GetMyDescription(opts *GetMyDescriptionOpts) (*types.BotDescription, error) {
    params := map[string]string{}
    data_params := map[string]string{}

    if opts != nil {
        params["language_code"] = opts.LanguageCode
    }


    r, err := b.Request("getMyDescription", params, data_params)
    if err != nil {
        return nil, err
    }

    
    var res *types.BotDescription
    return res, json.Unmarshal(r, &res) 

}

// SetMyShortDescription methods's optional params
type SetMyShortDescriptionOpts struct {
    ShortDescription string `json:"short_description,omitempty"`
    LanguageCode string `json:"language_code,omitempty"`
}

// Use this method to change the bot's short description, which is shown on the bot's profile page and is sent together with the link when users share the bot. Returns True on success.
func (b *Bot) SetMyShortDescription(opts *SetMyShortDescriptionOpts) (bool, error) {
    params := map[string]string{}
    data_params := map[string]string{}

    if opts != nil {
        params["short_description"] = opts.ShortDescription
        params["language_code"] = opts.LanguageCode
    }


    r, err := b.Request("setMyShortDescription", params, data_params)
    if err != nil {
        return false, err
    }

    
    var res bool
    return res, json.Unmarshal(r, &res) 

}

// GetMyShortDescription methods's optional params
type GetMyShortDescriptionOpts struct {
    LanguageCode string `json:"language_code,omitempty"`
}

// Use this method to get the current bot short description for the given user language. Returns BotShortDescription on success.
func (b *Bot) GetMyShortDescription(opts *GetMyShortDescriptionOpts) (*types.BotShortDescription, error) {
    params := map[string]string{}
    data_params := map[string]string{}

    if opts != nil {
        params["language_code"] = opts.LanguageCode
    }


    r, err := b.Request("getMyShortDescription", params, data_params)
    if err != nil {
        return nil, err
    }

    
    var res *types.BotShortDescription
    return res, json.Unmarshal(r, &res) 

}

// SetChatMenuButton methods's optional params
type SetChatMenuButtonOpts struct {
    ChatId int64 `json:"chat_id,omitempty"`
    MenuButton *types.MenuButton `json:"menu_button,omitempty"`
}

// Use this method to change the bot's menu button in a private chat, or the default menu button. Returns True on success.
func (b *Bot) SetChatMenuButton(opts *SetChatMenuButtonOpts) (bool, error) {
    params := map[string]string{}
    data_params := map[string]string{}

    if opts != nil {
        params["chat_id"] = strconv.FormatInt(opts.ChatId, 10)

        if opts.MenuButton != nil {
            bs, err := json.Marshal(opts.MenuButton)
            if err != nil {
                return false, fmt.Errorf("failed to marshal field menu_button: %w", err)
            }
            params["menu_button"] = string(bs)
        }

    }


    r, err := b.Request("setChatMenuButton", params, data_params)
    if err != nil {
        return false, err
    }

    
    var res bool
    return res, json.Unmarshal(r, &res) 

}

// GetChatMenuButton methods's optional params
type GetChatMenuButtonOpts struct {
    ChatId int64 `json:"chat_id,omitempty"`
}

// Use this method to get the current value of the bot's menu button in a private chat, or the default menu button. Returns MenuButton on success.
func (b *Bot) GetChatMenuButton(opts *GetChatMenuButtonOpts) (*types.MenuButton, error) {
    params := map[string]string{}
    data_params := map[string]string{}

    if opts != nil {
        params["chat_id"] = strconv.FormatInt(opts.ChatId, 10)
    }


    r, err := b.Request("getChatMenuButton", params, data_params)
    if err != nil {
        return nil, err
    }

    
    var res *types.MenuButton
    return res, json.Unmarshal(r, &res) 

}

// SetMyDefaultAdministratorRights methods's optional params
type SetMyDefaultAdministratorRightsOpts struct {
    Rights *types.ChatAdministratorRights `json:"rights,omitempty"`
    ForChannels bool `json:"for_channels,omitempty"`
}

// Use this method to change the default administrator rights requested by the bot when it's added as an administrator to groups or channels. These rights will be suggested to users, but they are free to modify the list before adding the bot. Returns True on success.
func (b *Bot) SetMyDefaultAdministratorRights(opts *SetMyDefaultAdministratorRightsOpts) (bool, error) {
    params := map[string]string{}
    data_params := map[string]string{}

    if opts != nil {

        if opts.Rights != nil {
            bs, err := json.Marshal(opts.Rights)
            if err != nil {
                return false, fmt.Errorf("failed to marshal field rights: %w", err)
            }
            params["rights"] = string(bs)
        }

        params["for_channels"] = strconv.FormatBool(opts.ForChannels)
    }


    r, err := b.Request("setMyDefaultAdministratorRights", params, data_params)
    if err != nil {
        return false, err
    }

    
    var res bool
    return res, json.Unmarshal(r, &res) 

}

// GetMyDefaultAdministratorRights methods's optional params
type GetMyDefaultAdministratorRightsOpts struct {
    ForChannels bool `json:"for_channels,omitempty"`
}

// Use this method to get the current default administrator rights of the bot. Returns ChatAdministratorRights on success.
func (b *Bot) GetMyDefaultAdministratorRights(opts *GetMyDefaultAdministratorRightsOpts) (*types.ChatAdministratorRights, error) {
    params := map[string]string{}
    data_params := map[string]string{}

    if opts != nil {
        params["for_channels"] = strconv.FormatBool(opts.ForChannels)
    }


    r, err := b.Request("getMyDefaultAdministratorRights", params, data_params)
    if err != nil {
        return nil, err
    }

    
    var res *types.ChatAdministratorRights
    return res, json.Unmarshal(r, &res) 

}

// EditMessageText methods's optional params
type EditMessageTextOpts struct {
    ChatId int64 `json:"chat_id,omitempty"`
    MessageId int64 `json:"message_id,omitempty"`
    InlineMessageId string `json:"inline_message_id,omitempty"`
    ParseMode string `json:"parse_mode,omitempty"`
    Entities []types.MessageEntity `json:"entities,omitempty"`
    LinkPreviewOptions *types.LinkPreviewOptions `json:"link_preview_options,omitempty"`
    ReplyMarkup *types.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// Use this method to edit text and game messages. On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned.
func (b *Bot) EditMessageText(text string, opts *EditMessageTextOpts) (*types.Message, error) {
    params := map[string]string{}
    data_params := map[string]string{}

    params["text"] = text
    if opts != nil {
        params["chat_id"] = strconv.FormatInt(opts.ChatId, 10)
        params["message_id"] = strconv.FormatInt(opts.MessageId, 10)
        params["inline_message_id"] = opts.InlineMessageId
        params["parse_mode"] = opts.ParseMode

        if opts.Entities != nil {
            bs, err := json.Marshal(opts.Entities)
            if err != nil {
                return nil, fmt.Errorf("failed to marshal field entities: %w", err)
            }
            params["entities"] = string(bs)
        }


        if opts.LinkPreviewOptions != nil {
            bs, err := json.Marshal(opts.LinkPreviewOptions)
            if err != nil {
                return nil, fmt.Errorf("failed to marshal field link_preview_options: %w", err)
            }
            params["link_preview_options"] = string(bs)
        }


        if opts.ReplyMarkup != nil {
            bs, err := json.Marshal(opts.ReplyMarkup)
            if err != nil {
                return nil, fmt.Errorf("failed to marshal field reply_markup: %w", err)
            }
            params["reply_markup"] = string(bs)
        }

    }


    r, err := b.Request("editMessageText", params, data_params)
    if err != nil {
        return nil, err
    }

    
    var res *types.Message
    return res, json.Unmarshal(r, &res) 

}

// EditMessageCaption methods's optional params
type EditMessageCaptionOpts struct {
    ChatId int64 `json:"chat_id,omitempty"`
    MessageId int64 `json:"message_id,omitempty"`
    InlineMessageId string `json:"inline_message_id,omitempty"`
    Caption string `json:"caption,omitempty"`
    ParseMode string `json:"parse_mode,omitempty"`
    CaptionEntities []types.MessageEntity `json:"caption_entities,omitempty"`
    ReplyMarkup *types.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// Use this method to edit captions of messages. On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned.
func (b *Bot) EditMessageCaption(opts *EditMessageCaptionOpts) (*types.Message, error) {
    params := map[string]string{}
    data_params := map[string]string{}

    if opts != nil {
        params["chat_id"] = strconv.FormatInt(opts.ChatId, 10)
        params["message_id"] = strconv.FormatInt(opts.MessageId, 10)
        params["inline_message_id"] = opts.InlineMessageId
        params["caption"] = opts.Caption
        params["parse_mode"] = opts.ParseMode

        if opts.CaptionEntities != nil {
            bs, err := json.Marshal(opts.CaptionEntities)
            if err != nil {
                return nil, fmt.Errorf("failed to marshal field caption_entities: %w", err)
            }
            params["caption_entities"] = string(bs)
        }


        if opts.ReplyMarkup != nil {
            bs, err := json.Marshal(opts.ReplyMarkup)
            if err != nil {
                return nil, fmt.Errorf("failed to marshal field reply_markup: %w", err)
            }
            params["reply_markup"] = string(bs)
        }

    }


    r, err := b.Request("editMessageCaption", params, data_params)
    if err != nil {
        return nil, err
    }

    
    var res *types.Message
    return res, json.Unmarshal(r, &res) 

}

// EditMessageMedia methods's optional params
type EditMessageMediaOpts struct {
    ChatId int64 `json:"chat_id,omitempty"`
    MessageId int64 `json:"message_id,omitempty"`
    InlineMessageId string `json:"inline_message_id,omitempty"`
    ReplyMarkup *types.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// Use this method to edit animation, audio, document, photo, or video messages. If a message is part of a message album, then it can be edited only to an audio for audio albums, only to a document for document albums and to a photo or a video otherwise. When an inline message is edited, a new file can't be uploaded; use a previously uploaded file via its file_id or specify a URL. On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned.
func (b *Bot) EditMessageMedia(media *types.InputMedia, opts *EditMessageMediaOpts) (*types.Message, error) {
    params := map[string]string{}
    data_params := map[string]string{}


    if media != nil {
        bs, err := json.Marshal(media)
        if err != nil {
            return nil, fmt.Errorf("failed to marshal field media: %w", err)
        }
        params["media"] = string(bs)
    }

    if opts != nil {
        params["chat_id"] = strconv.FormatInt(opts.ChatId, 10)
        params["message_id"] = strconv.FormatInt(opts.MessageId, 10)
        params["inline_message_id"] = opts.InlineMessageId

        if opts.ReplyMarkup != nil {
            bs, err := json.Marshal(opts.ReplyMarkup)
            if err != nil {
                return nil, fmt.Errorf("failed to marshal field reply_markup: %w", err)
            }
            params["reply_markup"] = string(bs)
        }

    }


    r, err := b.Request("editMessageMedia", params, data_params)
    if err != nil {
        return nil, err
    }

    
    var res *types.Message
    return res, json.Unmarshal(r, &res) 

}

// EditMessageLiveLocation methods's optional params
type EditMessageLiveLocationOpts struct {
    ChatId int64 `json:"chat_id,omitempty"`
    MessageId int64 `json:"message_id,omitempty"`
    InlineMessageId string `json:"inline_message_id,omitempty"`
    HorizontalAccuracy float64 `json:"horizontal_accuracy,omitempty"`
    Heading int64 `json:"heading,omitempty"`
    ProximityAlertRadius int64 `json:"proximity_alert_radius,omitempty"`
    ReplyMarkup *types.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// Use this method to edit live location messages. A location can be edited until its live_period expires or editing is explicitly disabled by a call to stopMessageLiveLocation. On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned.
func (b *Bot) EditMessageLiveLocation(latitude float64, longitude float64, opts *EditMessageLiveLocationOpts) (*types.Message, error) {
    params := map[string]string{}
    data_params := map[string]string{}

    params["latitude"] = strconv.FormatFloat(latitude, 'E', -1, 64)
    params["longitude"] = strconv.FormatFloat(longitude, 'E', -1, 64)
    if opts != nil {
        params["chat_id"] = strconv.FormatInt(opts.ChatId, 10)
        params["message_id"] = strconv.FormatInt(opts.MessageId, 10)
        params["inline_message_id"] = opts.InlineMessageId
        params["horizontal_accuracy"] = strconv.FormatFloat(opts.HorizontalAccuracy, 'E', -1, 64)
        params["heading"] = strconv.FormatInt(opts.Heading, 10)
        params["proximity_alert_radius"] = strconv.FormatInt(opts.ProximityAlertRadius, 10)

        if opts.ReplyMarkup != nil {
            bs, err := json.Marshal(opts.ReplyMarkup)
            if err != nil {
                return nil, fmt.Errorf("failed to marshal field reply_markup: %w", err)
            }
            params["reply_markup"] = string(bs)
        }

    }


    r, err := b.Request("editMessageLiveLocation", params, data_params)
    if err != nil {
        return nil, err
    }

    
    var res *types.Message
    return res, json.Unmarshal(r, &res) 

}

// StopMessageLiveLocation methods's optional params
type StopMessageLiveLocationOpts struct {
    ChatId int64 `json:"chat_id,omitempty"`
    MessageId int64 `json:"message_id,omitempty"`
    InlineMessageId string `json:"inline_message_id,omitempty"`
    ReplyMarkup *types.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// Use this method to stop updating a live location message before live_period expires. On success, if the message is not an inline message, the edited Message is returned, otherwise True is returned.
func (b *Bot) StopMessageLiveLocation(opts *StopMessageLiveLocationOpts) (*types.Message, error) {
    params := map[string]string{}
    data_params := map[string]string{}

    if opts != nil {
        params["chat_id"] = strconv.FormatInt(opts.ChatId, 10)
        params["message_id"] = strconv.FormatInt(opts.MessageId, 10)
        params["inline_message_id"] = opts.InlineMessageId

        if opts.ReplyMarkup != nil {
            bs, err := json.Marshal(opts.ReplyMarkup)
            if err != nil {
                return nil, fmt.Errorf("failed to marshal field reply_markup: %w", err)
            }
            params["reply_markup"] = string(bs)
        }

    }


    r, err := b.Request("stopMessageLiveLocation", params, data_params)
    if err != nil {
        return nil, err
    }

    
    var res *types.Message
    return res, json.Unmarshal(r, &res) 

}

// EditMessageReplyMarkup methods's optional params
type EditMessageReplyMarkupOpts struct {
    ChatId int64 `json:"chat_id,omitempty"`
    MessageId int64 `json:"message_id,omitempty"`
    InlineMessageId string `json:"inline_message_id,omitempty"`
    ReplyMarkup *types.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// Use this method to edit only the reply markup of messages. On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned.
func (b *Bot) EditMessageReplyMarkup(opts *EditMessageReplyMarkupOpts) (*types.Message, error) {
    params := map[string]string{}
    data_params := map[string]string{}

    if opts != nil {
        params["chat_id"] = strconv.FormatInt(opts.ChatId, 10)
        params["message_id"] = strconv.FormatInt(opts.MessageId, 10)
        params["inline_message_id"] = opts.InlineMessageId

        if opts.ReplyMarkup != nil {
            bs, err := json.Marshal(opts.ReplyMarkup)
            if err != nil {
                return nil, fmt.Errorf("failed to marshal field reply_markup: %w", err)
            }
            params["reply_markup"] = string(bs)
        }

    }


    r, err := b.Request("editMessageReplyMarkup", params, data_params)
    if err != nil {
        return nil, err
    }

    
    var res *types.Message
    return res, json.Unmarshal(r, &res) 

}

// StopPoll methods's optional params
type StopPollOpts struct {
    ReplyMarkup *types.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// Use this method to stop a poll which was sent by the bot. On success, the stopped Poll is returned.
func (b *Bot) StopPoll(chatId int64, messageId int64, opts *StopPollOpts) (*types.Poll, error) {
    params := map[string]string{}
    data_params := map[string]string{}

    params["chat_id"] = strconv.FormatInt(chatId, 10)
    params["message_id"] = strconv.FormatInt(messageId, 10)
    if opts != nil {

        if opts.ReplyMarkup != nil {
            bs, err := json.Marshal(opts.ReplyMarkup)
            if err != nil {
                return nil, fmt.Errorf("failed to marshal field reply_markup: %w", err)
            }
            params["reply_markup"] = string(bs)
        }

    }


    r, err := b.Request("stopPoll", params, data_params)
    if err != nil {
        return nil, err
    }

    
    var res *types.Poll
    return res, json.Unmarshal(r, &res) 

}

// Use this method to delete a message, including service messages, with the following limitations:
// - A message can only be deleted if it was sent less than 48 hours ago.
// - Service messages about a supergroup, channel, or forum topic creation can't be deleted.
// - A dice message in a private chat can only be deleted if it was sent more than 24 hours ago.
// - Bots can delete outgoing messages in private chats, groups, and supergroups.
// - Bots can delete incoming messages in private chats.
// - Bots granted can_post_messages permissions can delete outgoing messages in channels.
// - If the bot is an administrator of a group, it can delete any message there.
// - If the bot has can_delete_messages permission in a supergroup or a channel, it can delete any message there.
// Returns True on success.
func (b *Bot) DeleteMessage(chatId int64, messageId int64) (bool, error) {
    params := map[string]string{}
    data_params := map[string]string{}
    params["chat_id"] = strconv.FormatInt(chatId, 10)
    params["message_id"] = strconv.FormatInt(messageId, 10)

    r, err := b.Request("deleteMessage", params, data_params)
    if err != nil {
        return false, err
    }
    
    var res bool
    return res, json.Unmarshal(r, &res) 

}

// Use this method to delete multiple messages simultaneously. If some of the specified messages can't be found, they are skipped. Returns True on success.
func (b *Bot) DeleteMessages(chatId int64, messageIds []int64) (bool, error) {
    params := map[string]string{}
    data_params := map[string]string{}
    params["chat_id"] = strconv.FormatInt(chatId, 10)

    if messageIds != nil {
        bs, err := json.Marshal(messageIds)
        if err != nil {
            return false, fmt.Errorf("failed to marshal field message_ids: %w", err)
        }
        params["message_ids"] = string(bs)
    }


    r, err := b.Request("deleteMessages", params, data_params)
    if err != nil {
        return false, err
    }
    
    var res bool
    return res, json.Unmarshal(r, &res) 

}

// SendSticker methods's optional params
type SendStickerOpts struct {
    MessageThreadId int64 `json:"message_thread_id,omitempty"`
    Emoji string `json:"emoji,omitempty"`
    DisableNotification bool `json:"disable_notification,omitempty"`
    ProtectContent bool `json:"protect_content,omitempty"`
    ReplyParameters *types.ReplyParameters `json:"reply_parameters,omitempty"`
    ReplyMarkup types.ReplyMarkup `json:"reply_markup,omitempty"`
}

// Use this method to send static .WEBP, animated .TGS, or video .WEBM stickers. On success, the sent Message is returned.
func (b *Bot) SendSticker(chatId int64, sticker types.InputFile, opts *SendStickerOpts) (*types.Message, error) {
    params := map[string]string{}
    data_params := map[string]string{}

    params["chat_id"] = strconv.FormatInt(chatId, 10)

    if sticker != nil {
        switch f := sticker.(type) {
        case string:
            _, err := os.Stat(f)
            if err != nil {
                params["sticker"] = f
            } else {
                params["sticker"] = "attach://sticker"
                data_params["sticker"] = f
            }
        default:
            return nil, fmt.Errorf("unknown type for InputFile: %T", sticker)
        }
    }
    if opts != nil {
        params["message_thread_id"] = strconv.FormatInt(opts.MessageThreadId, 10)
        params["emoji"] = opts.Emoji
        params["disable_notification"] = strconv.FormatBool(opts.DisableNotification)
        params["protect_content"] = strconv.FormatBool(opts.ProtectContent)

        if opts.ReplyParameters != nil {
            bs, err := json.Marshal(opts.ReplyParameters)
            if err != nil {
                return nil, fmt.Errorf("failed to marshal field reply_parameters: %w", err)
            }
            params["reply_parameters"] = string(bs)
        }


        if opts.ReplyMarkup != nil {
            bs, err := json.Marshal(opts.ReplyMarkup)
            if err != nil {
                return nil, fmt.Errorf("failed to marshal field reply_markup: %w", err)
            }
            params["reply_markup"] = string(bs)
        }

    }


    r, err := b.Request("sendSticker", params, data_params)
    if err != nil {
        return nil, err
    }

    
    var res *types.Message
    return res, json.Unmarshal(r, &res) 

}

// Use this method to get a sticker set. On success, a StickerSet object is returned.
func (b *Bot) GetStickerSet(name string) (*types.StickerSet, error) {
    params := map[string]string{}
    data_params := map[string]string{}
    params["name"] = name

    r, err := b.Request("getStickerSet", params, data_params)
    if err != nil {
        return nil, err
    }
    
    var res *types.StickerSet
    return res, json.Unmarshal(r, &res) 

}

// Use this method to get information about custom emoji stickers by their identifiers. Returns an Array of Sticker objects.
func (b *Bot) GetCustomEmojiStickers(customEmojiIds []string) ([]types.Sticker, error) {
    params := map[string]string{}
    data_params := map[string]string{}

    if customEmojiIds != nil {
        bs, err := json.Marshal(customEmojiIds)
        if err != nil {
            return nil, fmt.Errorf("failed to marshal field custom_emoji_ids: %w", err)
        }
        params["custom_emoji_ids"] = string(bs)
    }


    r, err := b.Request("getCustomEmojiStickers", params, data_params)
    if err != nil {
        return nil, err
    }
    
    var res []types.Sticker
    return res, json.Unmarshal(r, &res) 

}

// Use this method to upload a file with a sticker for later use in the createNewStickerSet and addStickerToSet methods (the file can be used multiple times). Returns the uploaded File on success.
func (b *Bot) UploadStickerFile(userId int64, sticker types.InputFile, stickerFormat string) (*types.File, error) {
    params := map[string]string{}
    data_params := map[string]string{}
    params["user_id"] = strconv.FormatInt(userId, 10)

    if sticker != nil {
        switch f := sticker.(type) {
        case string:
            _, err := os.Stat(f)
            if err != nil {
                params["sticker"] = f
            } else {
                params["sticker"] = "attach://sticker"
                data_params["sticker"] = f
            }
        default:
            return nil, fmt.Errorf("unknown type for InputFile: %T", sticker)
        }
    }
    params["sticker_format"] = stickerFormat

    r, err := b.Request("uploadStickerFile", params, data_params)
    if err != nil {
        return nil, err
    }
    
    var res *types.File
    return res, json.Unmarshal(r, &res) 

}

// CreateNewStickerSet methods's optional params
type CreateNewStickerSetOpts struct {
    StickerType string `json:"sticker_type,omitempty"`
    NeedsRepainting bool `json:"needs_repainting,omitempty"`
}

// Use this method to create a new sticker set owned by a user. The bot will be able to edit the sticker set thus created. Returns True on success.
func (b *Bot) CreateNewStickerSet(userId int64, name string, title string, stickers []types.InputSticker, stickerFormat string, opts *CreateNewStickerSetOpts) (bool, error) {
    params := map[string]string{}
    data_params := map[string]string{}

    params["user_id"] = strconv.FormatInt(userId, 10)
    params["name"] = name
    params["title"] = title

    if stickers != nil {
        bs, err := json.Marshal(stickers)
        if err != nil {
            return false, fmt.Errorf("failed to marshal field stickers: %w", err)
        }
        params["stickers"] = string(bs)
    }

    params["sticker_format"] = stickerFormat
    if opts != nil {
        params["sticker_type"] = opts.StickerType
        params["needs_repainting"] = strconv.FormatBool(opts.NeedsRepainting)
    }


    r, err := b.Request("createNewStickerSet", params, data_params)
    if err != nil {
        return false, err
    }

    
    var res bool
    return res, json.Unmarshal(r, &res) 

}

// Use this method to add a new sticker to a set created by the bot. The format of the added sticker must match the format of the other stickers in the set. Emoji sticker sets can have up to 200 stickers. Animated and video sticker sets can have up to 50 stickers. Static sticker sets can have up to 120 stickers. Returns True on success.
func (b *Bot) AddStickerToSet(userId int64, name string, sticker *types.InputSticker) (bool, error) {
    params := map[string]string{}
    data_params := map[string]string{}
    params["user_id"] = strconv.FormatInt(userId, 10)
    params["name"] = name

    if sticker != nil {
        bs, err := json.Marshal(sticker)
        if err != nil {
            return false, fmt.Errorf("failed to marshal field sticker: %w", err)
        }
        params["sticker"] = string(bs)
    }


    r, err := b.Request("addStickerToSet", params, data_params)
    if err != nil {
        return false, err
    }
    
    var res bool
    return res, json.Unmarshal(r, &res) 

}

// Use this method to move a sticker in a set created by the bot to a specific position. Returns True on success.
func (b *Bot) SetStickerPositionInSet(sticker string, position int64) (bool, error) {
    params := map[string]string{}
    data_params := map[string]string{}
    params["sticker"] = sticker
    params["position"] = strconv.FormatInt(position, 10)

    r, err := b.Request("setStickerPositionInSet", params, data_params)
    if err != nil {
        return false, err
    }
    
    var res bool
    return res, json.Unmarshal(r, &res) 

}

// Use this method to delete a sticker from a set created by the bot. Returns True on success.
func (b *Bot) DeleteStickerFromSet(sticker string) (bool, error) {
    params := map[string]string{}
    data_params := map[string]string{}
    params["sticker"] = sticker

    r, err := b.Request("deleteStickerFromSet", params, data_params)
    if err != nil {
        return false, err
    }
    
    var res bool
    return res, json.Unmarshal(r, &res) 

}

// Use this method to change the list of emoji assigned to a regular or custom emoji sticker. The sticker must belong to a sticker set created by the bot. Returns True on success.
func (b *Bot) SetStickerEmojiList(sticker string, emojiList []string) (bool, error) {
    params := map[string]string{}
    data_params := map[string]string{}
    params["sticker"] = sticker

    if emojiList != nil {
        bs, err := json.Marshal(emojiList)
        if err != nil {
            return false, fmt.Errorf("failed to marshal field emoji_list: %w", err)
        }
        params["emoji_list"] = string(bs)
    }


    r, err := b.Request("setStickerEmojiList", params, data_params)
    if err != nil {
        return false, err
    }
    
    var res bool
    return res, json.Unmarshal(r, &res) 

}

// SetStickerKeywords methods's optional params
type SetStickerKeywordsOpts struct {
    Keywords []string `json:"keywords,omitempty"`
}

// Use this method to change search keywords assigned to a regular or custom emoji sticker. The sticker must belong to a sticker set created by the bot. Returns True on success.
func (b *Bot) SetStickerKeywords(sticker string, opts *SetStickerKeywordsOpts) (bool, error) {
    params := map[string]string{}
    data_params := map[string]string{}

    params["sticker"] = sticker
    if opts != nil {

        if opts.Keywords != nil {
            bs, err := json.Marshal(opts.Keywords)
            if err != nil {
                return false, fmt.Errorf("failed to marshal field keywords: %w", err)
            }
            params["keywords"] = string(bs)
        }

    }


    r, err := b.Request("setStickerKeywords", params, data_params)
    if err != nil {
        return false, err
    }

    
    var res bool
    return res, json.Unmarshal(r, &res) 

}

// SetStickerMaskPosition methods's optional params
type SetStickerMaskPositionOpts struct {
    MaskPosition *types.MaskPosition `json:"mask_position,omitempty"`
}

// Use this method to change the mask position of a mask sticker. The sticker must belong to a sticker set that was created by the bot. Returns True on success.
func (b *Bot) SetStickerMaskPosition(sticker string, opts *SetStickerMaskPositionOpts) (bool, error) {
    params := map[string]string{}
    data_params := map[string]string{}

    params["sticker"] = sticker
    if opts != nil {

        if opts.MaskPosition != nil {
            bs, err := json.Marshal(opts.MaskPosition)
            if err != nil {
                return false, fmt.Errorf("failed to marshal field mask_position: %w", err)
            }
            params["mask_position"] = string(bs)
        }

    }


    r, err := b.Request("setStickerMaskPosition", params, data_params)
    if err != nil {
        return false, err
    }

    
    var res bool
    return res, json.Unmarshal(r, &res) 

}

// Use this method to set the title of a created sticker set. Returns True on success.
func (b *Bot) SetStickerSetTitle(name string, title string) (bool, error) {
    params := map[string]string{}
    data_params := map[string]string{}
    params["name"] = name
    params["title"] = title

    r, err := b.Request("setStickerSetTitle", params, data_params)
    if err != nil {
        return false, err
    }
    
    var res bool
    return res, json.Unmarshal(r, &res) 

}

// SetStickerSetThumbnail methods's optional params
type SetStickerSetThumbnailOpts struct {
    Thumbnail types.InputFile `json:"thumbnail,omitempty"`
}

// Use this method to set the thumbnail of a regular or mask sticker set. The format of the thumbnail file must match the format of the stickers in the set. Returns True on success.
func (b *Bot) SetStickerSetThumbnail(name string, userId int64, opts *SetStickerSetThumbnailOpts) (bool, error) {
    params := map[string]string{}
    data_params := map[string]string{}

    params["name"] = name
    params["user_id"] = strconv.FormatInt(userId, 10)
    if opts != nil {

        if opts.Thumbnail != nil {
            switch f := opts.Thumbnail.(type) {
            case string:
                _, err := os.Stat(f)
                if err != nil {
                    params["thumbnail"] = f
                } else {
                    params["thumbnail"] = "attach://thumbnail"
                    data_params["thumbnail"] = f
                }
            default:
                return false, fmt.Errorf("unknown type for InputFile: %T", opts.Thumbnail)
            }
        }
    }


    r, err := b.Request("setStickerSetThumbnail", params, data_params)
    if err != nil {
        return false, err
    }

    
    var res bool
    return res, json.Unmarshal(r, &res) 

}

// SetCustomEmojiStickerSetThumbnail methods's optional params
type SetCustomEmojiStickerSetThumbnailOpts struct {
    CustomEmojiId string `json:"custom_emoji_id,omitempty"`
}

// Use this method to set the thumbnail of a custom emoji sticker set. Returns True on success.
func (b *Bot) SetCustomEmojiStickerSetThumbnail(name string, opts *SetCustomEmojiStickerSetThumbnailOpts) (bool, error) {
    params := map[string]string{}
    data_params := map[string]string{}

    params["name"] = name
    if opts != nil {
        params["custom_emoji_id"] = opts.CustomEmojiId
    }


    r, err := b.Request("setCustomEmojiStickerSetThumbnail", params, data_params)
    if err != nil {
        return false, err
    }

    
    var res bool
    return res, json.Unmarshal(r, &res) 

}

// Use this method to delete a sticker set that was created by the bot. Returns True on success.
func (b *Bot) DeleteStickerSet(name string) (bool, error) {
    params := map[string]string{}
    data_params := map[string]string{}
    params["name"] = name

    r, err := b.Request("deleteStickerSet", params, data_params)
    if err != nil {
        return false, err
    }
    
    var res bool
    return res, json.Unmarshal(r, &res) 

}

// AnswerInlineQuery methods's optional params
type AnswerInlineQueryOpts struct {
    CacheTime int64 `json:"cache_time,omitempty"`
    IsPersonal bool `json:"is_personal,omitempty"`
    NextOffset string `json:"next_offset,omitempty"`
    Button *types.InlineQueryResultsButton `json:"button,omitempty"`
}

// Use this method to send answers to an inline query. On success, True is returned.
// No more than 50 results per query are allowed.
func (b *Bot) AnswerInlineQuery(inlineQueryId string, results []types.InlineQueryResult, opts *AnswerInlineQueryOpts) (bool, error) {
    params := map[string]string{}
    data_params := map[string]string{}

    params["inline_query_id"] = inlineQueryId

    if results != nil {
        bs, err := json.Marshal(results)
        if err != nil {
            return false, fmt.Errorf("failed to marshal field results: %w", err)
        }
        params["results"] = string(bs)
    }

    if opts != nil {
        params["cache_time"] = strconv.FormatInt(opts.CacheTime, 10)
        params["is_personal"] = strconv.FormatBool(opts.IsPersonal)
        params["next_offset"] = opts.NextOffset

        if opts.Button != nil {
            bs, err := json.Marshal(opts.Button)
            if err != nil {
                return false, fmt.Errorf("failed to marshal field button: %w", err)
            }
            params["button"] = string(bs)
        }

    }


    r, err := b.Request("answerInlineQuery", params, data_params)
    if err != nil {
        return false, err
    }

    
    var res bool
    return res, json.Unmarshal(r, &res) 

}

// Use this method to set the result of an interaction with a Web App and send a corresponding message on behalf of the user to the chat from which the query originated. On success, a SentWebAppMessage object is returned.
func (b *Bot) AnswerWebAppQuery(webAppQueryId string, result *types.InlineQueryResult) (*types.SentWebAppMessage, error) {
    params := map[string]string{}
    data_params := map[string]string{}
    params["web_app_query_id"] = webAppQueryId

    if result != nil {
        bs, err := json.Marshal(result)
        if err != nil {
            return nil, fmt.Errorf("failed to marshal field result: %w", err)
        }
        params["result"] = string(bs)
    }


    r, err := b.Request("answerWebAppQuery", params, data_params)
    if err != nil {
        return nil, err
    }
    
    var res *types.SentWebAppMessage
    return res, json.Unmarshal(r, &res) 

}

// SendInvoice methods's optional params
type SendInvoiceOpts struct {
    MessageThreadId int64 `json:"message_thread_id,omitempty"`
    MaxTipAmount int64 `json:"max_tip_amount,omitempty"`
    SuggestedTipAmounts []int64 `json:"suggested_tip_amounts,omitempty"`
    StartParameter string `json:"start_parameter,omitempty"`
    ProviderData string `json:"provider_data,omitempty"`
    PhotoUrl string `json:"photo_url,omitempty"`
    PhotoSize int64 `json:"photo_size,omitempty"`
    PhotoWidth int64 `json:"photo_width,omitempty"`
    PhotoHeight int64 `json:"photo_height,omitempty"`
    NeedName bool `json:"need_name,omitempty"`
    NeedPhoneNumber bool `json:"need_phone_number,omitempty"`
    NeedEmail bool `json:"need_email,omitempty"`
    NeedShippingAddress bool `json:"need_shipping_address,omitempty"`
    SendPhoneNumberToProvider bool `json:"send_phone_number_to_provider,omitempty"`
    SendEmailToProvider bool `json:"send_email_to_provider,omitempty"`
    IsFlexible bool `json:"is_flexible,omitempty"`
    DisableNotification bool `json:"disable_notification,omitempty"`
    ProtectContent bool `json:"protect_content,omitempty"`
    ReplyParameters *types.ReplyParameters `json:"reply_parameters,omitempty"`
    ReplyMarkup *types.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// Use this method to send invoices. On success, the sent Message is returned.
func (b *Bot) SendInvoice(chatId int64, title string, description string, payload string, providerToken string, currency string, prices []types.LabeledPrice, opts *SendInvoiceOpts) (*types.Message, error) {
    params := map[string]string{}
    data_params := map[string]string{}

    params["chat_id"] = strconv.FormatInt(chatId, 10)
    params["title"] = title
    params["description"] = description
    params["payload"] = payload
    params["provider_token"] = providerToken
    params["currency"] = currency

    if prices != nil {
        bs, err := json.Marshal(prices)
        if err != nil {
            return nil, fmt.Errorf("failed to marshal field prices: %w", err)
        }
        params["prices"] = string(bs)
    }

    if opts != nil {
        params["message_thread_id"] = strconv.FormatInt(opts.MessageThreadId, 10)
        params["max_tip_amount"] = strconv.FormatInt(opts.MaxTipAmount, 10)

        if opts.SuggestedTipAmounts != nil {
            bs, err := json.Marshal(opts.SuggestedTipAmounts)
            if err != nil {
                return nil, fmt.Errorf("failed to marshal field suggested_tip_amounts: %w", err)
            }
            params["suggested_tip_amounts"] = string(bs)
        }

        params["start_parameter"] = opts.StartParameter
        params["provider_data"] = opts.ProviderData
        params["photo_url"] = opts.PhotoUrl
        params["photo_size"] = strconv.FormatInt(opts.PhotoSize, 10)
        params["photo_width"] = strconv.FormatInt(opts.PhotoWidth, 10)
        params["photo_height"] = strconv.FormatInt(opts.PhotoHeight, 10)
        params["need_name"] = strconv.FormatBool(opts.NeedName)
        params["need_phone_number"] = strconv.FormatBool(opts.NeedPhoneNumber)
        params["need_email"] = strconv.FormatBool(opts.NeedEmail)
        params["need_shipping_address"] = strconv.FormatBool(opts.NeedShippingAddress)
        params["send_phone_number_to_provider"] = strconv.FormatBool(opts.SendPhoneNumberToProvider)
        params["send_email_to_provider"] = strconv.FormatBool(opts.SendEmailToProvider)
        params["is_flexible"] = strconv.FormatBool(opts.IsFlexible)
        params["disable_notification"] = strconv.FormatBool(opts.DisableNotification)
        params["protect_content"] = strconv.FormatBool(opts.ProtectContent)

        if opts.ReplyParameters != nil {
            bs, err := json.Marshal(opts.ReplyParameters)
            if err != nil {
                return nil, fmt.Errorf("failed to marshal field reply_parameters: %w", err)
            }
            params["reply_parameters"] = string(bs)
        }


        if opts.ReplyMarkup != nil {
            bs, err := json.Marshal(opts.ReplyMarkup)
            if err != nil {
                return nil, fmt.Errorf("failed to marshal field reply_markup: %w", err)
            }
            params["reply_markup"] = string(bs)
        }

    }


    r, err := b.Request("sendInvoice", params, data_params)
    if err != nil {
        return nil, err
    }

    
    var res *types.Message
    return res, json.Unmarshal(r, &res) 

}

// CreateInvoiceLink methods's optional params
type CreateInvoiceLinkOpts struct {
    MaxTipAmount int64 `json:"max_tip_amount,omitempty"`
    SuggestedTipAmounts []int64 `json:"suggested_tip_amounts,omitempty"`
    ProviderData string `json:"provider_data,omitempty"`
    PhotoUrl string `json:"photo_url,omitempty"`
    PhotoSize int64 `json:"photo_size,omitempty"`
    PhotoWidth int64 `json:"photo_width,omitempty"`
    PhotoHeight int64 `json:"photo_height,omitempty"`
    NeedName bool `json:"need_name,omitempty"`
    NeedPhoneNumber bool `json:"need_phone_number,omitempty"`
    NeedEmail bool `json:"need_email,omitempty"`
    NeedShippingAddress bool `json:"need_shipping_address,omitempty"`
    SendPhoneNumberToProvider bool `json:"send_phone_number_to_provider,omitempty"`
    SendEmailToProvider bool `json:"send_email_to_provider,omitempty"`
    IsFlexible bool `json:"is_flexible,omitempty"`
}

// Use this method to create a link for an invoice. Returns the created invoice link as String on success.
func (b *Bot) CreateInvoiceLink(title string, description string, payload string, providerToken string, currency string, prices []types.LabeledPrice, opts *CreateInvoiceLinkOpts) (string, error) {
    params := map[string]string{}
    data_params := map[string]string{}

    params["title"] = title
    params["description"] = description
    params["payload"] = payload
    params["provider_token"] = providerToken
    params["currency"] = currency

    if prices != nil {
        bs, err := json.Marshal(prices)
        if err != nil {
            return "", fmt.Errorf("failed to marshal field prices: %w", err)
        }
        params["prices"] = string(bs)
    }

    if opts != nil {
        params["max_tip_amount"] = strconv.FormatInt(opts.MaxTipAmount, 10)

        if opts.SuggestedTipAmounts != nil {
            bs, err := json.Marshal(opts.SuggestedTipAmounts)
            if err != nil {
                return "", fmt.Errorf("failed to marshal field suggested_tip_amounts: %w", err)
            }
            params["suggested_tip_amounts"] = string(bs)
        }

        params["provider_data"] = opts.ProviderData
        params["photo_url"] = opts.PhotoUrl
        params["photo_size"] = strconv.FormatInt(opts.PhotoSize, 10)
        params["photo_width"] = strconv.FormatInt(opts.PhotoWidth, 10)
        params["photo_height"] = strconv.FormatInt(opts.PhotoHeight, 10)
        params["need_name"] = strconv.FormatBool(opts.NeedName)
        params["need_phone_number"] = strconv.FormatBool(opts.NeedPhoneNumber)
        params["need_email"] = strconv.FormatBool(opts.NeedEmail)
        params["need_shipping_address"] = strconv.FormatBool(opts.NeedShippingAddress)
        params["send_phone_number_to_provider"] = strconv.FormatBool(opts.SendPhoneNumberToProvider)
        params["send_email_to_provider"] = strconv.FormatBool(opts.SendEmailToProvider)
        params["is_flexible"] = strconv.FormatBool(opts.IsFlexible)
    }


    r, err := b.Request("createInvoiceLink", params, data_params)
    if err != nil {
        return "", err
    }

    
    var res string
    return res, json.Unmarshal(r, &res) 

}

// AnswerShippingQuery methods's optional params
type AnswerShippingQueryOpts struct {
    ShippingOptions []types.ShippingOption `json:"shipping_options,omitempty"`
    ErrorMessage string `json:"error_message,omitempty"`
}

// If you sent an invoice requesting a shipping address and the parameter is_flexible was specified, the Bot API will send an Update with a shipping_query field to the bot. Use this method to reply to shipping queries. On success, True is returned.
func (b *Bot) AnswerShippingQuery(shippingQueryId string, ok bool, opts *AnswerShippingQueryOpts) (bool, error) {
    params := map[string]string{}
    data_params := map[string]string{}

    params["shipping_query_id"] = shippingQueryId
    params["ok"] = strconv.FormatBool(ok)
    if opts != nil {

        if opts.ShippingOptions != nil {
            bs, err := json.Marshal(opts.ShippingOptions)
            if err != nil {
                return false, fmt.Errorf("failed to marshal field shipping_options: %w", err)
            }
            params["shipping_options"] = string(bs)
        }

        params["error_message"] = opts.ErrorMessage
    }


    r, err := b.Request("answerShippingQuery", params, data_params)
    if err != nil {
        return false, err
    }

    
    var res bool
    return res, json.Unmarshal(r, &res) 

}

// AnswerPreCheckoutQuery methods's optional params
type AnswerPreCheckoutQueryOpts struct {
    ErrorMessage string `json:"error_message,omitempty"`
}

// Once the user has confirmed their payment and shipping details, the Bot API sends the final confirmation in the form of an Update with the field pre_checkout_query. Use this method to respond to such pre-checkout queries. On success, True is returned. Note: The Bot API must receive an answer within 10 seconds after the pre-checkout query was sent.
func (b *Bot) AnswerPreCheckoutQuery(preCheckoutQueryId string, ok bool, opts *AnswerPreCheckoutQueryOpts) (bool, error) {
    params := map[string]string{}
    data_params := map[string]string{}

    params["pre_checkout_query_id"] = preCheckoutQueryId
    params["ok"] = strconv.FormatBool(ok)
    if opts != nil {
        params["error_message"] = opts.ErrorMessage
    }


    r, err := b.Request("answerPreCheckoutQuery", params, data_params)
    if err != nil {
        return false, err
    }

    
    var res bool
    return res, json.Unmarshal(r, &res) 

}

// Informs a user that some of the Telegram Passport elements they provided contains errors. The user will not be able to re-submit their Passport to you until the errors are fixed (the contents of the field for which you returned the error must change). Returns True on success.
// Use this if the data submitted by the user doesn't satisfy the standards your service requires for any reason. For example, if a birthday date seems invalid, a submitted document is blurry, a scan shows evidence of tampering, etc. Supply some details in the error message to make sure the user knows how to correct the issues.
func (b *Bot) SetPassportDataErrors(userId int64, errors []types.PassportElementError) (bool, error) {
    params := map[string]string{}
    data_params := map[string]string{}
    params["user_id"] = strconv.FormatInt(userId, 10)

    if errors != nil {
        bs, err := json.Marshal(errors)
        if err != nil {
            return false, fmt.Errorf("failed to marshal field errors: %w", err)
        }
        params["errors"] = string(bs)
    }


    r, err := b.Request("setPassportDataErrors", params, data_params)
    if err != nil {
        return false, err
    }
    
    var res bool
    return res, json.Unmarshal(r, &res) 

}

// SendGame methods's optional params
type SendGameOpts struct {
    MessageThreadId int64 `json:"message_thread_id,omitempty"`
    DisableNotification bool `json:"disable_notification,omitempty"`
    ProtectContent bool `json:"protect_content,omitempty"`
    ReplyParameters *types.ReplyParameters `json:"reply_parameters,omitempty"`
    ReplyMarkup *types.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// Use this method to send a game. On success, the sent Message is returned.
func (b *Bot) SendGame(chatId int64, gameShortName string, opts *SendGameOpts) (*types.Message, error) {
    params := map[string]string{}
    data_params := map[string]string{}

    params["chat_id"] = strconv.FormatInt(chatId, 10)
    params["game_short_name"] = gameShortName
    if opts != nil {
        params["message_thread_id"] = strconv.FormatInt(opts.MessageThreadId, 10)
        params["disable_notification"] = strconv.FormatBool(opts.DisableNotification)
        params["protect_content"] = strconv.FormatBool(opts.ProtectContent)

        if opts.ReplyParameters != nil {
            bs, err := json.Marshal(opts.ReplyParameters)
            if err != nil {
                return nil, fmt.Errorf("failed to marshal field reply_parameters: %w", err)
            }
            params["reply_parameters"] = string(bs)
        }


        if opts.ReplyMarkup != nil {
            bs, err := json.Marshal(opts.ReplyMarkup)
            if err != nil {
                return nil, fmt.Errorf("failed to marshal field reply_markup: %w", err)
            }
            params["reply_markup"] = string(bs)
        }

    }


    r, err := b.Request("sendGame", params, data_params)
    if err != nil {
        return nil, err
    }

    
    var res *types.Message
    return res, json.Unmarshal(r, &res) 

}

// SetGameScore methods's optional params
type SetGameScoreOpts struct {
    Force bool `json:"force,omitempty"`
    DisableEditMessage bool `json:"disable_edit_message,omitempty"`
    ChatId int64 `json:"chat_id,omitempty"`
    MessageId int64 `json:"message_id,omitempty"`
    InlineMessageId string `json:"inline_message_id,omitempty"`
}

// Use this method to set the score of the specified user in a game message. On success, if the message is not an inline message, the Message is returned, otherwise True is returned. Returns an error, if the new score is not greater than the user's current score in the chat and force is False.
func (b *Bot) SetGameScore(userId int64, score int64, opts *SetGameScoreOpts) (*types.Message, error) {
    params := map[string]string{}
    data_params := map[string]string{}

    params["user_id"] = strconv.FormatInt(userId, 10)
    params["score"] = strconv.FormatInt(score, 10)
    if opts != nil {
        params["force"] = strconv.FormatBool(opts.Force)
        params["disable_edit_message"] = strconv.FormatBool(opts.DisableEditMessage)
        params["chat_id"] = strconv.FormatInt(opts.ChatId, 10)
        params["message_id"] = strconv.FormatInt(opts.MessageId, 10)
        params["inline_message_id"] = opts.InlineMessageId
    }


    r, err := b.Request("setGameScore", params, data_params)
    if err != nil {
        return nil, err
    }

    
    var res *types.Message
    return res, json.Unmarshal(r, &res) 

}

// GetGameHighScores methods's optional params
type GetGameHighScoresOpts struct {
    ChatId int64 `json:"chat_id,omitempty"`
    MessageId int64 `json:"message_id,omitempty"`
    InlineMessageId string `json:"inline_message_id,omitempty"`
}

// Use this method to get data for high score tables. Will return the score of the specified user and several of their neighbors in a game. Returns an Array of GameHighScore objects.
func (b *Bot) GetGameHighScores(userId int64, opts *GetGameHighScoresOpts) ([]types.GameHighScore, error) {
    params := map[string]string{}
    data_params := map[string]string{}

    params["user_id"] = strconv.FormatInt(userId, 10)
    if opts != nil {
        params["chat_id"] = strconv.FormatInt(opts.ChatId, 10)
        params["message_id"] = strconv.FormatInt(opts.MessageId, 10)
        params["inline_message_id"] = opts.InlineMessageId
    }


    r, err := b.Request("getGameHighScores", params, data_params)
    if err != nil {
        return nil, err
    }

    
    var res []types.GameHighScore
    return res, json.Unmarshal(r, &res) 

}
