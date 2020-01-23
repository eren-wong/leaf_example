package service

import (
	"github.com/ida-wong/leaf/gate"
	"github.com/ida-wong/leaf/log"

	"client/proto"
)

func init() {
	register(new(proto.Hello), onHello)
}

func onHello(args []interface{}) {
	msg := args[0].(*proto.Hello)
	agent := args[1].(gate.Agent)

	log.Debug("client hello %v, id: %d", msg.Name, msg.Id)
	log.Debug("local addr: %s", agent.LocalAddr())
	log.Debug("remote addr: %s", agent.RemoteAddr())

	agent.Close()
}
