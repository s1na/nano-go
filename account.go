package nano

import (
	"github.com/s1na/nano-go/rpc"
)

type Account struct {
	Id string
}

func NewAccount() *Account {
	a := new(Account)

	return a
}
