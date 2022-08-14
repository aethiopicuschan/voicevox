package voicevox

/*
音声関連はここ
音声合成
https://voicevox.github.io/voicevox_engine/api/#tag/%E9%9F%B3%E5%A3%B0%E5%90%88%E6%88%90
*/

import (
	"encoding/json"
	"fmt"
)

/*
音声合成 > 音声合成する
https://voicevox.github.io/voicevox_engine/api/#tag/%E9%9F%B3%E5%A3%B0%E5%90%88%E6%88%90/operation/synthesis_synthesis_post
*/
func (c *Client) CreateVoice(speaker int, enableInterrogativeUpspeak bool, query *Query) (wav []byte, err error) {
	url := c.url()
	url.Path = "/synthesis"
	q := url.Query()
	q.Set("speaker", fmt.Sprint(speaker))
	q.Set("enable_interrogative_upspeak", fmt.Sprint(enableInterrogativeUpspeak))
	url.RawQuery = q.Encode()
	jsonQuery, err := json.Marshal(query)
	if err != nil {
		return nil, err
	}
	_, wav, err = c.request("POST", url, jsonQuery)
	return
}

/*
音声合成 > 複数まとめて音声合成する
https://voicevox.github.io/voicevox_engine/api/#tag/%E9%9F%B3%E5%A3%B0%E5%90%88%E6%88%90/operation/multi_synthesis_multi_synthesis_post
*/
func (c *Client) CreateVoiceMulti(speaker int, query *Query) (zip []byte, err error) {
	url := c.url()
	url.Path = "/multi_synthesis"
	q := url.Query()
	q.Set("speaker", fmt.Sprint(speaker))
	url.RawQuery = q.Encode()
	jsonQuery, err := json.Marshal(query)
	if err != nil {
		return nil, err
	}
	_, zip, err = c.request("POST", url, jsonQuery)
	return
}

/*
音声合成 > 2人の話者でモーフィングした音声を合成する
https://voicevox.github.io/voicevox_engine/api/#tag/%E9%9F%B3%E5%A3%B0%E5%90%88%E6%88%90/operation/_synthesis_morphing_synthesis_morphing_post
*/
func (c *Client) CreateVoiceMorphing(base int, target int, rate float64, query *Query) (wav []byte, err error) {
	url := c.url()
	url.Path = "/synthesis_morphing"
	q := url.Query()
	q.Set("base_speaker", fmt.Sprint(base))
	q.Set("target_speaker", fmt.Sprint(target))
	q.Set("morph_rate", fmt.Sprint(rate))
	url.RawQuery = q.Encode()
	jsonQuery, err := json.Marshal(query)
	if err != nil {
		return nil, err
	}
	_, wav, err = c.request("POST", url, jsonQuery)
	return
}
