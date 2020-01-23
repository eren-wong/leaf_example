package gate_const

import "time"

const (
	PendingWriteNum  = 2000
	MaxMsgLen        = uint32(4096)
	HandshakeTimeout = 10 * time.Second
	LenMsgLen        = 2
	AutoReconnect    = true
	ConnectInterval  = 10 * time.Second
	LittleEndian     = false
)
