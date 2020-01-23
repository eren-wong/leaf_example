package service

import (
	"reflect"

	. "client/internal"
	. "client/logic/internal"
)

func register(proto interface{}, method interface{}) {
	JsonProcessor.Register(proto)
	JsonProcessor.SetRouter(proto, Skeleton.ChanRPCServer)
	Skeleton.RegisterChanRPC(reflect.TypeOf(proto), method)
}
