package engines

import (
	"errors"
	execute "github.com/BGrewell/go-execute/v2"
	"github.com/bgrewell/go-firewall/internal/netlink"
	"github.com/bgrewell/go-firewall/pkg/nftables"
	"log"
	"sync"
	"time"
)

type DefaultEngine struct {
	Monitor     netlink.NetlinkWatcher
	Errors      <-chan error
	Lock        *sync.Mutex
	exec        execute.Executor
	running     bool
	ruleset     *nftables.Ruleset
	notifyChan  chan<- struct{} // Channel used to update the caller that a nftable change was detected
	monitorChan <-chan []byte   // Channel returned by the monitor that it triggers nftables changes on
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
	// Mark as running
	e.running = true

	// Setup sync triggering
	var err error
	e.monitorChan, e.Errors, err = e.Monitor.Monitor()
	if err != nil {
		return err
	}

	// Manually trigger once then monitor for triggers
	err = e.sync()
	if err != nil {
		return err
	}
	go e.monitorSyncTrigger()
	return nil
}

// Stop stops the execution of the rules engine
func (e *DefaultEngine) Stop() error {
	e.running = false
	return errors.New("this method has not been implemented")
}

// monitorSyncTrigger uses a netlink watcher to monitor the netlink messages and trigger a sync of the rules if
// an NFTABLES message is seen
func (e *DefaultEngine) monitorSyncTrigger() {
	for e.running {
		// Wait for a trigger from the watcher
		<-e.monitorChan

		// Acquire a lock before we modify the rules state, then sync and finally release the lock
		e.Lock.Lock()
		err := e.sync()
		e.Lock.Unlock()
		if err != nil {
			log.Printf("error syncing rules: %v\n", err)
		}
		if e.notifyChan != nil {
			e.notifyChan <- struct{}{}
		}
	}
}

// sync synchronizes the internal rule state with the external rule state
func (e *DefaultEngine) sync() error {
	output, err := e.exec.ExecuteWithTimeout("nft -aj list ruleset", 10*time.Second)
	if err != nil {
		return err
	}
	e.ruleset, err = nftables.UnmarshalNftables(output)
	if err != nil {
		return err
	}

	return nil
}
