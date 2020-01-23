package gate_const

import "time"

const (
	PendingWriteNum = 2000
	MaxMsgLen       = uint32(4096)
	HTTPTimeout     = 10 * time.Second
	LenMsgLen       = 2
	LittleEndian    = false
)
