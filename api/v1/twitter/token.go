package twitter

import (
	"encoding/base64"
	"encoding/json"
)

type Token string

type Page struct {
	PreID         int64 `json:"pre_id"`
	NextTimeAtUTC int64 `json:"next_time_at_utc"`
	PageSize      int64 `json:"page_size"`
}

func (p Page) Encode() Token {
	b, err := json.Marshal(p)
	if err != nil {
		return Token("")
	}
	return Token(base64.StdEncoding.EncodeToString(b))
}

func (t Token) Decode() Page {
	var result Page
	if len(t) == 0 {
		return result
	}
	bytes, err := base64.StdEncoding.DecodeString(string(t))
	if err != nil {
		return result
	}
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return result
	}
	return result
}
