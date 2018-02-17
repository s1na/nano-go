package nano

import (
	"encoding/json"
	"errors"
)

// Returns how many rai are in the public supply.
func AvailableSupply() (string, error) {
	return client.fetchString("available_supply", nil, "available")
}

// Reports the number of accounts in the ledger.
func FrontierCount() (int, error) {
	return client.fetchInt("frontier_count", nil, "count")
}

// Returns a map of representatives and their voting weights.
// If count > 0, limits the number of representatives returned.
// Optionally sorts representatives in descending order.
func Representatives(count int, sort bool) (map[string]string, error) {
	payload := map[string]interface{}{
		"sorting": sort,
	}

	if count > 0 {
		payload["count"] = count
	}

	return client.fetchMap("representatives", nil, "representatives")
}

// Returns frontier, open block, change representative block,
// balance, last modified timestamp from local database and
// block count starting at account up to count (>= v8.1).
// Optionally returns representative, voting weight,
// pending balance for each account.
// Optionally sorts accounts in descending order.
// Requires enable_control.
func Ledger(account string, count int, representative, weight, pending, sorting bool) (map[string]map[string]string, error) {
	payload := map[string]interface{}{
		"account":        account,
		"count":          count,
		"representative": representative,
		"weight":         weight,
		"pending":        pending,
		"sorting":        sorting,
	}

	raw, err := client.call("ledger", payload)
	if err != nil {
		return nil, err
	}

	var r map[string]map[string]map[string]string
	if err = json.Unmarshal(raw, &r); err != nil {
		return nil, err
	}

	accounts, ok := r["accounts"]
	if !ok {
		return nil, errors.New("Response of ledger is empty")
	}

	return accounts, nil
}
