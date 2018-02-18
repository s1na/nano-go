package rpc

import (
	"encoding/json"
	"errors"
)

type Block struct {
	Hash           string
	Work           string
	Previous       string
	Source         string
	Root           string
	Representative string
}

// Retrieves a json representation of block.
func GetBlock(hash string) (map[string]string, error) {
	payload := map[string]interface{}{
		"hash": hash,
	}

	return client.fetchMap("block", payload, "contents")
}

// Retrieves a json representations of blocks.
func Blocks(hashes []string) (map[string]map[string]string, error) {
	payload := map[string]interface{}{
		"hashes": hashes,
	}

	raw, err := client.call("blocks", payload)
	if err != nil {
		return nil, err
	}

	var r map[string]map[string]map[string]string
	if err = json.Unmarshal(raw, &r); err != nil {
		return nil, err
	}

	blocks, ok := r["blocks"]
	if !ok {
		return nil, errors.New("Response of blocks is empty")
	}

	return blocks, nil
}

// Retrieves a json representations of blocks with transaction
// amount & block account.
// Additionally checks if block is pending, returns source account
// for receive & open blocks (0 for send & change blocks) (>= v8.1).
func BlocksInfo(hashes []string, pending, source bool) (map[string]interface{}, error) {
	payload := map[string]interface{}{
		"hashes":  hashes,
		"pending": pending,
		"source":  source,
	}

	return client.fetchMapInterface("blocks_info", payload, "blocks")
}

// Returns the account containing block.
func BlockAccount(hash string) (string, error) {
	payload := map[string]interface{}{
		"hash": hash,
	}

	return client.fetchString("block_account", payload, "account")
}

// Reports the number of blocks in the ledger
// and unchecked synchronizing blocks.
func BlockCount() (map[string]string, error) {
	return client.fetchMap("block_count", nil, "")
}

// Reports the number of blocks in the ledger
// by type (send, receive, open, change).
func BlockCountType() (map[string]string, error) {
	return client.fetchMap("block_count_type", nil, "")
}

// Creates a json representations of a new open block
// based on input data & signed with private key (>= v8.1).
// Optionally uses work value for block from external source.
// Requires enable_control
func CreateOpenBlock(key, account, representative, source, work string) (map[string]string, error) {
	payload := map[string]interface{}{
		"type":           "open",
		"key":            key,
		"account":        account,
		"representative": representative,
		"source":         source,
	}

	if work != "" {
		payload["work"] = work
	}

	return client.fetchMap("block_create", payload, "")
}

// Creates a json representations of a new receive block (>= v8.1).
// Optionally uses work value for block from external source.
// Requires enable_control
func CreateReceiveBlock(wallet, account, source, previous, work string) (map[string]string, error) {
	payload := map[string]interface{}{
		"type":     "receive",
		"wallet":   wallet,
		"account":  account,
		"source":   source,
		"previous": previous,
	}

	if work != "" {
		payload["work"] = work
	}

	return client.fetchMap("block_create", payload, "")
}

// Creates a json representations of a new send block (>= v8.1).
// Optionally uses work value for block from external source.
// Requires enable_control
func CreateSendBlock(wallet, account, destination, balance, amount, previous, work string) (map[string]string, error) {
	payload := map[string]interface{}{
		"type":     "send",
		"wallet":   wallet,
		"account":  account,
		"balance":  balance,
		"amount":   amount,
		"previous": previous,
	}

	if work != "" {
		payload["work"] = work
	}

	return client.fetchMap("block_create", payload, "")
}

// Creates a json representations of a new change block (>= v8.1).
// Optionally uses work value for block from external source.
// Requires enable_control
func CreateChangeBlock(wallet, account, representative, previous, work string) (map[string]string, error) {
	payload := map[string]interface{}{
		"type":           "change",
		"wallet":         wallet,
		"account":        account,
		"representative": representative,
		"previous":       previous,
	}

	if work != "" {
		payload["work"] = work
	}

	return client.fetchMap("block_create", payload, "")
}

// Publishes block to the network.
func ProcessBlock(block map[string]string) (string, error) {
	payload := map[string]interface{}{
		"block": block,
	}

	return client.fetchString("process", payload, "hash")
}

// Checks whether block is pending by hash (>= v8.0).
func PendingExists(hash string) (bool, error) {
	payload := map[string]interface{}{
		"hash": hash,
	}

	return client.isSuccess("pending_exists", payload, "exists")
}

// Retrieves a json representation of unchecked synchronizing block by hash (>= v8.0).
func GetUncheckedBlock(hash string) (string, error) {
	payload := map[string]interface{}{
		"hash": hash,
	}

	return client.fetchString("unchecked_get", payload, "contents")
}

// Stops generating work for block.
// Requires enable_control.
func CancelWork(hash string) error {
	payload := map[string]interface{}{
		"hash": hash,
	}

	_, err := client.call("work_cancel", payload)

	return err
}

// Generates work for block.
// Requires enable_control.
func GenerateWork(hash string) (string, error) {
	payload := map[string]interface{}{
		"hash": hash,
	}

	return client.fetchString("work_generate", payload, "work")
}

// Checks whether work is valid for block.
func ValidateWork(work, hash string) (bool, error) {
	payload := map[string]interface{}{
		"work": work,
		"hash": hash,
	}

	return client.isSuccess("work_validate", payload, "valid")
}

// Returns a list of block hashes in the account
// chain ending at block up to count.
func Successors(block string, count int) ([]string, error) {
	payload := map[string]interface{}{
		"block": block,
		"count": count,
	}

	return client.fetchSlice("successors", payload, "blocks")
}

// Returns a list of block hashes in the account
// chain starting at block up to count.
func Chain(block string, count int) ([]string, error) {
	payload := map[string]interface{}{
		"block": block,
		"count": count,
	}

	return client.fetchSlice("chain", payload, "blocks")
}

// Reports send/receive information for a chain of blocks.
func History(hash string, count int) ([]map[string]string, error) {
	payload := map[string]interface{}{
		"hash":  hash,
		"count": count,
	}

	raw, err := client.call("history", payload)
	if err != nil {
		return nil, err
	}

	var r map[string][]map[string]string
	if err = json.Unmarshal(raw, &r); err != nil {
		return nil, err
	}

	history, ok := r["history"]
	if !ok {
		return nil, errors.New("Response of history is empty")
	}

	return history, nil
}
