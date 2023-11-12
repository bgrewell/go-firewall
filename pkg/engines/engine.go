package engines

import (
	execute "github.com/BGrewell/go-execute/v2"
	"github.com/bgrewell/go-firewall/internal/netlink"
	"sync"
)

func NewEngine() (engine Engine, err error) {
	nlm, err := netlink.NewNetlinkMonitor()
	if err != nil {
		return nil, err
	}

	de := DefaultEngine{
		Monitor:  nlm,
		Triggers: nil,
		Lock:     &sync.Mutex{},
		running:  false,
		exec:     execute.NewExecutor(),
	}

	return &de, nil
}

type Engine interface {
	Start() error
	Stop() error
}
