package rpc

import (
	"encoding/json"
	"errors"
)

type Account struct {
	Frontier            string
	OpenBlock           string
	RepresentativeBlock string
	Balance             string
	Modified            string
	BlockCount          string
	Representative      string
	Weight              string
	Pending             string
}

// Creates a new account, insert next deterministic key in wallet.
// If work is false, it disables work generation after creating account (>= v8.1).
// Requires enable_control.
func CreateAccount(wallet string, work bool) (string, error) {
	payload := map[string]interface{}{
		"wallet": wallet,
		"work":   work,
	}

	return client.fetchString("account_create", payload, "account")
}

// Returns account number corresponding to the public key.
func GetAccount(key string) (string, error) {
	payload := map[string]interface{}{
		"key": key,
	}

	return client.fetchString("account_get", payload, "account")
}

// Returns frontier, open block, change representative block,
// balance, last modified timestamp from local database
// and block count for account.
// Additionally returns representative, voting weight and
// pending balance for account, if respective parameters are set (>= v8.1).
func AccountInfo(account string, representative, weight, pending bool) (map[string]string, error) {
	payload := map[string]interface{}{
		"account":        account,
		"representative": representative,
		"weight":         weight,
		"pending":        pending,
	}

	return client.fetchMap("account_info", payload, "")
}

// Returns how many RAW is owned (balance) and how many
// have not yet been received by account (pending).
func AccountBalance(account string) (string, string, error) {
	payload := map[string]interface{}{
		"account": account,
	}

	r, err := client.fetchMap("account_balance", payload, "")
	if err != nil {
		return "", "", err
	}

	balance, ok := r["balance"]
	if !ok {
		return "", "", errors.New("Response of account_balance has no balance")
	}

	pending, ok := r["pending"]
	if !ok {
		return "", "", errors.New("Response of account_balance has no pending")
	}

	return balance, pending, nil
}

// Returns number of blocks for a specific account.
func AccountBlockCount(account string) (int, error) {
	payload := map[string]interface{}{
		"account": account,
	}

	return client.fetchInt("account_block_count", payload, "block_count")
}

// Reports send/receive information for a account.
func AccountHistory(account string, count int) ([]map[string]string, error) {
	payload := map[string]interface{}{
		"account": account,
		"count":   count,
	}

	raw, err := client.call("account_history", payload)
	if err != nil {
		return nil, err
	}

	var r []map[string]string
	if err = json.Unmarshal(raw, &r); err != nil {
		return nil, err
	}

	return r, nil
}

// Returns the public key for account.
func AccountKey(account string) (string, error) {
	payload := map[string]interface{}{
		"account": account,
	}

	return client.fetchString("account_key", payload, "key")
}

// Returns the representative for account.
func AccountRepresentative(account string) (string, error) {
	payload := map[string]interface{}{
		"account": account,
	}

	return client.fetchString("account_representative", payload, "representative")
}

// Sets the representative for account in wallet.
// If provided, uses work value for block from external source (>= v8.1).
// Returns the change block.
// Requires enable_control.
func SetAccountRepresentative(wallet, account, representative, work string) (string, error) {
	payload := map[string]interface{}{
		"wallet":         wallet,
		"account":        account,
		"representative": representative,
	}

	if work != "" {
		payload["work"] = work
	}

	return client.fetchString("account_representative_set", payload, "block")
}

// Returns the voting weight for account.
func AccountWeight(account string) (string, error) {
	payload := map[string]interface{}{
		"account": account,
	}

	return client.fetchString("account_weight", payload, "weight")
}

// Returns how many RAW is owned and
// how many have not yet been received by accounts list.
func AccountsBalances(accounts []string) (map[string]map[string]string, error) {
	payload := map[string]interface{}{
		"accounts": accounts,
	}

	raw, err := client.call("accounts_balances", payload)
	if err != nil {
		return nil, err
	}

	var r map[string]map[string]map[string]string
	if err = json.Unmarshal(raw, &r); err != nil {
		return nil, err
	}

	balances, ok := r["balances"]
	if !ok {
		return nil, errors.New("Response of accounts_balances is empty")
	}

	return balances, nil
}

// Returns a list of pairs of account and block hash
// representing the head block for accounts list.
func AccountsFrontiers(accounts []string) (map[string]string, error) {
	payload := map[string]interface{}{
		"accounts": accounts,
	}

	return client.fetchMap("accounts_frontiers", payload, "frontiers")
}

// Returns a list of block hashes which have not
// yet been received by these accounts.
// If threshold is not empty, returns a list of pending
// block hashes with amount more or equal to threshold (>= v8.0).
// If source is not empty, returns a list of pending
// block hashes with amount and source accounts (>= v8.1).
func AccountsPending(accounts []string, count int, threshold, source string) (map[string]interface{}, error) {
	payload := map[string]interface{}{
		"accounts": accounts,
		"count":    count,
	}

	if threshold != "" {
		payload["threshold"] = threshold
	}

	if source != "" {
		payload["source"] = source
	}

	return client.fetchMapInterface("accounts_pending", payload, "blocks")
}

// Returns a list of pairs of delegator names given
// account a representative and its balance (>= v8.0).
func Delegators(account string) (map[string]string, error) {
	payload := map[string]interface{}{
		"account": account,
	}

	return client.fetchMap("delegators", payload, "delegators")
}

// Get number of delegators for a specific
// representative account (>= v8.0).
func DelegatorsCount(account string) (int, error) {
	payload := map[string]interface{}{
		"account": account,
	}

	return client.fetchInt("delegators_count", payload, "count")
}

// Returns a list of pairs of account and block hash
// representing the head block starting at account up to count.
func Frontiers(account string, count int) (map[string]string, error) {
	payload := map[string]interface{}{
		"account": account,
		"count":   count,
	}

	return client.fetchMap("frontiers", payload, "frontiers")
}

// Waits for payment of 'amount' to arrive in 'account'
// or until 'timeout' milliseconds have elapsed.
func WaitPayment(account, amount string, timeout int) (string, error) {
	payload := map[string]interface{}{
		"account": account,
		"amount":  amount,
		"timeout": timeout,
	}

	return client.fetchString("payment_wait", payload, "status")
}

// Checks whether account is a valid account number.
func ValidateAccountNumber(account string) (bool, error) {
	payload := map[string]interface{}{
		"account": account,
	}

	return client.isSuccess("validate_account_number", payload, "valid")
}

// Returns a list of block hashes which have not
// yet been received by this account.
// Optionally returns a list of pending block hashes
// with amount more or equal to threshold (>= v8.0).
// Optionally returns a list of pending block hashes
// with amount and source accounts (>= v8.0).
func Pending(account string, count, threshold int, source bool) (interface{}, error) {
	payload := map[string]interface{}{
		"account": account,
		"count":   count,
		"source":  source,
	}

	if threshold > 0 {
		payload["threshold"] = threshold
	}

	return client.fetchInterface("pending", payload, "blocks")
}

// Retrieves work for account in wallet (>= v8.0).
// Requires enable_control.
func GetWork(wallet, account string) (string, error) {
	payload := map[string]interface{}{
		"wallet":  wallet,
		"account": account,
	}

	return client.fetchString("work_get", payload, "work")
}

// Sets work for account in wallet (>= v8.0).
// Requires enable_control.
func SetWork(wallet, account, work string) (bool, error) {
	payload := map[string]interface{}{
		"wallet":  wallet,
		"account": account,
		"work":    work,
	}

	return client.isSuccess("work_set", payload, "")
}
