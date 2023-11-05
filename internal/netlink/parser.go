package netlink

import (
	"encoding/binary"
	"fmt"
)

// Netlink message header
type NlMsgHdr struct {
	Len   uint32
	Type  uint16
	Flags uint16
	Seq   uint32
	Pid   uint32
}

// ParseNetlinkMessageHeader parses the Netlink message header
func ParseNetlinkMessageHeader(b []byte) (*NlMsgHdr, error) {
	if len(b) < 16 {
		return nil, fmt.Errorf("not enough data for netlink message header")
	}
	return &NlMsgHdr{
		Len:   binary.LittleEndian.Uint32(b[0:4]),
		Type:  binary.LittleEndian.Uint16(b[4:6]),
		Flags: binary.LittleEndian.Uint16(b[6:8]),
		Seq:   binary.LittleEndian.Uint32(b[8:12]),
		Pid:   binary.LittleEndian.Uint32(b[12:16]),
	}, nil
}
