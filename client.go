package nano

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Client struct {
	url string
}

func NewClient(url string) *Client {
	c := &Client{
		url: url,
	}

	return c
}

func (c *Client) Stop() (bool, error) {
	r, err := c.call("stop", nil)
	if err != nil {
		return false, err
	}

	if _, ok := r["success"]; !ok {
		return false, nil
	}

	return true, nil
}

func (c *Client) Version() (map[string]string, error) {
	r, err := c.call("version", nil)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (c *Client) call(action string, payload map[string]interface{}) (map[string]string, error) {
	if payload == nil {
		payload = make(map[string]interface{})
	}

	payload["action"] = action
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	res, err := http.Post(c.url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	raw, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var data map[string]string
	if err = json.Unmarshal(raw, &data); err != nil {
		return nil, err
	}

	return data, nil
}
