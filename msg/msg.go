package msg

import "github.com/trist725/myleaf/network/protobuf"

var (
	//前后端包不同的操作
	ClientProcessor = protobuf.NewProcessor()
	ServerProcessor = protobuf.NewProcessor()
)
