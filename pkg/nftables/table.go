package nftables

import "encoding/json"

type Table interface {
	//List(opts ...TableOptions)
	//Add(name string, family AddressFamily) error
	//Create(name string, family AddressFamily) error
	//Delete(name string, family AddressFamily) error
	//Flush(name string, family AddressFamily) error
	//Destroy(name string, family AddressFamily) error
	//Exists(name string, family AddressFamily) (bool, error)
}

func UnmarshalTable(tableJSON []byte) (table Table, err error) {
	var defaultTable DefaultTable
	if err := json.Unmarshal(tableJSON, &defaultTable); err != nil {
		return nil, err
	}

	return &defaultTable, nil
}
