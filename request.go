package bot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/KeralaBots/GoTGramBot/types"
)

const (
	APIURL  = "https://api.telegram.org"
	TIMEOUT = time.Second * 5
)

type Response struct {
	Ok          bool                      `json:"ok"`
	Result      json.RawMessage           `json:"result"`
	ErrorCode   int                       `json:"error_code"`
	Description string                    `json:"description"`
	Parameters  *types.ResponseParameters `json:"parameters"`
}

type FileReader struct {
	FileName string
	File     []byte
}

func GetApiURL(token string, method string) string {
	return fmt.Sprintf("%s/bot%s/%s", APIURL, token, method)
}

func (b *Bot) Request(method string, params map[string]string, files map[string]string) (json.RawMessage, error) {
	custom_byte := &bytes.Buffer{}
	contentType, err := generateContentType(params, files, custom_byte)
	if err != nil {
		return nil, fmt.Errorf("failed to generate content-type: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, GetApiURL(b.Token, method), custom_byte)
	if err != nil {
		return nil, fmt.Errorf("post request failed: %e", err)
	}

	req.Header.Set("Content-Type", contentType)

	res, err := b.Client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("post execution failed: %e", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %e", err)
	}

	var response Response
	json.Unmarshal(body, &response)

	res.Body.Close()
	if !response.Ok {
		return nil, fmt.Errorf("telegram error [%d] : %s", response.ErrorCode, err.Error())
	}

	return response.Result, nil
}
