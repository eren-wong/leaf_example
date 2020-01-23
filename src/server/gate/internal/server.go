package internal

import (
	"github.com/ida-wong/leaf/gate"

	"server/config"
	"server/const/gate"
	. "server/internal"
	"server/logic"
)

var Server = new(gate.Server)

func init() {
	Server.MaxConnNum = config.Server.MaxConnNum
	Server.PendingWriteNum = gate_const.PendingWriteNum
	Server.MaxMsgLen = gate_const.MaxMsgLen
	Server.Processor = JsonProcessor
	Server.AgentChanRPC = logic.ChanRpc

	Server.WSAddr = config.Server.WSAddr
	Server.HTTPTimeout = gate_const.HTTPTimeout
	Server.CertFile = config.Server.CertFile
	Server.KeyFile = config.Server.KeyFile

	Server.TCPAddr = config.Server.TCPAddr
	Server.LenMsgLen = gate_const.LenMsgLen
	Server.LittleEndian = gate_const.LittleEndian
}
