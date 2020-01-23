package chanrpc

import (
	"github.com/ida-wong/leaf/gate"
	"github.com/ida-wong/leaf/log"

	. "server/logic/internal"
)

func init() {
	Skeleton.RegisterChanRPC("NewTcpAgent", newTcpAgent)
	Skeleton.RegisterChanRPC("CloseTcpAgent", closeTcpAgent)
	Skeleton.RegisterChanRPC("NewWsAgent", newWsAgent)
	Skeleton.RegisterChanRPC("CloseWsAgent", closeWsAgent)
}

func newTcpAgent(args []interface{}) {
	agent := args[0].(gate.Agent)

	log.Debug("new tcp agent, remote address: %s, local address: %s", agent.RemoteAddr(), agent.LocalAddr())
}

func closeTcpAgent(args []interface{}) {
	agent := args[0].(gate.Agent)

	log.Debug("close tcp agent, remote address: %s, local address: %s", agent.RemoteAddr(), agent.LocalAddr())
}

func newWsAgent(args []interface{}) {
	agent := args[0].(gate.Agent)

	log.Debug("new web socket agent, remote address: %s, local address: %s", agent.RemoteAddr(), agent.LocalAddr())
}

func closeWsAgent(args []interface{}) {
	agent := args[0].(gate.Agent)

	log.Debug("close web socket agent, remote address: %s, local address: %s", agent.RemoteAddr(), agent.LocalAddr())
}
