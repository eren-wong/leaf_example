package rpc

import (
	"github.com/ida-wong/leaf/gate"

	"server/proto"
)

func Hello(agent gate.Agent, id int, name string) {
	msg := new(proto.Hello)
	msg.Id = id
	msg.Name = name

	agent.WriteMsg(msg)
}
