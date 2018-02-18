package nano

import (
	"github.com/s1na/nano-go/rpc"
)

type Wallet struct {
	Id       string
	Accounts map[string]*Account
	Seed     string
}

func NewWallet() *Wallet {
	w := new(Wallet)
	w.Accounts = make(map[string]*Account)

	return w
}

func (w *Wallet) CreateAccount() (*Account, error) {
	a := NewAccount()

	a.Id, err := rpc.CreateAccount(w.Id, true)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (w *Wallet) Accounts() ([]*Account, error) {
	ids, err := rpc.AccountList(w.Id)
	if err != nil {
		return nil, err
	}

	accounts := make([]*Account, len(ids))
	for i, id := range ids {
		a, ok := w.Accounts[id]
		if !ok {
			a = NewAccount()
			a.Id = id
			w.Accounts[id] = a
		}

		accounts[i] = a
	}

	return accounts
}
