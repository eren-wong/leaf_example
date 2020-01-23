package internal

import (
	"github.com/ida-wong/leaf/gate"
	"client/config"
	"client/const/gate"
	. "client/internal"
	"client/logic"
)

var Client = new(gate.Client)

func init() {
	Client.AutoReconnect = gate_const.AutoReconnect
	Client.ConnectInterval = gate_const.ConnectInterval
	Client.PendingWriteNum = gate_const.PendingWriteNum
	Client.MaxMsgLen = gate_const.MaxMsgLen
	Client.Processor = JsonProcessor
	Client.AgentChanRPC = logic.ChanRpc

	Client.WSAddr = config.Client.WSAddr
	Client.WsConnNum = 1
	Client.HandshakeTimeout = gate_const.HandshakeTimeout

	Client.TCPAddr = config.Client.TCPAddr
	Client.LenMsgLen = gate_const.LenMsgLen
	Client.LittleEndian = gate_const.LittleEndian
	Client.TcpConnNum = 1
}
