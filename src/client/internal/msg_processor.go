package internal

import (
	"github.com/ida-wong/leaf/network/json"
	"github.com/ida-wong/leaf/network/protobuf"
)

var (
	JsonProcessor     = json.NewProcessor()
	ProtobufProcessor = protobuf.NewProcessor()
)
