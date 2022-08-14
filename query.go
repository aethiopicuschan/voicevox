package voicevox

/*
クエリ関連はここ
クエリ作成
https://voicevox.github.io/voicevox_engine/api/#tag/%E3%82%AF%E3%82%A8%E3%83%AA%E4%BD%9C%E6%88%90
クエリ編集
https://voicevox.github.io/voicevox_engine/api/#tag/%E3%82%AF%E3%82%A8%E3%83%AA%E7%B7%A8%E9%9B%86
*/

import (
	"encoding/json"
	"fmt"
)

/*
クエリ作成 > 音声合成用のクエリを作成する
https://voicevox.github.io/voicevox_engine/api/#tag/%E3%82%AF%E3%82%A8%E3%83%AA%E4%BD%9C%E6%88%90/operation/audio_query_audio_query_post
*/
func (c *Client) CreateQuery(speaker int, text string) (*Query, error) {
	url := c.url()
	url.Path = "/audio_query"
	q := url.Query()
	q.Set("speaker", fmt.Sprint(speaker))
	q.Set("text", text)
	url.RawQuery = q.Encode()
	_, body, err := c.request("POST", url)
	if err != nil {
		return nil, err
	}
	var query Query
	err = json.Unmarshal(body, &query)
	return &query, err
}

/*
クエリ作成 > 音声合成用のクエリをプリセットを用いて作成する
https://voicevox.github.io/voicevox_engine/api/#tag/%E3%82%AF%E3%82%A8%E3%83%AA%E4%BD%9C%E6%88%90/operation/audio_query_from_preset_audio_query_from_preset_post
*/
func (c *Client) CreateQueryWithPreset(preset int, text string) (*Query, error) {
	url := c.url()
	url.Path = "/audio_query_from_preset"
	q := url.Query()
	q.Set("preset_id", fmt.Sprint(preset))
	q.Set("text", text)
	url.RawQuery = q.Encode()
	_, body, err := c.request("POST", url)
	if err != nil {
		return nil, err
	}
	var query Query
	err = json.Unmarshal(body, &query)
	return &query, err
}

/*
クエリ編集 > テキストからアクセント句を得る
https://voicevox.github.io/voicevox_engine/api/#tag/%E3%82%AF%E3%82%A8%E3%83%AA%E7%B7%A8%E9%9B%86/operation/accent_phrases_accent_phrases_post
*/
func (c *Client) GetAccentPhrases(speaker int, text string, isKana bool) (*[]AccentPhrases, error) {
	url := c.url()
	url.Path = "/accent_phrases"
	q := url.Query()
	q.Set("speaker", fmt.Sprint(speaker))
	q.Set("text", text)
	q.Set("is_kana", fmt.Sprint(isKana))
	url.RawQuery = q.Encode()
	_, body, err := c.request("POST", url)
	if err != nil {
		return nil, err
	}
	var accentPhrases []AccentPhrases
	err = json.Unmarshal(body, &accentPhrases)
	return &accentPhrases, err
}

/*
クエリ編集 > アクセント句から音高・音素長を得る
https://voicevox.github.io/voicevox_engine/api/#tag/%E3%82%AF%E3%82%A8%E3%83%AA%E7%B7%A8%E9%9B%86/operation/mora_data_mora_data_post
*/
func (c *Client) GetPitchAndPhonemeLength(speaker int, accentPhrases *[]AccentPhrases) (*[]AccentPhrases, error) {
	url := c.url()
	url.Path = "/mora_data"
	q := url.Query()
	q.Set("speaker", fmt.Sprint(speaker))
	url.RawQuery = q.Encode()
	ap, err := json.Marshal(accentPhrases)
	if err != nil {
		return nil, err
	}
	_, body, err := c.request("POST", url, ap)
	if err != nil {
		return nil, err
	}
	var result []AccentPhrases
	err = json.Unmarshal(body, &result)
	return &result, err
}

/*
クエリ編集 > アクセント句から音素長を得る
https://voicevox.github.io/voicevox_engine/api/#tag/%E3%82%AF%E3%82%A8%E3%83%AA%E7%B7%A8%E9%9B%86/operation/mora_length_mora_length_post
*/
func (c *Client) GetPhonemeLength(speaker int, accentPhrases *[]AccentPhrases) (*[]AccentPhrases, error) {
	url := c.url()
	url.Path = "/mora_length"
	q := url.Query()
	q.Set("speaker", fmt.Sprint(speaker))
	url.RawQuery = q.Encode()
	ap, err := json.Marshal(accentPhrases)
	if err != nil {
		return nil, err
	}
	_, body, err := c.request("POST", url, ap)
	if err != nil {
		return nil, err
	}
	var result []AccentPhrases
	err = json.Unmarshal(body, &result)
	return &result, err
}

/*
クエリ編集 > アクセント句から音高を得る
https://voicevox.github.io/voicevox_engine/api/#tag/%E3%82%AF%E3%82%A8%E3%83%AA%E7%B7%A8%E9%9B%86/operation/mora_pitch_mora_pitch_post
*/
func (c *Client) GetPitch(speaker int, accentPhrases *[]AccentPhrases) (*[]AccentPhrases, error) {
	url := c.url()
	url.Path = "/mora_pitch"
	q := url.Query()
	q.Set("speaker", fmt.Sprint(speaker))
	url.RawQuery = q.Encode()
	ap, err := json.Marshal(accentPhrases)
	if err != nil {
		return nil, err
	}
	_, body, err := c.request("POST", url, ap)
	if err != nil {
		return nil, err
	}
	var result []AccentPhrases
	err = json.Unmarshal(body, &result)
	return &result, err
}
