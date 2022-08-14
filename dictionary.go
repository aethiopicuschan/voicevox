package voicevox

/*
ユーザー辞書関連はここ
ユーザー辞書
https://voicevox.github.io/voicevox_engine/api/#tag/%E3%83%A6%E3%83%BC%E3%82%B6%E3%83%BC%E8%BE%9E%E6%9B%B8
*/

import (
	"encoding/json"
	"fmt"
)

/*
その他 > Get User Dict Words
https://voicevox.github.io/voicevox_engine/api/#tag/%E3%83%A6%E3%83%BC%E3%82%B6%E3%83%BC%E8%BE%9E%E6%9B%B8/operation/get_user_dict_words_user_dict_get
*/
func (c *Client) GetUserDict() (dict Dict, err error) {
	url := c.url()
	url.Path = "/user_dict"
	_, body, err := c.request("GET", url)
	err = json.Unmarshal(body, &dict)
	return
}

/*
その他 > Add User Dict Word
https://voicevox.github.io/voicevox_engine/api/#tag/%E3%83%A6%E3%83%BC%E3%82%B6%E3%83%BC%E8%BE%9E%E6%9B%B8/operation/add_user_dict_word_user_dict_word_post
*/
func (c *Client) AddUserDictWord(surface string, pronunciation string, accentType int, wordType *string, priority *int) (uuid string, err error) {
	url := c.url()
	url.Path = "/user_dict_word"
	q := url.Query()
	q.Set("surface", surface)
	q.Set("pronunciation", pronunciation)
	q.Set("accent_type", fmt.Sprint(accentType))
	if wordType != nil {
		q.Set("word_type", *wordType)
	}
	if priority != nil {
		q.Set("priority", fmt.Sprint(*priority))
	}
	url.RawQuery = q.Encode()
	_, body, err := c.request("POST", url)
	uuid = string(body)
	return
}

/*
その他 > Rewrite User Dict Word
https://voicevox.github.io/voicevox_engine/api/#tag/%E3%83%A6%E3%83%BC%E3%82%B6%E3%83%BC%E8%BE%9E%E6%9B%B8/operation/rewrite_user_dict_word_user_dict_word__word_uuid__put
*/
func (c *Client) RewriteUserDictWord(uuid string, surface string, pronunciation string, accentType int, wordType *string, priority *int) (err error) {
	url := c.url()
	url.Path = fmt.Sprintf("/user_dict_word/%s", uuid)
	q := url.Query()
	q.Set("surface", surface)
	q.Set("pronunciation", pronunciation)
	q.Set("accent_type", fmt.Sprint(accentType))
	if wordType != nil {
		q.Set("word_type", *wordType)
	}
	if priority != nil {
		q.Set("priority", fmt.Sprint(*priority))
	}
	url.RawQuery = q.Encode()
	_, _, err = c.request("PUT", url)
	return
}

/*
その他 > Delete User Dict Word
https://voicevox.github.io/voicevox_engine/api/#tag/%E3%83%A6%E3%83%BC%E3%82%B6%E3%83%BC%E8%BE%9E%E6%9B%B8/operation/delete_user_dict_word_user_dict_word__word_uuid__delete
*/
func (c *Client) DeleteUserDictWord(uuid string) (err error) {
	url := c.url()
	url.Path = fmt.Sprintf("/user_dict_word/%s", uuid)
	_, _, err = c.request("DELETE", url)
	return
}

/*
その他 > Import User Dict Words
https://voicevox.github.io/voicevox_engine/api/#tag/%E3%83%A6%E3%83%BC%E3%82%B6%E3%83%BC%E8%BE%9E%E6%9B%B8/operation/import_user_dict_words_import_user_dict_post
*/
func (c *Client) ImportUserDict(dict Dict, override bool) (err error) {
	url := c.url()
	url.Path = "/import_user_dict"
	q := url.Query()
	q.Set("override", fmt.Sprint(override))
	url.RawQuery = q.Encode()
	jsonDict, err := json.Marshal(dict)
	if err != nil {
		return
	}
	_, _, err = c.request("POST", url, jsonDict)
	return
}
