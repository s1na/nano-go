package rpc

import (
	"encoding/json"
	"errors"
)

// Lists all the accounts inside wallet.
func AccountList(wallet string) ([]string, error) {
	payload := map[string]interface{}{
		"wallet": wallet,
	}

	return client.fetchSlice("account_list", payload, "accounts")
}

// Moves accounts from source to wallet.
// Returns true if accounts were moved successfully.
// Requires enable_control.
func MoveAccounts(wallet, source string, accounts []string) (bool, error) {
	payload := map[string]interface{}{
		"wallet":   wallet,
		"source":   source,
		"accounts": accounts,
	}

	return client.isSuccess("account_move", payload, "moved")
}

// Removes account from wallet.
// Returns true if account was removed successfully.
// Requires enable_control.
func RemoveAccount(wallet, account string) (bool, error) {
	payload := map[string]interface{}{
		"wallet":  wallet,
		"account": account,
	}

	return client.isSuccess("account_remove", payload, "removed")
}

// Creates new accounts, insert next deterministic keys in wallet up to count (>= v8.1).
// Optionally disables work generation after creating account.
// Requires enable_control
func CreateAccounts(wallet string, count int, work bool) ([]string, error) {
	payload := map[string]interface{}{
		"wallet": wallet,
		"count":  count,
		"work":   work,
	}

	return client.fetchSlice("accounts_create", payload, "accounts")
}

// Begins a new payment session. Searches wallet for an account that's marked
// as available and has a 0 balance. If one is found, the account number
// is returned and is marked as unavailable. If no account is found,
// a new account is created, placed in the wallet, and returned.
func BeginPayment(wallet string) (string, error) {
	payload := map[string]interface{}{
		"wallet": wallet,
	}

	return client.fetchString("payment_begin", payload, "account")
}

// Marks all accounts in wallet as available for being used as a payment session.
// Returns status.
func InitPayment(wallet string) (string, error) {
	payload := map[string]interface{}{
		"wallet": wallet,
	}

	return client.fetchString("payment_init", payload, "status")
}

// Ends a payment session. Marks the account as available for use in a payment session.
func EndPayment(wallet, account string) error {
	payload := map[string]interface{}{
		"wallet":  wallet,
		"account": account,
	}

	_, err := client.call("payment_end", payload)

	return err
}

// Receives pending block for account in wallet.
// Optionally Uses work value for block from external source (>= v8.1).
func ReceiveBlock(wallet, account, block, work string) (string, error) {
	payload := map[string]interface{}{
		"wallet":  wallet,
		"account": account,
		"block":   block,
	}

	if work != "" {
		payload["work"] = work
	}

	return client.fetchString("receive", payload, "block")
}

// Returns the default representative for wallet.
func WalletRepresentative(wallet string) (string, error) {
	payload := map[string]interface{}{
		"wallet": wallet,
	}

	return client.fetchString("wallet_representative", payload, "representative")
}

// Sets the default representative for wallet.
// Requires enable_control.
func SetWalletRepresentative(wallet, representative string) (bool, error) {
	payload := map[string]interface{}{
		"wallet":         wallet,
		"representative": representative,
	}

	return client.isSuccess("wallet_representative_set", payload, "set")
}

// Tells the node to look for pending blocks for any account in wallet.
// Requires enable_control.
func SearchPending(wallet string) (bool, error) {
	payload := map[string]interface{}{
		"wallet": wallet,
	}

	return client.isSuccess("search_pending", payload, "started")
}

// Send amount from source in wallet to destination.
// Proof of Work is precomputed for one transaction in the background.
// If it has been a while since your last transaction it will send
// instantly, the next one will need to wait for Proof of Work to be generated.
// A unique id (per node) should be specified for each spend to provide idempotency.
// That means that if you call send two times with the same id, the second request
// won't send any additional Nano, and will return the first block instead (>= 10.0).
// Using the same id for requests with different parameters
// (wallet, source, destination, and amount) is undefined behavior
// and may result in an error in the future.
// Optionally uses work value for block from external source (>= v8.1).
// Requires enable_control.
func Send(wallet, source, destination, id string, amount int, work string) (string, error) {
	payload := map[string]interface{}{
		"wallet":      wallet,
		"source":      source,
		"destination": destination,
		"id":          id,
		"amount":      amount,
	}

	if work != "" {
		payload["work"] = work
	}

	return client.fetchString("send", payload, "block")
}

// Adds an adhoc private key key to wallet.
// Optionally disables work generation after adding account (>= v8.1).
// Requires enable_control.
func WalletAdd(wallet, key string, work bool) (string, error) {
	payload := map[string]interface{}{
		"wallet": wallet,
		"key":    key,
		"work":   work,
	}

	return client.fetchString("wallet_add", payload, "key")
}

// Returns the sum of all accounts balances in wallet.
func WalletTotalBalance(wallet string) (map[string]string, error) {
	payload := map[string]interface{}{
		"wallet": wallet,
	}

	return client.fetchMap("wallet_balance_total", payload, "")
}

// Returns how many rai is owned and how many have not
// yet been received by all accounts in wallet.
// If threshold > 0, returns wallet accounts balances more or equal to threshold (>= v8.1).
func WalletBalances(wallet string, threshold int) (map[string]map[string]string, error) {
	payload := map[string]interface{}{
		"wallet": wallet,
	}

	if threshold > 0 {
		payload["threshold"] = threshold
	}

	raw, err := client.call("wallet_balances", payload)
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

// Changes seed for wallet to seed.
// Requires enable_control.
func ChangeWalletSeed(wallet, seed string) (bool, error) {
	payload := map[string]interface{}{
		"wallet": wallet,
		"seed":   seed,
	}

	return client.isSuccess("wallet_change_seed", payload, "")
}

// Checks whether wallet contains account.
func WalletContains(wallet, account string) (bool, error) {
	payload := map[string]interface{}{
		"wallet":  wallet,
		"account": account,
	}

	return client.isSuccess("wallet_contains", payload, "exists")
}

// Creates a new random wallet id.
// Requires enable_control.
func CreateWallet() (string, error) {
	return client.fetchString("wallet_create", nil, "wallet")
}

// Destroys wallet and all contained accounts.
// Requires enable_control.
func DestroyWallet(wallet string) error {
	payload := map[string]interface{}{
		"wallet": wallet,
	}

	_, err := client.call("wallet_destroy", payload)

	return err
}

// Returns a json representation of wallet.
func ExportWallet(wallet string) (string, error) {
	payload := map[string]interface{}{
		"wallet": wallet,
	}

	return client.fetchString("wallet_export", payload, "json")
}

// Returns a list of pairs of account and block hash representing
// the head block starting for accounts from wallet.
func WalletFrontiers(wallet string) (map[string]string, error) {
	payload := map[string]interface{}{
		"wallet": wallet,
	}

	return client.fetchMap("wallet_frontiers", payload, "frontiers")
}

// Returns a list of block hashes which have not yet been
// received by accounts in this wallet (>= v8.0).
// If threshold > 0, Returns a list of pending block hashes
// with amount more or equal to threshold.
// Optionally, Returns a list of pending block hashes with
// amount and source accounts (>= v8.1).
// Requires enable_control.
func WalletPending(wallet string, count, threshold int, source bool) (map[string]interface{}, error) {
	payload := map[string]interface{}{
		"wallet": wallet,
		"count":  count,
		"source": source,
	}

	if threshold > 0 {
		payload["threshold"] = threshold
	}

	return client.fetchMapInterface("wallet_pending", payload, "blocks")
}

// Rebroadcasts blocks for accounts from wallet starting
// at frontier down to count to the network (>= v8.0).
// Requires enable_control
func WalletRepublish(wallet string, count int) ([]string, error) {
	payload := map[string]interface{}{
		"wallet": wallet,
		"count":  count,
	}

	return client.fetchSlice("wallet_republish", payload, "blocks")
}

// Returns a map of account and work from wallet (>= v8.0).
// Requires enable_control
func WalletWorkGet(wallet string) (map[string]string, error) {
	payload := map[string]interface{}{
		"wallet": wallet,
	}

	return client.fetchMap("wallet_work_get", payload, "works")
}

// Changes the password for wallet to password.
// Requires enable_control
func ChangeWalletPassword(wallet, password string) (bool, error) {
	payload := map[string]interface{}{
		"wallet":   wallet,
		"password": password,
	}

	return client.isSuccess("password_change", payload, "changed")
}

// Enters the password in to wallet.
func EnterWalletPassword(wallet, password string) (bool, error) {
	payload := map[string]interface{}{
		"wallet":   wallet,
		"password": password,
	}

	return client.isSuccess("password_enter", payload, "valid")
}

// Checks whether the password entered for wallet is valid.
func WalletPasswordValid(wallet, password string) (bool, error) {
	payload := map[string]interface{}{
		"wallet":   wallet,
		"password": password,
	}

	return client.isSuccess("password_valid", payload, "valid")
}

// Checks whether wallet is locked.
func IsWalletLocked(wallet string) (bool, error) {
	payload := map[string]interface{}{
		"wallet": wallet,
	}

	return client.isSuccess("password_locked", payload, "locked")
}
