package nano

import (
	"encoding/json"
	"errors"
)

// Returns receive minimum for node (>= v8.0).
// Requires enable_control.
func GetReceiveMinimum() (string, error) {
	return client.fetchString("receive_minimum", nil, "amount")
}

// Sets amount as new receive minimum for node until restart (>= v8.0).
// Returns true if minimum receive was successfully set.
// Requires enable_control.
func SetReceiveMinimum(amount string) (bool, error) {
	payload := map[string]interface{}{
		"amount": amount,
	}

	return client.isSuccess("receive_minimum_set", payload, "")
}

// Tells the node to look for pending blocks for any account in all
// available wallets (>= v8.0).
// Returns true if search started successfully, and false otherwise.
// Requires enable_control.
func SearchAllPending() (bool, error) {
	return client.isSuccess("search_pending_all", nil, "")
}

// Returns a map of unchecked synchronizing block hashes and their json
// representation up to count (>= v8.0).
func UncheckedBlocks(count int) (map[string]map[string]string, error) {
	payload := map[string]interface{}{
		"count": count,
	}

	raw, err := client.call("unchecked", payload)
	if err != nil {
		return nil, err
	}

	var r map[string]map[string]map[string]string
	if err = json.Unmarshal(raw, &r); err != nil {
		return nil, err
	}

	blocks, ok := r["blocks"]
	if !ok {
		return nil, errors.New("Response of unchecked has no blocks")
	}

	return blocks, nil
}

// Clears unchecked synchronizing blocks (>= v8.0).
// Returns true if successfully cleared.
// Requires enable_control
func ClearUncheckedBlocks() (bool, error) {
	return client.isSuccess("unchecked_clear", nil, "")
}

// Tells the node to send a keepalive packet to address:port.
// Requires enable_control.
func SendKeepalive(address string, port int) error {
	payload := map[string]interface{}{
		"address": address,
		"port":    port,
	}

	_, err := client.call("keepalive", payload)

	return err
}

// Returns a map of peer addresses (IPv6:port) and their node network versions.
func Peers() (map[string]string, error) {
	return client.fetchMap("peers", nil, "peers")
}

// Adds a specific IP address and port as work peer for node until restart (>= v8.0).
// Returns true if work peer was added successfully.
// Requires enable_control.
func AddWorkPeer(address, port string) (bool, error) {
	payload := map[string]interface{}{
		"address": address,
		"port":    port,
	}

	return client.isSuccess("work_peer_add", payload, "")
}

// Retrieves work peers (>= v8.0).
// Requires enable_control.
func GetWorkPeers() ([]string, error) {
	return client.fetchSlice("work_peers", nil, "work_peers")
}

// Clears work peers node list until restart (>= v8.0).
// Requires enable_control.
func ClearWorkPeers() (bool, error) {
	return client.isSuccess("work_peers_clear", nil, "")
}

// Initializes bootstrap to specific IP address and port.
// Returns true if bootstrap was started successfully.
func Bootstrap(address string, port int) (bool, error) {
	payload := map[string]interface{}{
		"address": address,
		"port":    port,
	}

	return client.isSuccess("bootstrap", payload, "")
}

// Initialize multi-connection bootstrap to random peers.
// Returns true if bootstrap was started successfully.
func BootstrapAny() (bool, error) {
	return client.isSuccess("bootstrap_any", nil, "")
}

// Rebroadcasts blocks starting at hash to the network.
// If sources > 0, additionally rebroadcast source
// chain blocks for receive/open up to sources depth (>= v8.0).
// If destinations > 0, additionally rebroadcast destination
// chain blocks from receive up to destinations depth (>= v8.0).
func Republish(hash string, count, sources, destinations int) ([]string, error) {
	payload := map[string]interface{}{
		"hash": hash,
	}

	if count > 0 {
		payload["count"] = count
	}

	if sources > 0 {
		payload["sources"] = sources
	}

	if destinations > 0 {
		payload["destinations"] = destinations
	}

	return client.fetchSlice("republish", payload, "blocks")
}

// Returns version information for RPC, Store & Node (Major & Minor version).
// RPC Version always retruns "1" as of 13/01/2018.
func Version() (map[string]string, error) {
	return client.fetchMap("version", nil, "")
}

// Stops the node safely.
func Stop() (bool, error) {
	return client.isSuccess("stop", nil, "")
}
