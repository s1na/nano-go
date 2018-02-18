package rpc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

var (
	client = NewClient("http://localhost:7076")
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

func SetRPCServer(url string) {
	client.url = url
}

// Derive deterministic keypair from seed based on index.
func DeterministicKey(seed string, index int) (map[string]string, error) {
	payload := map[string]interface{}{
		"seed":  seed,
		"index": index,
	}

	return client.fetchMap("deterministic_key", payload, "")
}

// Generates an adhoc random keypair.
func KeyCreate() (map[string]string, error) {
	return client.fetchMap("key_create", nil, "")
}

// Derives public key and account number from private key.
func KeyExpand(key string) (map[string]string, error) {
	payload := map[string]interface{}{
		"key": key,
	}

	return client.fetchMap("key_expand", payload, "")
}

// Retrieves unchecked database keys, blocks hashes & a json
// representations of unchecked pending blocks
// starting from key up to count (>= v8.0).
func UncheckedKeys(key string, count int) (map[string]interface{}, error) {
	payload := map[string]interface{}{
		"key":   key,
		"count": count,
	}

	return client.fetchMapInterface("unchecked_keys", payload, "unchecked")
}

func (c *Client) call(action string, payload map[string]interface{}) ([]byte, error) {
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

	return ioutil.ReadAll(res.Body)
}

func (c *Client) fetchMap(action string, payload map[string]interface{}, key string) (map[string]string, error) {
	raw, err := c.call(action, payload)
	if err != nil {
		return nil, err
	}

	var r map[string]string
	if key == "" {
		if err = json.Unmarshal(raw, &r); err != nil {
			return nil, err
		}
	} else {
		var data map[string]map[string]string
		if err = json.Unmarshal(raw, &data); err != nil {
			return nil, err
		}

		var ok bool
		r, ok = data[key]
		if !ok {
			return nil, fmt.Errorf("Response of %s doesn't contain key %s.\n", action, key)
		}
	}

	return r, nil
}

func (c *Client) fetchMapInterface(action string, payload map[string]interface{}, key string) (map[string]interface{}, error) {
	raw, err := c.call(action, payload)
	if err != nil {
		return nil, err
	}

	var r map[string]interface{}
	if key == "" {
		if err = json.Unmarshal(raw, &r); err != nil {
			return nil, err
		}
	} else {
		var data map[string]map[string]interface{}
		if err = json.Unmarshal(raw, &data); err != nil {
			return nil, err
		}

		var ok bool
		r, ok = data[key]
		if !ok {
			return nil, fmt.Errorf("Response of %s doesn't contain key %s.\n", action, key)
		}
	}

	return r, nil
}

func (c *Client) isSuccess(action string, payload map[string]interface{}, key string) (bool, error) {
	r, err := c.fetchMap(action, payload, "")
	if err != nil {
		return false, err
	}

	var ok bool
	if key == "" {
		_, ok = r["success"]
	} else {
		r, exists := r[key]
		ok = exists && r == "1"
	}

	return ok, nil
}

func (c *Client) fetchString(action string, payload map[string]interface{}, key string) (string, error) {
	r, err := c.fetchMap(action, payload, "")
	if err != nil {
		return "", err
	}

	val, ok := r[key]
	if !ok {
		return "", fmt.Errorf("Response of %s doesn't contain key %s.\n", action, key)
	}

	return val, nil
}

func (c *Client) fetchInt(action string, payload map[string]interface{}, key string) (int, error) {
	r, err := c.fetchMap(action, payload, "")
	if err != nil {
		return 0, err
	}

	raw, ok := r[key]
	if !ok {
		return 0, fmt.Errorf("Response of %s doesn't contain key %s.\n", action, key)
	}

	return strconv.Atoi(raw)
}

func (c *Client) fetchInterface(action string, payload map[string]interface{}, key string) (interface{}, error) {
	r, err := c.fetchMapInterface(action, payload, "")
	if err != nil {
		return nil, err
	}

	val, ok := r[key]
	if !ok {
		return nil, fmt.Errorf("Response of %s doesn't contain key %s.\n", action, key)
	}

	return val, nil
}

func (c *Client) fetchSlice(action string, payload map[string]interface{}, key string) ([]string, error) {
	rawVal, err := c.fetchInterface(action, payload, key)
	if err != nil {
		return nil, err
	}

	switch val := rawVal.(type) {
	case []string:
		return val, nil
	case string:
		if val == "" {
			return []string{}, nil
		} else {
			return nil, fmt.Errorf("Key %s in response of %s contains a string instead of slice.\n", key, action)
		}
	case []interface{}:
		v := make([]string, len(val))
		for i, item := range val {
			v[i] = item.(string)
		}
		return v, nil
	default:
		return nil, fmt.Errorf("Key %s in response of %s contains an invalid type %T.\n", key, action, val)
	}
}
