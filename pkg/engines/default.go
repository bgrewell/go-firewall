package engines

import (
	"errors"
	"github.com/bgrewell/go-firewall/internal/netlink"
	"log"
	"sync"
)

type DefaultEngine struct {
	Monitor  netlink.NetlinkWatcher
	Triggers <-chan []byte
	Lock     *sync.Mutex
	running  bool
}

func (e *DefaultEngine) CreateTable() {
	//TODO implement me
	panic("implement me")
}

func (e *DefaultEngine) DeleteTable() {
	//TODO implement me
	panic("implement me")
}

func (e *DefaultEngine) CreateChain() {
	//TODO implement me
	panic("implement me")
}

func (e *DefaultEngine) DeleteChain() {
	//TODO implement me
	panic("implement me")
}

func (e *DefaultEngine) CreateRule() {
	//TODO implement me
	panic("implement me")
}

func (e *DefaultEngine) UpdateRule() {
	//TODO implement me
	panic("implement me")
}

func (e *DefaultEngine) DeleteRule() {
	//TODO implement me
	panic("implement me")
}

// Start starts the execution of the rules engine
func (e *DefaultEngine) Start() error {
	e.running = true
	return errors.New("this method has not been implemented")
}

// Stop stops the execution of the rules engine
func (e *DefaultEngine) Stop() error {
	e.running = false
	return errors.New("this method has not been implemented")
}

// triggerRuleSync uses a netlink watcher to monitor the netlink messages and trigger a sync of the rules if
// an NFTABLES message is seen
func (e *DefaultEngine) triggerRuleSync() {
	for e.running {
		// Wait for a trigger from the watcher
		<-e.Triggers

		// Acquire a lock before we modify the rules state, then sync and finally release the lock
		e.Lock.Lock()
		err := e.sync()
		e.Lock.Unlock()
		if err != nil {
			log.Printf("error syncing rules: %v\n", err)
		}
	}
}

// sync synchronizes the internal rule state with the external rule state
func (e *DefaultEngine) sync() error {
	return errors.New("this method is not yet implemented")
}
