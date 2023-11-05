package iptables

import (
	"github.com/bgrewell/go-firewall/internal/netlink"
	"github.com/bgrewell/go-firewall/pkg/engines"
	"sync"
)

func NewIptablesEngine() (iptEngine Engine, err error) {
	nlm, _ := netlink.NewNetlinkMonitor()
	triggers, _, setupErr := nlm.Monitor()
	if setupErr != nil {
		return nil, err
	}

	e := &engines.DefaultEngine{
		Monitor:  nlm,
		Triggers: triggers,
		Lock:     &sync.Mutex{},
	}

	return e, nil
}

type Engine interface {
	Start() error
	Stop() error
	CreateTable()
	DeleteTable()
	CreateChain()
	DeleteChain()
	CreateRule()
	UpdateRule()
	DeleteRule()
}
