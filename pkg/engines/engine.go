package engines

import (
	execute "github.com/BGrewell/go-execute/v2"
	"github.com/bgrewell/go-firewall/internal/netlink"
	"sync"
)

// EngineOptions defines a generic function that takes and modifies the listOptions struct
type EngineOptions func(options *engineOptions)

// engineOptions is a struct that contains the various options for setting up a firewall engine
type engineOptions struct {
	monitor    netlink.NetlinkWatcher
	notifyChan chan<- struct{}
}

// WithMonitor adds a Netlink Monitor to the engineOptions. This allows the engine to monitor changes
// to the nftables configuration in realtime which allows it to update the rule state and send notifications.
func WithMonitor(monitor netlink.NetlinkWatcher) EngineOptions {
	return func(opts *engineOptions) {
		opts.monitor = monitor
	}
}

// WithChangeNotify adds a channel that the engine can send update triggers on. For this to work you must provide a monitor
// in addition to the trigger as this will only trigger on external changes, not changes that are made via this library.
func WithChangeNotify(notifyChan chan<- struct{}) EngineOptions {
	return func(opts *engineOptions) {
		opts.notifyChan = notifyChan
	}
}

func NewEngine(options ...EngineOptions) (engine Engine, err error) {
	opts := &engineOptions{}

	for _, option := range options {
		option(opts)
	}

	de := DefaultEngine{
		Monitor:    opts.monitor,
		Lock:       &sync.Mutex{},
		running:    false,
		exec:       execute.NewExecutor(),
		notifyChan: opts.notifyChan,
	}

	return &de, nil
}

type Engine interface {
	Start() error
	Stop() error
}
