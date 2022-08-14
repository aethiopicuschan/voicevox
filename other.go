package voicevox

/*
その他扱いされてるやつはここ
その他
https://voicevox.github.io/voicevox_engine/api/#tag/%E3%81%9D%E3%81%AE%E4%BB%96
*/

import (
	"encoding/json"
)

/*
その他 > base64エンコードされた複数のwavデータを一つに結合する
https://voicevox.github.io/voicevox_engine/api/#tag/%E3%81%9D%E3%81%AE%E4%BB%96/operation/connect_waves_connect_waves_post
*/
func (c *Client) MargeWav(wavs []string) (wav []byte, err error) {
	url := c.url()
	url.Path = "/connect_waves"
	json, err := json.Marshal(wavs)
	if err != nil {
		return nil, err
	}
	_, wav, err = c.request("POST", url, json)
	return wav, err
}

/*
その他 > Get Presets
https://voicevox.github.io/voicevox_engine/api/#tag/%E3%81%9D%E3%81%AE%E4%BB%96/operation/get_presets_presets_get
*/
func (c *Client) GetPresets() (presets []Preset, err error) {
	url := c.url()
	url.Path = "/presets"
	_, body, err := c.request("GET", url)
	err = json.Unmarshal(body, &presets)
	return presets, err
}

/*
その他 > Version
https://voicevox.github.io/voicevox_engine/api/#tag/%E3%81%9D%E3%81%AE%E4%BB%96/operation/version_version_get
*/
func (c *Client) GetVersion() (version string, err error) {
	url := c.url()
	url.Path = "/version"
	_, body, err := c.request("GET", url)
	err = json.Unmarshal(body, &version)
	return version, err
}

/*
その他 > Speakers
https://voicevox.github.io/voicevox_engine/api/#tag/%E3%81%9D%E3%81%AE%E4%BB%96/operation/speakers_speakers_get
*/
func (c *Client) GetSpeakers() (speakers []Speaker, err error) {
	url := c.url()
	url.Path = "/speakers"
	_, body, err := c.request("GET", url)
	err = json.Unmarshal(body, &speakers)
	return speakers, err
}

/*
その他 > Speaker Info
https://voicevox.github.io/voicevox_engine/api/#tag/%E3%81%9D%E3%81%AE%E4%BB%96/operation/speaker_info_speaker_info_get
*/
func (c *Client) GetSpeakerInfo(uuid string) (info SpeakerInfo, err error) {
	url := c.url()
	url.Path = "/speaker_info"
	q := url.Query()
	q.Set("speaker_uuid", uuid)
	url.RawQuery = q.Encode()
	_, body, err := c.request("GET", url)
	err = json.Unmarshal(body, &info)
	return info, err
}
