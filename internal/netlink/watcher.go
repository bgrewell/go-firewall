package netlink

func NewNetlinkMonitor() (watcher NetlinkWatcher, err error) {
	return &DefaultNetlinkWatcher{}, nil
}

type NetlinkWatcher interface {
	Monitor() (events <-chan []byte, errors <-chan error, setupErr error)
	Stop() error
	Running() bool
}
