package _old

// Cmd represents the default commands supported by iptables
type Cmd string

const (
	CmdAppend      Cmd = "append"
	CmdCheck       Cmd = "check"
	CmdDelete      Cmd = "delete"
	CmdInsert      Cmd = "insert"
	CmdReplace     Cmd = "replace"
	CmdList        Cmd = "list"
	CmdFlush       Cmd = "flush"
	CmdZero        Cmd = "zero"
	CmdNewChain    Cmd = "new-chain"
	CmdDeleteChain Cmd = "delete-chain"
	CmdPolicy      Cmd = "policy"
	CmdRenameChain Cmd = "rename-chain"
)
