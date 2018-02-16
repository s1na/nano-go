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

// Returns how many RAW is owned and
// how many have not yet been received by accounts list.
func (c *Client) AccountsBalances(accounts []string) (map[string]string, error) {
	payload := map[string]interface{}{
		"accounts": accounts,
	}

	r, err := c.call("accounts_balances", payload)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Creates new accounts, insert next deterministic keys in wallet up to count (>= v8.1).
// Optionally disables work generation after creating account.
// Requires enable_control
func (c *Client) AccountsCreate(wallet string, count int, work bool) (map[string]string, error) {
	payload := map[string]interface{}{
		"wallet": wallet,
		"count":  count,
		"work":   work,
	}

	r, err := c.call("accounts_create", payload)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Returns a list of pairs of account and block hash
// representing the head block for accounts list.
func (c *Client) AccountsFrontiers(accounts []string) (map[string]string, error) {
	payload := map[string]interface{}{
		"accounts": accounts,
	}

	r, err := c.call("accounts_frontiers", payload)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Returns a list of block hashes which have not
// yet been received by these accounts.
// If threshold is not empty, returns a list of pending
// block hashes with amount more or equal to threshold (>= v8.0).
// If source is not empty, returns a list of pending
// block hashes with amount and source accounts (>= v8.1).
func (c *Client) AccountsPending(accounts []string, count int, threshold, source string) (map[string]string, error) {
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

	r, err := c.call("accounts_pending", payload)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Returns how many rai are in the public supply.
func (c *Client) AvailableSupply() (map[string]string, error) {
	r, err := c.call("accounts_pending", nil)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Retrieves a json representation of block.
func (c *Client) Block(hash string) (map[string]string, error) {
	payload := map[string]interface{}{
		"hash": hash,
	}

	r, err := c.call("block", payload)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Retrieves a json representations of blocks.
func (c *Client) Blocks(hashes []string) (map[string]string, error) {
	payload := map[string]interface{}{
		"hashes": hashes,
	}

	r, err := c.call("blocks", payload)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Retrieves a json representations of blocks with transaction
// amount & block account.
// Additionally checks if block is pending, returns source account
// for receive & open blocks (0 for send & change blocks) (>= v8.1).
func (c *Client) BlocksInfo(hashes []string, pending, source bool) (map[string]string, error) {
	payload := map[string]interface{}{
		"hashes":  hashes,
		"pending": pending,
		"source":  source,
	}

	r, err := c.call("blocks_info", payload)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Returns the account containing block.
func (c *Client) BlockAccount(hash string) (map[string]string, error) {
	payload := map[string]interface{}{
		"hash": hash,
	}

	r, err := c.call("block_account", payload)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Reports the number of blocks in the ledger
// and unchecked synchronizing blocks.
func (c *Client) BlockCount() (map[string]string, error) {
	r, err := c.call("block_count", nil)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Reports the number of blocks in the ledger
// by type (send, receive, open, change).
func (c *Client) BlockCountType() (map[string]string, error) {
	r, err := c.call("block_count_type", nil)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Initializes bootstrap to specific IP address and port.
func (c *Client) Bootstrap(address string, port int) (map[string]string, error) {
	payload := map[string]interface{}{
		"address": address,
		"port":    port,
	}

	r, err := c.call("bootstrap", payload)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Initialize multi-connection bootstrap to random peers.
func (c *Client) BootstrapAny() (map[string]string, error) {
	r, err := c.call("bootstrap_any", nil)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Returns a list of block hashes in the account
// chain starting at block up to count.
func (c *Client) Chain(block string, count int) (map[string]string, error) {
	payload := map[string]interface{}{
		"block": block,
		"count": count,
	}

	r, err := c.call("chain", payload)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Returns a list of pairs of delegator names given
// account a representative and its balance (>= v8.0).
func (c *Client) Delegators(account string) (map[string]string, error) {
	payload := map[string]interface{}{
		"account": account,
	}

	r, err := c.call("delegators", payload)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Get number of delegators for a specific
// representative account (>= v8.0).
func (c *Client) DelegatorsCount(account string) (map[string]string, error) {
	payload := map[string]interface{}{
		"account": account,
	}

	r, err := c.call("delegators_count", payload)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Derive deterministic keypair from seed based on index.
func (c *Client) DeterministicKey(seed string, index int) (map[string]string, error) {
	payload := map[string]interface{}{
		"seed":  seed,
		"index": index,
	}

	r, err := c.call("deterministic_key", payload)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Returns a list of pairs of account and block hash
// representing the head block starting at account up to count.
func (c *Client) Frontiers(account string, count int) (map[string]string, error) {
	payload := map[string]interface{}{
		"account": account,
		"count":   count,
	}

	r, err := c.call("frontiers", payload)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Reports the number of accounts in the ledger.
func (c *Client) FrontierCount() (map[string]string, error) {
	r, err := c.call("frontier_count", nil)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Reports send/receive information for a chain of blocks.
func (c *Client) History(hash string, count int) (map[string]string, error) {
	payload := map[string]interface{}{
		"hash":  hash,
		"count": count,
	}

	r, err := c.call("history", payload)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Tells the node to send a keepalive packet to address:port.
// Requires enable_control.
func (c *Client) Keepalive(address string, port int) (map[string]string, error) {
	payload := map[string]interface{}{
		"address": address,
		"port":    port,
	}

	r, err := c.call("keepalive", payload)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Generates an adhoc random keypair.
func (c *Client) KeyCreate() (map[string]string, error) {
	r, err := c.call("key_create", nil)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Derives public key and account number from private key.
func (c *Client) KeyExpand(key string) (map[string]string, error) {
	payload := map[string]interface{}{
		"key": key,
	}

	r, err := c.call("key_expand", payload)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Returns frontier, open block, change representative block,
// balance, last modified timestamp from local database and
// block count starting at account up to count (>= v8.1).
// Optionally returns representative, voting weight,
// pending balance for each account.
// Optionally sorts accounts in descending order.
// Requires enable_control.
func (c *Client) Ledger(account string, count int, representative, weight, pending, sorting bool) (map[string]string, error) {
	payload := map[string]interface{}{
		"account":        account,
		"count":          count,
		"representative": representative,
		"weight":         weight,
		"pending":        pending,
		"sorting":        sorting,
	}

	r, err := c.call("ledger", payload)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Creates a json representations of a new open block
// based on input data & signed with private key (>= v8.1).
// Optionally uses work value for block from external source.
// Requires enable_control
func (c *Client) CreateOpenBlock(key, account, representative, source, work string) (map[string]string, error) {
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

	r, err := c.call("block_create", payload)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Creates a json representations of a new receive block (>= v8.1).
// Optionally uses work value for block from external source.
// Requires enable_control
func (c *Client) CreateReceiveBlock(wallet, account, source, previous, work string) (map[string]string, error) {
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

	r, err := c.call("block_create", payload)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Creates a json representations of a new send block (>= v8.1).
// Optionally uses work value for block from external source.
// Requires enable_control
func (c *Client) CreateSendBlock(wallet, account, destination, balance, amount, previous, work string) (map[string]string, error) {
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

	r, err := c.call("block_create", payload)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Creates a json representations of a new change block (>= v8.1).
// Optionally uses work value for block from external source.
// Requires enable_control
func (c *Client) CreateChangeBlock(wallet, account, representative, previous, work string) (map[string]string, error) {
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

	r, err := c.call("block_create", payload)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Begins a new payment session. Searches wallet for an account that's marked
// as available and has a 0 balance. If one is found, the account number
// is returned and is marked as unavailable. If no account is found,
// a new account is created, placed in the wallet, and returned.
func (c *Client) BeginPayment(wallet string) (map[string]string, error) {
	payload := map[string]interface{}{
		"wallet": wallet,
	}

	r, err := c.call("payment_begin", payload)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Marks all accounts in wallet as available for being used as a payment session.
func (c *Client) InitPayment(wallet string) (map[string]string, error) {
	payload := map[string]interface{}{
		"wallet": wallet,
	}

	r, err := c.call("payment_init", payload)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Ends a payment session. Marks the account as available for use in a payment session.
func (c *Client) EndPayment(wallet, account string) (map[string]string, error) {
	payload := map[string]interface{}{
		"wallet":  wallet,
		"account": account,
	}

	r, err := c.call("payment_end", payload)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Waits for payment of 'amount' to arrive in 'account'
// or until 'timeout' milliseconds have elapsed.
func (c *Client) WaitPayment(account string, amount, timeout int) (map[string]string, error) {
	payload := map[string]interface{}{
		"account": account,
		"amount":  amount,
		"timeout": timeout,
	}

	r, err := c.call("payment_wait", payload)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Publishes block to the network.
func (c *Client) ProcessBlock(block map[string]string) (map[string]string, error) {
	payload := map[string]interface{}{
		"block": block,
	}

	r, err := c.call("process", payload)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Receives pending block for account in wallet.
// Optionally Uses work value for block from external source (>= v8.1).
func (c *Client) ReceiveBlock(wallet, account, block, work string) (map[string]string, error) {
	payload := map[string]interface{}{
		"wallet":  wallet,
		"account": account,
		"block":   block,
	}

	if work != "" {
		payload["work"] = work
	}

	r, err := c.call("receive", payload)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Returns receive minimum for node (>= v8.0).
// Requires enable_control.
func (c *Client) ReceiveMinimum() (map[string]string, error) {
	r, err := c.call("receive_minimum", nil)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Sets amount as new receive minimum for node until restart (>= v8.0).
// Requires enable_control.
func (c *Client) SetReceiveMinimum(amount int) (map[string]string, error) {
	payload := map[string]interface{}{
		"amount": amount,
	}

	r, err := c.call("receive_minimum_set", payload)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Returns a list of pairs of representative and its voting weight.
// Optionally (if count > 0) returns a list of pairs of representative
// and its voting weight up to count.
// Optionally sorts representatives in descending order.
func (c *Client) Representatives(count int, sort bool) (map[string]string, error) {
	payload := map[string]interface{}{
		"sorting": sort,
	}

	if count > 0 {
		payload["count"] = count
	}

	r, err := c.call("representatives", nil)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Returns the default representative for wallet.
func (c *Client) WalletRepresentative(wallet string) (map[string]string, error) {
	payload := map[string]interface{}{
		"wallet": wallet,
	}

	r, err := c.call("wallet_representative", payload)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Sets the default representative for wallet.
// Requires enable_control.
func (c *Client) SetWalletRepresentative(wallet, representative string) (map[string]string, error) {
	payload := map[string]interface{}{
		"wallet":         wallet,
		"representative": representative,
	}

	r, err := c.call("wallet_representative_set", payload)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Rebroadcasts blocks starting at hash to the network.
// If sources > 0, additionally rebroadcast source
// chain blocks for receive/open up to sources depth (>= v8.0).
// If destinations > 0, additionally rebroadcast destination
// chain blocks from receive up to destinations depth (>= v8.0).
func (c *Client) Republish(hash string, count, sources, destinations int) (map[string]string, error) {
	payload := map[string]interface{}{
		"hash": hash,
	}

	if sources > 0 {
		payload["count"] = count
		payload["sources"] = sources
	}

	if destinations > 0 {
		payload["count"] = count
		payload["destinations"] = destinations
	}

	r, err := c.call("republish", payload)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Tells the node to look for pending blocks for any account in wallet.
// Requires enable_control.
func (c *Client) SearchPending(wallet string) (map[string]string, error) {
	payload := map[string]interface{}{
		"wallet": wallet,
	}

	r, err := c.call("search_pending", payload)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Tells the node to look for pending blocks for
// any account in all available wallets (>= v8.0).
// Requires enable_control.
func (c *Client) SearchPendingAll() (map[string]string, error) {
	r, err := c.call("search_pending_all", nil)
	if err != nil {
		return nil, err
	}

	return r, nil
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
func (c *Client) Send(wallet, source, destination, id string, amount int, work string) (map[string]string, error) {
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

	r, err := c.call("send", payload)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Checks whether account is a valid account number.
func (c *Client) ValidateAccountNumber(account string) (map[string]string, error) {
	payload := map[string]interface{}{
		"account": account,
	}

	r, err := c.call("validate_account_number", payload)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Returns a list of block hashes in the account
// chain ending at block up to count.
func (c *Client) Successors(block string, count int) (map[string]string, error) {
	payload := map[string]interface{}{
		"block": block,
		"count": count,
	}

	r, err := c.call("successors", payload)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Returns a list of pairs of peer IPv6:port and its node network version.
func (c *Client) Peers() (map[string]string, error) {
	r, err := c.call("peers", nil)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Returns a list of block hashes which have not
// yet been received by this account.
// Optionally returns a list of pending block hashes
// with amount more or equal to threshold (>= v8.0).
// Optionally returns a list of pending block hashes
// with amount and source accounts (>= v8.0).
func (c *Client) Pending(account string, count, threshold int, source bool) (map[string]string, error) {
	payload := map[string]interface{}{
		"account": account,
		"count":   count,
		"source":  source,
	}

	if threshold > 0 {
		payload["threshold"] = threshold
	}

	r, err := c.call("pending", payload)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Checks whether block is pending by hash (>= v8.0).
func (c *Client) PendingExists(hash string) (map[string]string, error) {
	payload := map[string]interface{}{
		"hash": hash,
	}

	r, err := c.call("pending_exists", payload)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Returns a list of pairs of unchecked synchronizing block
// hash and its json representation up to count (>= v8.0).
func (c *Client) UncheckedBlocks(count int) (map[string]string, error) {
	payload := map[string]interface{}{
		"count": count,
	}

	r, err := c.call("unchecked", payload)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Clears unchecked synchronizing blocks (>= v8.0).
// Requires enable_control
func (c *Client) ClearUncheckedBlocks() (map[string]string, error) {
	r, err := c.call("unchecked_clear", nil)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Retrieves a json representation of unchecked synchronizing block by hash (>= v8.0).
func (c *Client) GetUncheckedBlock(hash string) (map[string]string, error) {
	payload := map[string]interface{}{
		"hash": hash,
	}

	r, err := c.call("unchecked_get", payload)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Retrieves unchecked database keys, blocks hashes & a json
// representations of unchecked pending blocks
// starting from key up to count (>= v8.0).
func (c *Client) UncheckedKeys(key string, count int) (map[string]string, error) {
	payload := map[string]interface{}{
		"key":   key,
		"count": count,
	}

	r, err := c.call("unchecked_keys", payload)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Adds an adhoc private key key to wallet.
// Optionally disables work generation after adding account (>= v8.1).
// Requires enable_control.
func (c *Client) WalletAdd(wallet, key string, work bool) (map[string]string, error) {
	payload := map[string]interface{}{
		"wallet": wallet,
		"key":    key,
		"work":   work,
	}

	r, err := c.call("wallet_add", payload)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Returns the sum of all accounts balances in wallet.
func (c *Client) WalletTotalBalance(wallet string) (map[string]string, error) {
	payload := map[string]interface{}{
		"wallet": wallet,
	}

	r, err := c.call("wallet_balance_total", payload)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Returns how many rai is owned and how many have not
// yet been received by all accounts in wallet.
// If threshold > 0, returns wallet accounts balances more or equal to threshold (>= v8.1).
func (c *Client) WalletBalances(wallet string, threshold int) (map[string]string, error) {
	payload := map[string]interface{}{
		"wallet": wallet,
	}

	if threshold > 0 {
		payload["threshold"] = threshold
	}

	r, err := c.call("wallet_balances", payload)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Changes seed for wallet to seed.
// Requires enable_control.
func (c *Client) ChangeWalletSeed(wallet, seed string) (map[string]string, error) {
	payload := map[string]interface{}{
		"wallet": wallet,
		"seed":   seed,
	}

	r, err := c.call("wallet_change_seed", payload)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Checks whether wallet contains account.
func (c *Client) WalletContains(wallet, account string) (map[string]string, error) {
	payload := map[string]interface{}{
		"wallet":  wallet,
		"account": account,
	}

	r, err := c.call("wallet_contains", payload)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Creates a new random wallet id.
// Requires enable_control.
func (c *Client) CreateWallet() (map[string]string, error) {
	r, err := c.call("wallet_create", nil)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Destroys wallet and all contained accounts.
// Requires enable_control.
func (c *Client) DestroyWallet(wallet string) (map[string]string, error) {
	payload := map[string]interface{}{
		"wallet": wallet,
	}

	r, err := c.call("wallet_destroy", payload)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Returns a json representation of wallet.
func (c *Client) ExportWallet(wallet string) (map[string]string, error) {
	payload := map[string]interface{}{
		"wallet": wallet,
	}

	r, err := c.call("wallet_export", payload)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Returns a list of pairs of account and block hash representing
// the head block starting for accounts from wallet.
func (c *Client) WalletFrontiers(wallet string) (map[string]string, error) {
	payload := map[string]interface{}{
		"wallet": wallet,
	}

	r, err := c.call("wallet_frontiers", payload)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Returns a list of block hashes which have not yet been
// received by accounts in this wallet (>= v8.0).
// If threshold > 0, Returns a list of pending block hashes
// with amount more or equal to threshold.
// Optionally, Returns a list of pending block hashes with
// amount and source accounts (>= v8.1).
// Requires enable_control.
func (c *Client) WalletPending(wallet string, count, threshold int, source bool) (map[string]string, error) {
	payload := map[string]interface{}{
		"wallet": wallet,
		"count":  count,
		"source": source,
	}

	if threshold > 0 {
		payload["threshold"] = threshold
	}

	r, err := c.call("wallet_pending", payload)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Rebroadcasts blocks for accounts from wallet starting
// at frontier down to count to the network (>= v8.0).
// Requires enable_control
func (c *Client) WalletRepublish(wallet string, count int) (map[string]string, error) {
	payload := map[string]interface{}{
		"wallet": wallet,
		"count":  count,
	}

	r, err := c.call("wallet_republish", payload)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Returns a list of pairs of account and work from wallet (>= v8.0).
// Requires enable_control
func (c *Client) WalletWorkGet(wallet string) (map[string]string, error) {
	payload := map[string]interface{}{
		"wallet": wallet,
	}

	r, err := c.call("wallet_work_get", payload)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Changes the password for wallet to password.
// Requires enable_control
func (c *Client) ChangeWalletPassword(wallet, password string) (map[string]string, error) {
	payload := map[string]interface{}{
		"wallet":   wallet,
		"password": password,
	}

	r, err := c.call("password_change", payload)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Enters the password in to wallet.
func (c *Client) EnterWalletPassword(wallet, password string) (map[string]string, error) {
	payload := map[string]interface{}{
		"wallet":   wallet,
		"password": password,
	}

	r, err := c.call("password_enter", payload)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Checks whether the password entered for wallet is valid.
func (c *Client) WalletPasswordValid(wallet, password string) (map[string]string, error) {
	payload := map[string]interface{}{
		"wallet":   wallet,
		"password": password,
	}

	r, err := c.call("password_valid", payload)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Checks whether wallet is locked.
func (c *Client) IsWalletLocked(wallet string) (map[string]string, error) {
	payload := map[string]interface{}{
		"wallet": wallet,
	}

	r, err := c.call("password_locked", payload)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Stops generating work for block.
// Requires enable_control.
func (c *Client) CancelWork(hash string) (map[string]string, error) {
	payload := map[string]interface{}{
		"hash": hash,
	}

	r, err := c.call("work_cancel", payload)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Generates work for block.
// Requires enable_control.
func (c *Client) GenerateWork(hash string) (map[string]string, error) {
	payload := map[string]interface{}{
		"hash": hash,
	}

	r, err := c.call("work_generate", payload)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Retrieves work for account in wallet (>= v8.0).
// Requires enable_control.
func (c *Client) GetWork(wallet, account string) (map[string]string, error) {
	payload := map[string]interface{}{
		"wallet":  wallet,
		"account": account,
	}

	r, err := c.call("work_get", payload)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Sets work for account in wallet (>= v8.0).
// Requires enable_control.
func (c *Client) SetWork(wallet, account, work string) (map[string]string, error) {
	payload := map[string]interface{}{
		"wallet":  wallet,
		"account": account,
		"work":    work,
	}

	r, err := c.call("work_set", payload)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Adds specific IP address and port as work peer
// for node until restart (>= v8.0).
// Requires enable_control.
func (c *Client) AddWorkPeer(address string, port int) (map[string]string, error) {
	payload := map[string]interface{}{
		"address": address,
		"port":    port,
	}

	r, err := c.call("work_peer_add", payload)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Retrieve work peers (>= v8.0).
// Requires enable_control.
func (c *Client) GetWorkPeers() (map[string]string, error) {
	r, err := c.call("work_peers", nil)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Clear work peers node list until restart (>= v8.0).
// Requires enable_control.
func (c *Client) ClearWorkPeers() (map[string]string, error) {
	r, err := c.call("work_peers_clear", nil)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Checks whether work is valid for block.
func (c *Client) ValidateWork(work, hash string) (map[string]string, error) {
	payload := map[string]interface{}{
		"work": work,
		"hash": hash,
	}

	r, err := c.call("work_validate", payload)
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
