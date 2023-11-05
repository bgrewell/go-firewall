package nftables

type TableOptions func(options *listOptions)

type Table interface {
	List(opts ...TableOptions)
	Add(name string, family AddressFamily) error
	Create(name string, family AddressFamily) error
	Delete(name string, family AddressFamily) error
	Flush(name string, family AddressFamily) error
	Destroy(name string, family AddressFamily) error
	Exists(name string, family AddressFamily) (bool, error)
}
