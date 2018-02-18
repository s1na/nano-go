package nano

import (
	"github.com/s1na/nano-go/rpc"
)

var (
	node *Node
)

type Node struct {
	Wallets map[string]*Wallet
	Version string
}

func NewNode() *Node {
	n := new(Node)
	n.Wallets = make(map[string]*Wallet)

	return n
}

func GetNode() *Node {
	if node != nil {
		return node
	}

	node = NewNode()

	return node
}

func (n *Node) GetWallet(id string) *Wallet {
	w, _ := n.Wallets[id]

	return w
}

func (n *Node) CreateWallet() (*Wallet, error) {
	w := NewWallet()

	w.Id, err := rpc.CreateWallet()
	if err != nil {
		return nil, err
	}

	return w
}

func (n *Node) Version() (string, error) {
	if n.Version != "" {
		return n.Version
	}

	v, err := rpc.Version()
	if err != nil {
		return err
	}

	n.Version = v["node_version"]

	return n.Version
}

func (n *Node) Stop() error {
	return rpc.Stop()
}
