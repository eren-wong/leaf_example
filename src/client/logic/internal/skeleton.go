package internal

import (
	"github.com/ida-wong/leaf/chanrpc"
	"github.com/ida-wong/leaf/module"

	"client/const/skeleton"
)

var (
	Skeleton = new(module.Skeleton)
)

func init() {
	Skeleton.GoLen = skeleton_const.GoLen
	Skeleton.TimerDispatcherLen = skeleton_const.TimerDispatcherLen
	Skeleton.AsyncCallLen = skeleton_const.AsyncCallLen
	Skeleton.ChanRPCServer = chanrpc.NewServer(skeleton_const.ChanRPCLen)
	Skeleton.Init()
}
