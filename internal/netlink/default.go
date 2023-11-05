package netlink

import (
	"golang.org/x/sys/unix"
)

type DefaultNetlinkWatcher struct {
	monitorSocketFd int
	running         bool
	eventsChan      chan []byte
	errorsChan      chan error
}

func (w *DefaultNetlinkWatcher) Monitor() (events <-chan []byte, errors <-chan error, setupErr error) {
	// Open up a netlink socket
	w.monitorSocketFd, setupErr = unix.Socket(unix.AF_NETLINK, unix.SOCK_RAW, unix.NETLINK_NETFILTER)
	if setupErr != nil {
		return nil, nil, setupErr
	}

	// Setup socket address
	sa := &unix.SockaddrNetlink{
		Family: unix.AF_NETLINK,
		Pid:    uint32(0),
		Groups: NFNLGRP_NFTABLES,
	}

	// Bind to socket
	if setupErr = unix.Bind(w.monitorSocketFd, sa); setupErr != nil {
		return nil, nil, setupErr
	}

	// Setup reporting channels
	w.running = true
	w.eventsChan = make(chan []byte, 4096)
	w.errorsChan = make(chan error, 4096)
	go w.processMessages()

	return w.eventsChan, w.errorsChan, nil
}

func (w *DefaultNetlinkWatcher) Stop() error {
	w.running = false
	return unix.Close(w.monitorSocketFd)
}

func (w *DefaultNetlinkWatcher) Running() bool {
	return w.running
}

func (w *DefaultNetlinkWatcher) processMessages() {
	for w.running {
		buf := make([]byte, 4096)
		n, _, err := unix.Recvfrom(w.monitorSocketFd, buf, 0)
		if err != nil {
			// Handle the case when the socket is closed (e.g., via w.Stop())
			if err == unix.EBADF { // This is the error returned on a read from a closed file descriptor
				break // Exit the loop
			}
			w.errorsChan <- err
			continue
		}
		if n > 0 {
			// Technically we could just return a bool as we don't expect the contents to be used as this
			// is more of a trigger to notify the wrapper that the firewall rules have changed
			w.eventsChan <- buf[:n]
		}
	}
	close(w.eventsChan) // Close the channel to signal no more events will be sent
	close(w.errorsChan) // Close the channel to signal no more errors will be sent
}
