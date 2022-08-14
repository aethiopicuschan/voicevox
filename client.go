package voicevox

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	scheme string
	host   string
}

func NewClient(scheme string, host string) *Client {
	client := new(Client)
	client.scheme = scheme
	client.host = host
	return client
}

// 各関数はこれを呼んで *url.URL を入手し、そこにパスとクエリを追加して request に渡す
func (c *Client) url() *url.URL {
	url := &url.URL{}
	url.Scheme = c.scheme
	url.Host = c.host
	return url
}

// bodyにjsonを含めるときは適切にヘッダーをセットしたりする
func newRequest(method string, url string, json [][]byte) (*http.Request, error) {
	if len(json) > 0 {
		req, err := http.NewRequest(method, url, bytes.NewBuffer(json[len(json)-1]))
		if err != nil {
			return nil, err
		}
		req.Header.Set("Content-Type", "application/json")
		return req, nil
	}
	return http.NewRequest(method, url, nil)
}

// bodyのjsonをoptionalにするために可変長引数を利用している
func (c *Client) request(method string, url *url.URL, json ...[]byte) (*http.Response, []byte, error) {
	req, err := newRequest(method, url.String(), json)
	if err != nil {
		return nil, nil, err
	}
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}
	if resp.StatusCode != 200 && resp.StatusCode != 204 {
		return resp, body, errors.New(fmt.Sprintf("StatusCode was %s.", resp.Status))
	}
	return resp, body, err
}
