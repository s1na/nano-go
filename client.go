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

// Returns how many RAW is owned and how
// many have not yet been received by account.
func (c *Client) AccountBalance(account string) (map[string]string, error) {
	payload := map[string]interface{}{
		"account": account,
	}

	r, err := c.call("account_balance", payload)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Returns number of blocks for a specific account.
func (c *Client) AccountBlockCount(account string) (map[string]string, error) {
	payload := map[string]interface{}{
		"account": account,
	}

	r, err := c.call("account_block_count", payload)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Returns frontier, open block, change representative block,
// balance, last modified timestamp from local database
// and block count for account.
// Additionally returns representative, voting weight and
// pending balance for account, if respective parameters are set (>= v8.1).
func (c *Client) AccountInfo(account string, representative, weight, pending bool) (map[string]string, error) {
	payload := map[string]interface{}{
		"account":        account,
		"representative": representative,
		"weight":         weight,
		"pending":        pending,
	}

	r, err := c.call("account_info", payload)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Creates a new account, insert next deterministic key in wallet.
// If work is false, it disables work generation after creating account (>= v8.1).
// Requires enable_control
func (c *Client) AccountCreate(wallet string, work bool) (map[string]string, error) {
	payload := map[string]interface{}{
		"wallet": wallet,
		"work":   work,
	}

	r, err := c.call("account_create", payload)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Returns account number corresponding to the public key.
func (c *Client) AccountGet(key string) (map[string]string, error) {
	payload := map[string]interface{}{
		"key": key,
	}

	r, err := c.call("account_get", payload)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Reports send/receive information for a account.
func (c *Client) AccountHistory(account string) (map[string]string, error) {
	payload := map[string]interface{}{
		"account": account,
	}

	r, err := c.call("account_history", payload)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Lists all the accounts inside wallet.
func (c *Client) AccountList(wallet string) (map[string]string, error) {
	payload := map[string]interface{}{
		"wallet": wallet,
	}

	r, err := c.call("account_list", payload)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Moves accounts from source to wallet.
// Requires enable_control.
func (c *Client) AccountMove(wallet, source string, accounts []string) (map[string]string, error) {
	payload := map[string]interface{}{
		"wallet":   wallet,
		"source":   source,
		"accounts": accounts,
	}

	r, err := c.call("account_move", payload)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Returns the public key for account.
func (c *Client) AccountKey(account string) (map[string]string, error) {
	payload := map[string]interface{}{
		"account": account,
	}

	r, err := c.call("account_key", payload)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Removes account from wallet.
// Requires enable_control.
func (c *Client) AccountRemove(wallet, account string) (map[string]string, error) {
	payload := map[string]interface{}{
		"wallet":  wallet,
		"account": account,
	}

	r, err := c.call("account_remove", payload)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Returns the representative for account.
func (c *Client) AccountRepresentative(account string) (map[string]string, error) {
	payload := map[string]interface{}{
		"account": account,
	}

	r, err := c.call("account_representative", payload)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Sets the representative for account in wallet.
// If provided, uses work value for block from external source (>= v8.1).
// Requires enable_control.
func (c *Client) AccountRepresentativeSet(wallet, account, representative, work string) (map[string]string, error) {
	payload := map[string]interface{}{
		"wallet":         wallet,
		"account":        account,
		"representative": representative,
	}

	if work != "" {
		payload["work"] = work
	}

	r, err := c.call("account_representative_set", payload)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Returns the voting weight for account.
func (c *Client) AccountWeight(account string) (map[string]string, error) {
	payload := map[string]interface{}{
		"account": account,
	}

	r, err := c.call("account_weight", payload)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Stops the node safely.
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

// Returns version information for RPC, Store & Node (Major & Minor version).
// RPC Version always retruns "1" as of 13/01/2018.
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
