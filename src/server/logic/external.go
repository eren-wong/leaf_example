package logic

import "server/logic/internal"

var (
	Module  = new(internal.Module)
	ChanRpc = internal.Skeleton.ChanRPCServer
)
